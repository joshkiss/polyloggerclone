package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshkiss/polyloggerclone/models"
	"github.com/joshkiss/polyloggerclone/utils"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Validate()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
