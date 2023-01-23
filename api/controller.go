package api

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

const MAX_PEERS = 3

type Controller struct {
	Counter uint
	Closer  chan struct{}
	mu      sync.Mutex
}

func (controller *Controller) increaseCounter() bool {
	controller.mu.Lock()
	defer controller.mu.Unlock()

	if controller.Counter >= MAX_PEERS {
		return false
	}

	controller.Counter++
	return true
}

func (controller *Controller) HandleCloseRequest(c *fiber.Ctx) error {
	controller.Closer <- struct{}{}

	return c.SendStatus(fiber.StatusOK)
}

func (controller *Controller) HandleReserveRequest(c *fiber.Ctx) error {
	if controller.increaseCounter() {
		return c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusForbidden)
}
