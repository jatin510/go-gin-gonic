package controllers

import (
	"net/http"

	"gin-gonic-api/forms"
	"gin-gonic-api/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u UserController) SignUpUser(c *gin.Context) {
	userSignUpForm := forms.UserSignUp{}

	if err := c.Bind(&userSignUpForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user, err := userModel.Create(userSignUpForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}
