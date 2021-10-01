package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"messege": "Status OK",
		})
	})

	app.Get("/:id", getUser)
	app.Get("/follower/:username", getFollowers)

	app.Listen(":3000")
}

func getFollowers(c *fiber.Ctx) error {
	param := c.Params("username")
	if val, ok := data[param]; ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"messege":   "Status OK",
			"followers": val.Followers,
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"messege": "Not Found",
	})
}

func getUser(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 32, 10)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"messege": "Params Not Found",
		})
	}

	for _, v := range data {
		if v.ID == uint(id) {
			return c.Status(fiber.StatusOK).JSON(v)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"messege": "Not Found",
	})
}

type user struct {
	ID        uint   `json:"-"`
	Username  string `json:"username"`
	Followers uint   `json:"followers"`
}

var data = map[string]user{
	"sammy": {
		ID:        1,
		Username:  "SammyShark",
		Followers: 987,
	},
	"jesse": {
		ID:        2,
		Username:  "JesseOctopus",
		Followers: 432,
	},
	"drew": {
		ID:        3,
		Username:  "DrewSquid",
		Followers: 312,
	},
	"jamie": {
		ID:        4,
		Username:  "JamieMantisShrimp",
		Followers: 654,
	},
}
