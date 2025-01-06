package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type  MongoInstance struct {
	Client *mongo.Client
	Db *mongo.Database
}
type CODE struct{
	Ok int
	Created int
	Accepted int
	Bad_request int
	Forbiden int
	Not_found int
	Time_out int
	Server_error int
	Insufiscient_staorage int
	Loop int
}
var HTTP_CODE=CODE{
	Ok:200,
	Created:201,
	Accepted:202,
	Bad_request:400,
	Forbiden:403,
	Not_found:404,
	Time_out:408,
	Server_error:500,
	Insufiscient_staorage:507,
	Loop:508,
}




var mg MongoInstance
const dbName="groove_DB"
const mongoURI="mongodb://localhost:27017/" + dbName

func connect() error {
	client,err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db:= client.Database(dbName)
	if err!= nil {
		return err
	}
	mg= MongoInstance{Client:client,Db:db}
	return nil
}


//Model for User
type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	User_name string             `bson:"user_name"`
	Birth_date string             `bson:"birth_date"`
	Interest []string             `bson:"interest"`
}


func CreateUser(c *fiber.Ctx)error{
	collection := mg.Db.Collection("Users")
	var newUser User
	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	newUser.ID = ""
	result,err:=collection.InsertOne(c.Context(),newUser)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	
	return c.Status(HTTP_CODE.Created).SendString("user " + newUser.Name + " created with id " + result.InsertedID.(primitive.ObjectID).Hex())
 
}

func GetUser(c *fiber.Ctx)error{
	collection:=mg.Db.Collection("Users")
	filter:=bson.D{{}}
	var users []User
	 result,err := collection.Find(c.Context(),filter)
	 if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("user not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 err = result.All(c.Context(),&users)
	 if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 return c.Status(HTTP_CODE.Ok).JSON(users)

}

func updateUser(c *fiber.Ctx)error{
id:=c.Params("id")
userID,err:= primitive.ObjectIDFromHex(id)
if err != nil {
	return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
}
user:= new(User)
err = c.BodyParser(user)
if err != nil {
	return c.Status(HTTP_CODE.Bad_request).SendString(err.Error())
}
collection:=mg.Db.Collection("Users")
filter:=bson.D{{Key:"_id",Value: userID}}
update:=bson.D{{Key:"$set",Value: bson.D{
	{Key:"name",Value: user.Name},
	{Key:"email",Value: user.Email},
	{Key:"password",Value: user.Password},
	{Key:"user_name",Value: user.User_name},
	{Key:"birth_date",Value: user.Birth_date},
	{Key:"interest",Value: user.Interest},
}}}
err = collection.FindOneAndUpdate(c.Context(),filter,update).Err()
if err != nil {
	if err == mongo.ErrNoDocuments {
		return c.Status(HTTP_CODE.Not_found).SendString("user not found")
	}
	return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
}
return c.Status(HTTP_CODE.Ok).JSON(user)
}

func deleteUser(c *fiber.Ctx)error{
	id:=c.Params("id")
	userId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}
	collection:=mg.Db.Collection("Users")
	filter:=bson.D{{Key:"_id",Value: userId}}
	err = collection.FindOneAndDelete(c.Context(),filter).Err()
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("user not foound")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).SendString("user deleted")


}

func main() {
	err:=connect()
	if err != nil{
		log.Fatal(err)
	}
	app := fiber.New()
	app.Post("/user",CreateUser)
	app.Get("/users",GetUser)
	app.Put("/user/:id",updateUser)
	app.Delete("/user/:id",deleteUser)


	app.Get("/",func (c *fiber.Ctx)error{
		return c.SendString("Welcome to the user page")
	})
	app.Listen("localhost:3000")

}