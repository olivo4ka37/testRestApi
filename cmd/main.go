package main

import (
	"github.com/olivo4ka37/testRestApi"
	"github.com/olivo4ka37/testRestApi/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(testRestApi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
