package auth

import (
	"demo/linker/configs"
	"demo/linker/pkg/jwt"
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
		email, err := handler.AuthService.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := AuthResponse{
			Token: token,
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
		email, err := handler.AuthService.Register(regRequest.Email, regRequest.Password, regRequest.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := AuthResponse{
			Token: token,
		}
		response.Json(w, data, 200)
	}
}
