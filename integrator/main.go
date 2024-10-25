package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FullUser struct {
	PublicID  string `json:"public_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birth_date"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Response struct {
	User User `json:"user"`
}

type UserDataResponse struct {
	User FullUser `json:"user"`
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

// func timer(name string) func() {
// 	start := time.Now()
// 	return func() {
// 		fmt.Printf("%s took %v\n", name, time.Since(start))
// 	}
// }

func main() {
	app := fiber.New()

	app.Get("/get-users", func(c *fiber.Ctx) error {
		// defer timer("main")()
		var wg sync.WaitGroup
		var lock = sync.RWMutex{}
		usersMap := map[string]FullUser{}
		perPage := 1000
		pagesData := UsersList{}
		getJson("http://gi-api:8000/users?page=1&per_page=1", &pagesData)
		totalPages := int(math.Ceil(float64(pagesData.Metadata.TotalCount) / float64(perPage)))

		ch := make(chan UsersList)

		for page := 1; page <= totalPages; page++ {
			wg.Add(2)

			// make request
			go func(p int) {
				defer wg.Done()
				url := fmt.Sprintf("http://gi-api:8000/users?page=%d&per_page=%d", p, perPage)
				usersList := UsersList{}
				getJson(url, &usersList)
				ch <- usersList
			}(page)

			// process request
			go func() {
				lock.Lock()
				defer lock.Unlock()
				defer wg.Done()
				list := <-ch
				for i := 0; i < len(list.Users); i++ {
					usersMap[list.Users[i].PublicID] = FullUser{}
				}
			}()
		}

		wg.Wait()

		ch2 := make(chan UserDataResponse)

		// break all users IDs in chunks to prevent API overflow
		idsOnChunk := 0
		maxUsersOnChunk := 500
		chunks := [][]string{}
		chunk := make([]string, 0, maxUsersOnChunk)

		for userID := range usersMap {
			chunk = append(chunk, userID)
			idsOnChunk++
			if idsOnChunk == maxUsersOnChunk {
				tmp := make([]string, len(chunk))
				copy(tmp, chunk)
				chunks = append(chunks, tmp)
				chunk = chunk[:0]
				idsOnChunk = 0
			}
		}
		chunks = append(chunks, chunk)

		// get each user information
		for _, usersIDs := range chunks {
			for _, userID := range usersIDs {
				wg.Add(2)

				// make user request
				go func() {
					defer wg.Done()
					userData := UserDataResponse{}
					url := fmt.Sprintf("http://gi-api:8000/users/%s", userID)
					getJson(url, &userData)
					ch2 <- userData
				}()

				// process request
				go func() {
					lock.Lock()
					defer lock.Unlock()
					defer wg.Done()
					user := <-ch2
					usersMap[user.User.PublicID] = user.User
				}()
			}

			wg.Wait()
			fmt.Println(" ")
			fmt.Println("  -> going to next chunk")
			fmt.Println(" ")
			time.Sleep(1 * time.Second)
		}

		wg.Wait()

		return c.JSON(fiber.Map{
			"status":          "done",
			"users_processed": len(usersMap),
			"has_errors":      false,
		})
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
