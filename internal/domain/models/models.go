package models

type User struct {
	Name       string `Json: "name"`
	Email      string `Json: "email"`
	Password   string `Json: "password"`
	CreateTime string `Json: "create_time"`
}
