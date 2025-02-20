package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)
	NewGenDice(router)

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
