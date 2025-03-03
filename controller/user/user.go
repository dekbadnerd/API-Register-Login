package user

import (
	"net/http"

	"github.com/dekbadnerd/api-register-login/orm"
	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"message": "Read all users",
		"users":   users,
	})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Read sucessfully",
		"users":   user,
	})
}