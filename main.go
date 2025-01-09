package main

import (
	"log"

	"github.com/amadou-toure/groove_api/Database"
	"github.com/amadou-toure/groove_api/HTTP_CODE"
	"github.com/amadou-toure/groove_api/handlers"
	"github.com/gofiber/fiber/v2"
)



func main() {
	 err:=Database.Connect()
	if err != nil{
	 	log.Fatal(err.Error())
	}

	app := fiber.New()
	 app.Post("/user",handlers.CreateUser)
	app.Get("/users",handlers.GetUsers)
	// app.Get("/user/:id",handlers.GetUser)
	app.Put("/user/:id",handlers.UpdateUser)
	app.Delete("/user/:id",handlers.DeleteUser)


	 app.Get("/",func (c *fiber.Ctx)error{
	 	err := c.SendString("server is running")
		if err != nil{
			return c.Status(HTTP_CODE.Server_error).SendString("Server failed to respond")
		}
		return err
	 })
	app.Listen("localhost:3000")

}