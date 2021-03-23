package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/GoFiberSQLLogin/Controller"
	"github.com/kriangkrai/GoFiberSQLLogin/Router"
)

func main() {
	if err := Controller.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	api := app.Group("/api") // /api

	api.Get("/", Router.Gets)

	// var port = os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	port := "3000"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}

	app.Listen(":" + port)

}
