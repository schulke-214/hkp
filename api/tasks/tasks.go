package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/schulke-214/hkp/pkg/db"
	"github.com/schulke-214/hkp/pkg/domain/models"
)

// Create creates a task
func Create(ctx *gin.Context) {
	var task models.Task

	if err := ctx.BindJSON(&task); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Connection.Create(&task).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// QueryAll returns all task
func QueryAll(ctx *gin.Context) {
	var categories []models.Task

	if err := db.Connection.Find(&categories).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// QueryByID returns a task by id
func QueryByID(ctx *gin.Context) {
	var task models.Task
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).First(&task).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// UpdateByID updates a task by id and returns it
func UpdateByID(ctx *gin.Context) {
	var task models.Task
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).First(&task).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := ctx.BindJSON(&task); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := db.Connection.Save(&task).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// DeleteByID deletes a task by id
func DeleteByID(ctx *gin.Context) {
	var task models.Task
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).Delete(&task).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}
