package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	}
}

type GenDice struct{}

func NewGenDice(router *http.ServeMux) {
	handler := &GenDice{}
	router.HandleFunc("/gendice", handler.GenDice())
}

func (handler *GenDice) GenDice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		num := rand.Intn(5) + 1
		fmt.Println(num)
	}
}
