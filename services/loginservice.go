package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hagios2/gopharma/models"
	"github.com/hagios2/gopharma/util"
	"net/http"
)

type LoginService interface {
	Login(email string, password string)
	//Logout(ctx gin.Context)
}

type UserCredentials struct {
	Email    string `form:"user" json:"user" xml:"user"  binding:"required,email"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (credential UserCredentials) Login(email string, password string) {
	var ctx *gin.Context
	password, err := util.HashPassword(password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	user := models.User{Email: email, Password: password}
	util.DB.Where("email = ? and password = ?").First(&user)
	jwtService := NewJWTService()
	token := jwtService.GenerateToken(user.Email, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
	return
}

func Logout(ctx *gin.Context) {

}
