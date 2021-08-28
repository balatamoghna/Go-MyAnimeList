package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"trademarkia.com/animeAPI/backend"
)

func TestScraper(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{1, "Cowboy Bebop"},
		{5114, "Fullmetal Alchemist: Brotherhood"},
		{14719, "JoJo no Kimyou na Bouken (TV)"},
	}

	for _, test := range tests {
		if output := backend.GetByID(test.input).Title; output != test.expected {
			t.Error("Test Failed: {} inputed, \n{} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}

func TestHandler(t *testing.T) {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(400)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/test", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 400, resp.StatusCode, "Status Code")
}
