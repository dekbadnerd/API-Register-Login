package main

import (
	"github.com/dekbadnerd/api-register-login/controller/auth"
	"github.com/dekbadnerd/api-register-login/controller/user"
	"github.com/dekbadnerd/api-register-login/middleware"
	"github.com/dekbadnerd/api-register-login/orm"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	authGroup := r.Group("/users", middleware.JWTAuthen())
	authGroup.GET("/readall", user.ReadAll)
	authGroup.GET("/profile", user.Profile)

	r.Run("localhost:8080")
}
