package handlers

import (
	"net/http"

	"github.com/damianpugliese/go_rest_api_events_booking/models"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}
	
	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the user", "error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}