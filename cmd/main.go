package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	ConnectDb()
	app:=fiber.New()

	SetupRoutes(app)

	app.Listen(":3000")
}