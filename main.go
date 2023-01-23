package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/st-matskevich/go-dummyroom/api"
)

func main() {
	log.Println("Starting dummy server")

	app := fiber.New()

	app.Use(
		logger.New(),
	)

	closer := make(chan struct{})
	controller := api.Controller{Closer: closer, Counter: 0}

	app.Post("/close", controller.HandleCloseRequest)
	app.Post("/reserve", controller.HandleReserveRequest)

	go func() {
		log.Fatal(app.Listen(":3000"))
	}()

	<-closer

	time.Sleep(time.Second)
}
