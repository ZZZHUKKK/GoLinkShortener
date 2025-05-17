package main

import (
	"demo/linker/configs"
	"demo/linker/internal/auth"
	"demo/linker/internal/link"
	"demo/linker/internal/user"
	"demo/linker/pkg/db"
	"demo/linker/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
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
		Config:         conf,
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
