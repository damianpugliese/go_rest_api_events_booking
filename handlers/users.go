package handlers

import (
	"net/http"

	"github.com/damianpugliese/go_rest_api_events_booking/models"
	"github.com/damianpugliese/go_rest_api_events_booking/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not save the user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}
	
	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Could not authenticate the user",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not generate token",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User authenticated successfully",
		"token":   token,
	})
}