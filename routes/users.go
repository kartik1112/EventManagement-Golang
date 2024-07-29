package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/EventManagement-Golang/models"
	"github.com/kartik1112/EventManagement-Golang/utils"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "COuld not parse Data",
		})
		return
	}
	err = user.Save()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "COuld not save user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not Authenticate user",
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not Authenticate user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User logged in",
		"token":   token,
	})

}
