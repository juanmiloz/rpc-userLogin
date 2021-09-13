package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmPassword"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Birthdate       string 	  `json:"birthdate"`
}

var users = []user{
	{Username: "juanmiloz", Password: "juancamilo", ConfirmPassword: "juancamilo", FirstName: "Juan", LastName: "Zorrilla", Birthdate: "13/03/2002"},
	{Username: "jpSanin", Password: "sanincho", ConfirmPassword: "sanincho", FirstName: "Juan", LastName: "Sanin", Birthdate: "16/05/2001"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "create.html", "users.html")
	router.GET("/",defaultRedirect)
	router.GET("/create", addUser)

	router.Run("localhost:8080")
}

func defaultRedirect(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthdate := c.PostForm("birthdate")

	if len(username) > 0 && len(password) > 0 && len(confirmPassword) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(birthdate) > 0{
		if password == confirmPassword{
			newUser := user{Username: username, Password: password, ConfirmPassword: confirmPassword, FirstName: firstname, LastName: lastname, Birthdate: birthdate}
			users = append(users, newUser)
			c.HTML(http.StatusOK , "create.html", gin.H{
				"answer": "your user was create successfully",
			})
		}else{
			
		}
	}else{

	}
}