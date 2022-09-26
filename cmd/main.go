package main

import (
	"github.com/sirupsen/logrus"
	"github.com/strimenko/store-strimenko/api"
	"github.com/strimenko/store-strimenko/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(api.Server)
	if err := srv.Run(":8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
