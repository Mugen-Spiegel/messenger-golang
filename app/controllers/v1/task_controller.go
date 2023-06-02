package v1

import (
	"fmt"
	"messenger/app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TaskList(r *gin.Context) {
	user_id, _ := strconv.Atoi(r.Param("user_id"))
	task := []models.Task{}
	result := models.DB.Where("user_id = ?", user_id).Find(&task)
	println(result.RowsAffected)
	r.JSON(http.StatusOK, gin.H{"data": task})
}

func TaskCreate(r *gin.Context) {
	user_id, _ := strconv.Atoi(r.Param("user_id"))
	task_name := r.PostForm("task_name")
	task_description := r.PostForm("task_description")
	status := r.PostForm("status")

	task := models.Task{
		Task:            task_name,
		TaskDescription: task_description,
		Status:          status,
		UserID:          user_id,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	result := models.DB.Create(&task)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}
	models.DB.Preload("User").Find(&task)

	r.JSON(http.StatusOK, gin.H{"data": task})
}

type TaskDeleteParams struct {
	UserID int `uri:"user_id"`
	TaskID int `uri:"task_id"`
}

func TaskDelete(r *gin.Context) {
	var task_params TaskDeleteParams
	if r.ShouldBindUri(&task_params) != nil {
		panic("Invalid Params")
	}
	task := models.Task{
		ID:     task_params.TaskID,
		UserID: task_params.UserID,
	}
	fmt.Println(task, "asdasasas")
	result := models.DB.Delete(&task)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}

	r.JSON(http.StatusOK, "")
}

type TaskUpdateParams struct {
	UserID          int    `json:"user_id" binding:"required"`
	TaskID          int    `json:"task_id" binding:"required"`
	Status          string `json:"status"`
	Task            string `json:"task"`
	TaskDescription string `json:"task_description"`
}

func TaskUpdate(r *gin.Context) {
	var task_params TaskUpdateParams
	if err := r.ShouldBindJSON(&task_params); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	task := models.Task{ID: task_params.TaskID}
	result := models.DB.Model(&task).Select("*").Updates(models.Task{
		ID:              task_params.TaskID,
		UserID:          task_params.UserID,
		Status:          task_params.Status,
		Task:            task_params.Task,
		TaskDescription: task_params.TaskDescription,
		UpdatedAt:       time.Now(),
	})

	models.DB.Preload("User").Find(&task)
	if result.Error != nil {
		// error handling...
		panic(result.Error)
	}

	r.JSON(http.StatusOK, gin.H{"data": task})
}
