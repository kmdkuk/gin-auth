package handler

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kmdkuk/gin-auth/constants"
	"github.com/kmdkuk/gin-auth/model"
	"github.com/kmdkuk/gin-auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	CreateUser(c *gin.Context)
	GetCurrentUser(c *gin.Context)
}

type userHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) UserHandler {
	return &userHandler{
		userRepository: userRepository,
	}
}

func (u *userHandler) Login(c *gin.Context) {
	var request model.User
	err := c.BindJSON(&request)
	if err != nil {
		log.Println("error bind json", err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := u.userRepository.GetUserByID(request.UserID)
	if err != nil {
		log.Println("invalid userid")
		c.Status(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("invalid password")
		c.Status(http.StatusBadRequest)
		return
	}

	session := sessions.Default(c)
	session.Set(constants.USERID_KEY, user.UserID)
	if err := session.Save(); err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (u *userHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (u *userHandler) CreateUser(c *gin.Context) {
	var request model.User
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err = u.userRepository.CreateUser(request)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func (u *userHandler) GetCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	loginUserID := session.Get(constants.USERID_KEY)

	user, err := u.userRepository.GetUserByID(loginUserID.(string))
	if err != nil {
		log.Println("user not found ", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
