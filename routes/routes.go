package routes

import (
	"crud-api-homework/handler"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	fmt.Println("Masuk route")
	app := fiber.New()
	app.Post("/create-movie", handler.HandlerNewMovies)           //Insert
	app.Get("/get-movie/:slug", handler.HandlerGetMovies)         //Read
	app.Put("/update-movie/:slug", handler.HandlerUpMovies)       //Update
	app.Delete("/delete-movie/:slug", handler.HandlerDeleteMovie) //Delete

	app.Listen(":8000")
}
