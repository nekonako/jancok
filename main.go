package main

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
)

var lastCode int

func main() {

	runtime.GOMAXPROCS(4)

	db := newDb()
	repo := newRepository(db)
	service := NewService(repo)
	handler := newHandler(service)

	app := fiber.New()
	app.Get("/", handler.hanldeCreateCode)

	app.Listen(":3000")

}
