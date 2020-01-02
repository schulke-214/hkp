package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/schulke-214/hkp/pkg/db"
	"github.com/schulke-214/hkp/pkg/domain/models"
)

// CreateEntry creates an entry and associates it with a given task id
func CreateEntry(ctx *gin.Context) {
	var entry models.TaskEntry
	var count int
	id := ctx.Params.ByName("id")

	if err := db.Connection.Model(&models.Task{}).Where("id = ?", id).Count(&count).Error; err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if count != 1 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := ctx.BindJSON(&entry); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entry.TaskID = uid

	if err := db.Connection.Create(&entry).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

func QueryAllEntries(ctx *gin.Context) {
	var entries []models.TaskEntry

	if err := db.Connection.Find(&entries).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, entries)
}

// THIS IS NOT OK
// U SHOULDNT HAVE TO QUERY ENTRIES FOR A TASK ID
// @todo(max)
func QueryAllEntriesByTask(ctx *gin.Context) {
	var entries []models.TaskEntry
	taskID := ctx.Params.ByName("id")

	if err := db.Connection.Where("task_id = ?", taskID).Find(&entries).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, entries)
}

func DeleteEntry(ctx *gin.Context) {
	var entry models.TaskEntry
	taskID := ctx.Params.ByName("id")
	entryID := ctx.Params.ByName("entry_id")

	if err := db.Connection.Where("id = ? AND task_id = ?", entryID, taskID).Delete(&entry).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}
