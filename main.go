package main

import (
	"github.com/gofiber/fiber/v2"
)

var lastCode int

var code Code

func main() {

	db := newDb()
	repo := newRepository(db)
	service := NewService(repo)
	handler := newHandler(service)

	app := fiber.New()
	app.Get("/", handler.hanldeCreateCode)

	app.Listen(":3000")

}
