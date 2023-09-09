package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id				string
	Firstname  string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "John",
		Lastname: "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusCreated).JSON(user)
}


func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/users")
	userGroup.Get("",handleUser)
	userGroup.Post("" , handleCreateUser)

	app.Listen(":3000")
}
