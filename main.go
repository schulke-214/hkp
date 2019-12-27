package main

import (

	// autoconfigure the environment
	_ "github.com/schulke-214/hkp/pkg/util/env"

	"github.com/gin-gonic/gin"

	"github.com/schulke-214/hkp/api/tasks"
	"github.com/schulke-214/hkp/pkg/db"
)

func main() {
	router := gin.Default()

	taskRouter := router.Group("/tasks")

	taskRouter.POST("/", tasks.Create)
	taskRouter.GET("/", tasks.QueryAll)
	taskRouter.GET("/:id", tasks.QueryByID)
	taskRouter.PATCH("/:id", tasks.UpdateByID)
	taskRouter.DELETE("/:id", tasks.DeleteByID)

	entriesRouter := taskRouter.Group("/:id/entries")

	entriesRouter.GET("/", tasks.QueryAllEntries)
	entriesRouter.POST("/", tasks.CreateEntry)
	entriesRouter.DELETE("/:entry_id", tasks.DeleteEntry)

	router.Run(":8080")

	defer db.Connection.Close()
}
