package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hagios2/gopharma/services"
	"net/http"
)

func Login(ctx *gin.Context) {
	var credentials services.UserCredentials
	if err := ctx.ShouldBindJSON(credentials); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return
	}

	return services.LoginService.Login(credentials.Email, credentials.Password)
}
