package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Birthdate       string `json:"birthdate"`
}

var userLogged = []user{}

var users = []user{
	{Username: "juanmiloz", Password: "juancamilo", ConfirmPassword: "juancamilo", FirstName: "Juan", LastName: "Zorrilla", Birthdate: "13/03/2002"},
	{Username: "jpSanin", Password: "sanincho", ConfirmPassword: "sanincho", FirstName: "Juan Pablo", LastName: "Sanin", Birthdate: "16/05/2001"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("login.html", "create.html", "users.html")
	router.GET("/", defaultRedirect)
	router.GET("/users", loadViewLogin)
	router.POST("/users", login)
	router.GET("/create", loadViewCreate)
	router.POST("/create", addUser)
	router.GET("/list", loadViewList)
	router.GET("/logout", logout)
	router.Run("localhost:8080")
}

func defaultRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/users")
}

func loadViewList(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

func loadViewCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"answer":" ",
	})
}

func loadViewLogin(c *gin.Context) {

	if len(userLogged) != 0 {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"user":  userLogged,
			"users": users,
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"message": " ",
		})
	}
}

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	birthdate := c.PostForm("birthdate")

	if len(username) > 0 && len(password) > 0 && len(confirmPassword) > 0 && len(firstname) > 0 && len(lastname) > 0 && len(birthdate) > 0 {
		if password == confirmPassword {
			newUser := user{Username: username, Password: password, ConfirmPassword: confirmPassword, FirstName: firstname, LastName: lastname, Birthdate: birthdate}
			users = append(users, newUser)
			c.HTML(http.StatusOK, "login.html", gin.H{
				"message": "your user was create successfully",
			})
		} else {
			c.HTML(http.StatusOK, "create.html", gin.H{
				"answer": "The passwords are not equals",
			})
		}
	} else {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"answer": "There can be no empty fields",
		})
	}
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

func logout(c *gin.Context) {
	userLogged = []user{}
	c.Redirect(http.StatusMovedPermanently, "/users")
}
