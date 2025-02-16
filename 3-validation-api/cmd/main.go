package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/auth"
	"3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	confVerify := configs.LoadVerifyConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandler{
		Config: conf,
	})
	verify.NewVerifyHandler(router, verify.VerifyHandler{
		VerifyMailConfig: confVerify,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server start 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
