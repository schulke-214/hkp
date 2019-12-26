package api

import (
	"github.com/gin-gonic/gin"
)

// RegisterCategories registers all category endpoints on a gin router under a given path
func RegisterCategories(router *gin.Engine, path string) {
	group := router.Group(path)

	group.GET("/lol", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
