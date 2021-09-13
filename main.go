package main;

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Username String `json:"username"`
	Password String `json:"password"`
	ConfirmPassword String `json:"confirmPassword"`
	Username String `json:"firstName"`
	Username String `json:"lastName"`
}