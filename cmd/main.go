package main

import (
	"context"
	"demo/linker/configs"
	"demo/linker/internal/auth"
	"demo/linker/internal/link"
	"demo/linker/internal/user"
	"demo/linker/pkg/db"
	"demo/linker/pkg/middleware"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctxWithTimeOut, cencel := context.WithTimeout(ctx, 4*time.Second)
	defer cencel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Done task")
	case <-ctxWithTimeOut.Done():
		fmt.Println("Timeout")
	}
}

func main2() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	//repositories
	linkrep := link.NewLinkRepo(database)
	userrepo := user.NewUserRepository(database)

	//Services
	authService := auth.NewAuthService(userrepo)

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkrep,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
