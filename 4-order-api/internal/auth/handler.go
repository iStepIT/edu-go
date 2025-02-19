package auth

import (
	"4-order-api/configs"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.AuthConfig
}
type AuthHandler struct {
	*configs.AuthConfig
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthConfig: deps.AuthConfig,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
	}
}
