package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/gin-auth/db"
	"github.com/kmdkuk/gin-auth/handler"
	"github.com/kmdkuk/gin-auth/middleware"
)

func main() {
	if err := db.Init(); err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("MYSESSION", store))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowWildcard:   true,
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"*",
		},
	}))

	router.POST("/login", handler.Login)

	logout := router.Group("/logout")
	logout.Use(middleware.LoginCheckMiddleware())
	logout.POST("", handler.Logout)

	router.POST("/users", handler.CreateUser)
	user := router.Group("/users")
	user.Use(middleware.LoginCheckMiddleware())
	user.GET("", handler.GetCurrentUser)

	router.Run(":3000")
}
