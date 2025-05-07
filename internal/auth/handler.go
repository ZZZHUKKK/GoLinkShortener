package auth

import (
	"demo/linker/configs"
	"demo/linker/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
		var loginRequest LoginRequest
		err := json.NewDecoder(req.Body).Decode(&loginRequest)
		if err != nil {
			response.Json(w, err.Error(), 402)
		}
		fmt.Println(loginRequest)
		data := AuthResponse{
			Token: "123",
		}
		response.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}

}
