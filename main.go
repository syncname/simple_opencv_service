package main

import (
	"cvrecognizer/internal/api"
	"cvrecognizer/internal/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	c, err := config.ReadConfig("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	routerApi := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app, err := api.NewApp(c, routerApi)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Start())

}
