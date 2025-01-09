package models

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	User_name string             `bson:"user_name"`
	Birth_date string             `bson:"birth_date"`
	Interest []string             `bson:"interest"`
}
