package main

import (
	"github.com/gofiber/fiber/v2"
	"trademarkia.com/animeAPI/backend"
)

//Routers function used to route each request to appropriate url
func Routers(app *fiber.App) {
	app.Get("/", backend.Homepage)
	app.Get("/about", backend.AboutMe)
	app.Get("/anime/:id", backend.FindAnime)
}

func main() {
	backend.InitialMigration()
	app := fiber.New()

	Routers(app)
	app.Listen(":3000")

}
