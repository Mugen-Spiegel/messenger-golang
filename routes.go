package main

import (
	"messenger/app/controllers/admin"
	v1 "messenger/app/controllers/v1"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/admin/user", admin.UserCreate)
	router.GET("/admin/users", admin.UserList)
	router.GET("/:user_id/task", v1.TaskList)
	router.POST("/:user_id/task", v1.TaskCreate)
	router.DELETE(":user_id/task/:task_id", v1.TaskDelete)
	router.PATCH("/task", v1.TaskUpdate)
	return router
}
