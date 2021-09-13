package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmPassword"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Birthdate       time.Time `json:"time"`
}

var users = []user{
	{Username: "juanmiloz", Password: "juancamilo", ConfirmPassword: "juancamilo", FirstName: "Juan", LastName: "Zorrilla", Birthdate: time.Now()},
	{Username: "jpSanin", Password: "sanincho", ConfirmPassword: "sanincho", FirstName: "Juan", LastName: "Sanin", Birthdate: time.Now()},
}

func main() {
	router := gin.Default()
	router.GET("/create", addUser)

	router.Run("localhost:8080")
}

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthdate := c.PostForm("birthdate")

	newUser := user{Username: username, Password: password, ConfirmPassword: confirmPassword, FirstName: firstname, LastName: lastname, Birthdate: birthdate}

	append(users, newUser)

	/*if username != nil && password != nil && confirmPassword != nil && firstname != nil && lastname != nil && birthdate != nil{
		if password == confirmPassword{

		}else{

		}
	}else{

	}*/
}
