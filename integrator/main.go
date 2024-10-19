package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func setUpRoutes(app *fiber.App) {
// 	app.Get("/hello", routes.Hello)
// 	// app.Get("/allbooks", routes.AllBooks)
// 	// app.Post("/addbook", routes.AddBook)
// 	// app.Post("/book", routes.Book)
// 	// app.Put("/update", routes.Update)
// 	// app.Delete("/delete", routes.Delete)
// }

type User struct {
	Name  string
	Email string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := User{
			Name:  "Grame",
			Email: "grame@test.com",
		}
		return c.JSON(data)
		// return c.SendString("Hello integrator!")
	})

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3002"))
}
