package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshkiss/polyloggerclone/models"
)

func createEntry(ctx *gin.Context) {
	var entry models.Entry
	err := ctx.ShouldBindJSON(&entry)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := ctx.GetInt64("userId")
	entry.UserId = userId
	err = entry.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save entry."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Entry saved.", "entry": entry})
}
