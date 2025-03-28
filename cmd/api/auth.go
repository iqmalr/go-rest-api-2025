package main

import (
	"go-rest-api/cmd/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
	Name string `json:"name" binding:"required,min=2"`
	// Email    string `json:"email"`
	// Name     string `json:"name"`
}

func (app *application) registerUser(c *gin.Context){
	var register registerRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error":"Something went wrong"})
		return
	}

	register.Password = string(hashedPassword)
	user := database.User{
		Email: register.Email,
		Password: register.Password,
		Name: register.Name,
	}
	err=app.models.Users.Insert(&user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}