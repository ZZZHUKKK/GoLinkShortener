package auth

import (
	"demo/linker/configs"
	"demo/linker/pkg/request"
	"demo/linker/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Login")
		loginRequest, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
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
		regRequest, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.Json(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		fmt.Println(regRequest)
		_, err = handler.AuthService.Register(regRequest.Email, regRequest.Password, regRequest.Name)
		if err != nil {
			response.Json(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		response.Json(w, map[string]string{"message": "User registered successfully"}, http.StatusCreated)
	}
}
