package main

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/",Home)
	app.Post("/createmovie",CreateMovie)
	app.Get("/listmovies",ListMovies)
	app.Patch("/updatemovie/:id",UpdateMovie)
	app.Delete("/deletemovie/:id",DeleteMovie)
}