package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/gin-auth/constants"
	"github.com/kmdkuk/gin-auth/model"
	"github.com/kmdkuk/gin-auth/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var request model.User
	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println("error bind json", err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByID(request.ID)
	if err != nil {
		fmt.Println("invalid userid")
		c.Status(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("invalid password")
		c.Status(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	session.Set(constants.USERID_KEY, user.ID)
	session.Save()
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusOK)
}

func CreateUser(c *gin.Context) {
	var request model.User
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = repository.CreateUser(request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func GetCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	loginUserID := session.Get(constants.USERID_KEY)

	user, err := repository.GetUserByID(loginUserID.(string))
	if err != nil {
		fmt.Println("user not found ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
