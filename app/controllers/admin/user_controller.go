package admin

import (
	"messenger/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserList(request *gin.Context) {
	user := []models.User{}
	result := models.DB.Find(&user)
	println(result.RowsAffected)
	request.JSON(http.StatusOK, gin.H{"data": user})
}

func UserCreate(c *gin.Context) {
	lastname := c.Query("lastname")
	user := models.User{
		FirstName:  "Jophat",
		MiddleName: "Honor",
		LastName:   lastname,
		Email:      "jophat.tamayo@gmail.com",
	}
	models.DB.Create(&user)
	c.String(http.StatusOK, lastname)
}
