package main

import (
	"log"

	"github.com/acool-kaz/api-gateway-service/internal/app"
	"github.com/acool-kaz/api-gateway-service/internal/config"
)

func main() {
	cfg, err := config.InitConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	app := app.InitApp(cfg)

	app.Run()
}
