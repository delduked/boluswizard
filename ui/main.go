package main

import (
	"ui/controllers"
	"ui/middleware"
	"ui/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))
	app.Use(middleware.Log)

	app.Get("/error/:error", models.ErrorPage)
	app.Use("/ws", middleware.SocketUpgrade)
	app.Get("/ws/:id", websocket.New(controllers.Ws))
	go controllers.ListenToChannel()

	app.Get("/Login", models.Login)
	app.Get("/SignUp", models.Signup) // change to new sign up page

	app.Use("/u", middleware.ValidateToken)
	app.Use("/u/:user", middleware.VerifyUser)

	// final URL /u/nate.delduca@gmail.com/home
	app.Get("/u/:user/home", models.Home)

	assets := app.Group("/assets", middleware.ValidateToken)
	assets.Static("/", "./assets")

	// handle page not found errors
	app.Use(models.ErrorRedirect)

	app.Listen(":80")
}
