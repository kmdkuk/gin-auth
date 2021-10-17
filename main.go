package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/gin-auth/db"
	"github.com/kmdkuk/gin-auth/handler"
	"github.com/kmdkuk/gin-auth/middleware"
	"github.com/kmdkuk/gin-auth/repository"
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

	userHandler := handler.NewUserHandler(repository.NewUserRepository(db.GetDB()))

	router.POST("/login", userHandler.Login)

	logout := router.Group("/logout")
	logout.Use(middleware.LoginCheckMiddleware())
	logout.POST("", userHandler.Logout)

	router.POST("/users", userHandler.CreateUser)
	user := router.Group("/users")
	user.Use(middleware.LoginCheckMiddleware())
	user.GET("", userHandler.GetCurrentUser)

	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}
