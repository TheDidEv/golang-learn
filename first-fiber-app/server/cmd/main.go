package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	todo "github.com/thedidev/golang-projects/first-fiber-app/server"
)

func main() {
	app := fiber.New()

	// middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []todo.Todo{}

	app.Get("/api/todo", func(c *fiber.Ctx) error { //func(c *fiber.Ctx) - Context
		return c.JSON(todos)
	})

	app.Post("/api/todo", func(c *fiber.Ctx) error {
		td := &todo.Todo{}

		if err := c.BodyParser(td); err != nil {
			return err
		}

		td.Id = len(todos) + 1

		todos = append(todos, *td)

		return c.JSON(todos)
	})

	app.Patch("/api/todo/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.Id == id {
				todos[i].Done = !todos[i].Done
				break
			}
		}

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
