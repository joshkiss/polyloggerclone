package routes

import (
	"net/http"
	"strconv"

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

func updateEntry(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	entry, err := models.GetEntryByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch entry."})
		return
	}

	if entry.UserId != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update entry"})
		return
	}

	var updatedEntry models.Entry
	err = ctx.ShouldBindJSON(&updatedEntry)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEntry.ID = id
	err = updatedEntry.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update entry."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Entry updated successfully."})
}

// Delete entry
func deleteEntry(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse entry id."})
		return
	}

	entry, err := models.GetEntryByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch entry."})
		return
	}

	if entry.UserId != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = entry.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete entry."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
