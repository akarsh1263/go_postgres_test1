package main

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hey there delilah!")
}

func CreateMovie(c *fiber.Ctx) error {
	movie:=new(Movie)
	if err:=c.BodyParser(movie); err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	DB.Db.Create(&movie)

	return c.Status(200).JSON(movie)
}

func ListMovies(c *fiber.Ctx) error {
	movies:= []Movie{}

	DB.Db.Find(&movies)

	return c.Status(200).JSON(movies)
}

func UpdateMovie(c *fiber.Ctx) error {
	movie:= new(Movie)
	id:= c.Params("id")

	if err:= c.BodyParser(movie); err!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result:= DB.Db.Model(&movie).Where("id=?",id).Updates(movie)
	if result.Error!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(200).JSON(movie)
}

func DeleteMovie(c *fiber.Ctx) error {
	movie:= Movie{}
	id:= c.Params("id")

	result:= DB.Db.Where("id=?",id).Delete(&movie)
	if result.Error!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ListMovies(c)
}