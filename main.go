package main

import (
	"github.com/dekbadnerd/api-register-login/Controller/auth"
	"github.com/dekbadnerd/api-register-login/orm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
)

// Binding from JSON
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

func main() {
	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", auth.Register)
	r.Run("localhost:8080")
}
