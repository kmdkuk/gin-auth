package middleware

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/gin-auth/constants"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUserID := session.Get(constants.USERID_KEY)
		log.Println(loginUserID)

		if loginUserID == nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
