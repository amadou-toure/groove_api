package main

import (
	"log"
	"os"

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
	app.Post("/user/login",handlers.Login)
	app.Get("/users",handlers.GetUsers)
	app.Get("/user/:id",handlers.GetOneUser)
	app.Put("/user/:id",handlers.UpdateUser)
	app.Delete("/user/:id",handlers.DeleteUser)


	 app.Get("/",func (c *fiber.Ctx)error{
	 	err := c.SendString(":" + os.Getenv("PORT"))
		if err != nil{
			return c.Status(HTTP_CODE.Server_error).SendString("Server failed to respond")
		}
		return err
	 })
	app.Listen(":3000")

}