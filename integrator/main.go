package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Response struct {
	User User `json:"user"`
}

func getJson(url string, target *Response) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := Response{}
		getJson("http://gi-api:8000/users/mFL2DS5KmpcL", &response)
		fmt.Println(response.User)

		data := User{
			PublicID: response.User.PublicID,
			Name:     response.User.Name,
			Email:    response.User.Email,
		}
		return c.JSON(data)
	})

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3002"))
}
