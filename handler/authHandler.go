package handler

import (
	"camp-backend/initial"
	"camp-backend/types"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(c *gin.Context) {
	var request types.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	currentUser := new(types.TMember)
	response := new(types.LoginResponse)

	err := initial.Db.First(currentUser, "username = ?", request.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Code = types.WrongPassword

		response.Data.UserID = request.Username
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if currentUser.IsDeleted {
		response.Code = types.WrongPassword
		response.Data.UserID = currentUser.UserID
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if currentUser.Password != request.Password {
		response.Code = types.WrongPassword
		response.Data.UserID = currentUser.UserID
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	_, err = c.Cookie("camp-session")
	if err != nil {
		c.SetCookie("camp-session", currentUser.UserID, 3600, "/", "localhost", false, true)
	}
	response.Code = types.OK
	response.Data.UserID = currentUser.UserID
	c.JSON(http.StatusOK, response)
}

func Logout(c *gin.Context) {
	response := new(types.LogoutResponse)

	_, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = types.LoginRequired
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	c.SetCookie("camp-session", "", -1, "/", "localhost", false, true)
	response.Code = types.OK
	c.JSON(http.StatusOK, response)
}

func Whoami(c *gin.Context) {
	response := new(types.WhoAmIResponse)

	UserID, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = types.LoginRequired
		response.Data = types.TMember{}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	currentUser := new(types.TMember)
	if err = initial.Db.First(currentUser, UserID).Error; err != nil {
		response.Code = types.UnknownError
		response.Data = types.TMember{}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	response.Code = types.OK
	response.Data = *currentUser
	c.JSON(http.StatusOK, response)
}
