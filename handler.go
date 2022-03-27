package main

import "github.com/gofiber/fiber/v2"

type handler struct {
	service Service
}

func newHandler(service Service) handler {
	return handler{
		service: service,
	}
}

func (h handler) hanldeCreateCode(c *fiber.Ctx) error {
	code, err := h.service.Create()
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(code)
}
