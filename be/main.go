package main

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "*",
	}))

	app.Use(controllers.Logger)

	app.Get("/dashboard", monitor.New())
	app.Get("/health", controllers.Health)

	app.Post("/SignIn", controllers.Signin)
	app.Post("/SignUp", controllers.SignUp)

	wizard := app.Group("/wizard", controllers.VerifyMiddleWare)
	//wizard.Get("/GetCorrection/:Key", controllers.Correction)

	wizard.Get("/NewCorrection", controllers.Bolus)
	wizard.Post("/Corrections", controllers.SaveCorrections)
	wizard.Get("/Corrections", controllers.Corrections)
	wizard.Get("/CorrectionRange", controllers.CorrectionRange)

	// saves an array of blood sugar targets with time stamps through out the day, used to calculates carb corrections
	wizard.Post("/Targets", controllers.SaveTargets)
	wizard.Get("/Targets", controllers.Targets)
	wizard.Get("/CurrentTarget", controllers.CurrentTarget)

	// saves an array of carb ratios with time stamps through out the day, used to calculate carb corrections
	wizard.Post("/Ratios", controllers.SaveRatios)
	wizard.Get("/Ratios", controllers.Ratios)
	wizard.Get("/CurrentRatio", controllers.CurrentRatio)

	// saves an array of insuling sensitivity factors, used to calculate blood sugar and carb corrections
	wizard.Post("/ISFs", controllers.SaveISFs)
	wizard.Get("/ISFs", controllers.ISFs)
	wizard.Get("/CurrentISF", controllers.CurrentIsf)

	// saves an array of insuling sensitivity factors, used to calculate blood sugar and carb corrections
	wizard.Post("/Duration", controllers.SaveDuration)
	wizard.Get("/Duration", controllers.Duration)

	app.Listen(":80")
}
