package main

import (
	"4-order-api/configs"
	"4-order-api/internal/auth"
	"4-order-api/internal/link"
	"4-order-api/internal/product"
	"4-order-api/internal/verify"
	"4-order-api/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	dbConf := db.NewDB(conf)
	authConf := configs.LoadAuthConfig()
	verifyConf := configs.LoadVerifyConfig()
	router := http.NewServeMux()

	//Repository
	linkRepository := link.NewLinkRepository(dbConf)
	productRepository := product.NewProductRepository(dbConf)
	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthConfig: authConf,
	})

	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		VerifyMailConfig: verifyConf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	//
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
