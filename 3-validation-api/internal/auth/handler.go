package auth

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/req"
	"3-validation-api/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, config AuthHandler) {
	handler := &AuthHandler{
		Config: config.Config,
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
