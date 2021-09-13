package main

import (
	"net/http"
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

var userLogged = []user{}

var users = []user{
	{Username: "juanmiloz", Password: "juancamilo", ConfirmPassword: "juancamilo", FirstName: "Juan", LastName: "Zorrilla", Birthdate: time.Now()},
	{Username: "jpSanin", Password: "sanincho", ConfirmPassword: "sanincho", FirstName: "Juan", LastName: "Sanin", Birthdate: time.Now()},
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "create.html", "users.html")
	router.GET("/", defaultRedirect)
	router.GET("/users", loadViewLogin)
	router.POST("/users", login)
	//router.GET("/create", addUser)

	router.Run("localhost:8080")
}

func defaultRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func loadViewLogin(c *gin.Context) {
	if len(userLogged) != 0 {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"user":  userLogged,
			"users": users,
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
func addUser(c *gin.Context) {
	/*username := c.PostForm("username")
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

func login(c *gin.Context) {
	username := c.PostForm("Username")
	password := c.PostForm("Password")

	for _, a := range users {
		if username == a.Username {
			if password == a.Password {
				userLogged := a
				c.HTML(http.StatusOK, "users.html", gin.H{
					"username": userLogged.Username,
					"users":    users,
				})
				return
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"message": "Incorrect Password",
				})
				return
			}
		}

	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "This user doesn't exist",
	})

}
