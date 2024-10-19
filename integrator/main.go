package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Response struct {
	User User `json:"user"`
}

type Metadata struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}

type UsersList struct {
	Users    []User   `json:"users"`
	Metadata Metadata `json:"metadata"`
}

func getJson(url string, target interface{}) error {
	// TODO: better error handle
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func main() {
	app := fiber.New()

	app.Get("/get-users", func(c *fiber.Ctx) error {
		page := 1
		perPage := 1000
		totalPages := 1
		allUsers := []string{}

		for {
			url := fmt.Sprintf("http://gi-api:8000/users?page=%d&per_page=%d", page, perPage)
			usersList := UsersList{}
			getJson(url, &usersList)

			totalPages = int(math.Ceil(float64(usersList.Metadata.TotalCount) / float64(usersList.Metadata.PerPage)))

			for i := 0; i < len(usersList.Users); i++ {
				allUsers = append(allUsers, usersList.Users[i].PublicID)
			}

			page++
			if page == totalPages {
				break
			}
		}

		return c.JSON(allUsers[0:10])
	})

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
