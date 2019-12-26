package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/schulke-214/hkp/pkg/db"
	"github.com/schulke-214/hkp/pkg/domain/models"
)

// Create creates a category
func Create(ctx *gin.Context) {
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Connection.Create(&category).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// QueryAll returns all categories
func QueryAll(ctx *gin.Context) {
	var categories []models.Category

	if err := db.Connection.Find(&categories).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

// QueryByID returns a category by id
func QueryByID(ctx *gin.Context) {
	var category models.Category
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).First(&category).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// UpdateByID updates a category by id and returns it
func UpdateByID(ctx *gin.Context) {
	var category models.Category
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).First(&category).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := ctx.BindJSON(&category); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := db.Connection.Save(&category).Error; err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteByID deletes a category by id
func DeleteByID(ctx *gin.Context) {
	var category models.Category
	id := ctx.Params.ByName("id")

	if err := db.Connection.Where("id = ?", id).Delete(&category).Error; err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}

// RegisterCategories registers all category endpoints on a gin router under a given path
func RegisterCategories(router *gin.Engine) {
	group := router.Group("/categories")

	group.GET("/", QueryAll)
	group.GET("/:id", QueryByID)
	group.POST("/create", Create)
	group.PATCH("/:id", UpdateByID)
	group.DELETE("/:id", DeleteByID)
}
