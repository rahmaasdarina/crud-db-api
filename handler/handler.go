package handler

import (
	config "crud-api-homework/configs"
	"crud-api-homework/models"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var err error

func HandlerNewMovies(c *fiber.Ctx) error {
	//return c.SendString("Insert Data Movies")
	Db := config.DbConn

	movie := new(models.Movies)
	//movie := &models.Movies{}
	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	Db.Create(&movie)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    "success create data",
		"result": &movie,
	})
}

func HandlerGetMovies(c *fiber.Ctx) error {
	//return c.SendString("Read Data Movies")
	Db := config.DbConn

	slug := c.Params("slug")
	var movie models.Movies
	Db.Find(&movie, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "data not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success retrieve data",
		"result": &movie,
	})
}

func HandlerUpMovies(c *fiber.Ctx) error {
	//return c.SendString("Update Movies")
	Db := config.DbConn

	slug := c.Params("slug")
	movie := new(models.Movies)
	Db.First(&movie, slug)
	var mysqlErr *mysql.MySQLError

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "data not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": &movie,
	})
}

func HandlerDeleteMovie(c *fiber.Ctx) error {
	//return c.SendString("Delete Movies")
	Db := config.DbConn

	slug := c.Params("slug")
	var movie models.Movies
	Db.First(&movie, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "data not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	Db.Delete(&movie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "success",
	})

}
