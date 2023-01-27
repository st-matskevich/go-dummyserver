package api

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

const MAX_PEERS = 3

type Controller struct {
	Reservations map[string]struct{}
	Closer       chan struct{}
	mu           sync.Mutex
}

func (controller *Controller) reserve(id string) bool {
	controller.mu.Lock()
	defer controller.mu.Unlock()

	if len(controller.Reservations) >= MAX_PEERS {
		return false
	}

	controller.Reservations[id] = struct{}{}
	return true
}

func (controller *Controller) checkReservation(id string) bool {
	controller.mu.Lock()
	defer controller.mu.Unlock()

	_, ok := controller.Reservations[id]
	return ok
}

func (controller *Controller) HandleCloseRequest(c *fiber.Ctx) error {
	controller.Closer <- struct{}{}

	return c.SendStatus(fiber.StatusOK)
}

func (controller *Controller) HandlePostReservationRequest(c *fiber.Ctx) error {
	stringID := c.Params("id")

	if controller.reserve(stringID) {
		return c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusForbidden)
}

func (controller *Controller) HandleGetReservationRequest(c *fiber.Ctx) error {
	stringID := c.Params("id")

	if controller.checkReservation(stringID) {
		return c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusNotFound)
}
