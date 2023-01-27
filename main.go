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
	controller := api.Controller{Closer: closer, Reservations: make(map[string]struct{})}

	app.Post("/close", controller.HandleCloseRequest)

	app.Post("/reservation/:id", controller.HandlePostReservationRequest)
	app.Get("/reservation/:id", controller.HandleGetReservationRequest)

	go func() {
		log.Fatal(app.Listen(":3000"))
	}()

	<-closer

	time.Sleep(time.Second)
}
