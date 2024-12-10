package controllers

import (
	"flockstay_api/initializers"
	"flockstay_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ginCtx *gin.Context) {
	var body struct {
		Phone    string `json:"phone"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if ginCtx.BindJSON(&body) != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to hash password",
		})
		return
	}

	newUser := models.User{Phone: body.Phone, Username: body.Username, Password: string(hashPassword)}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to create user",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"msg": "successfully created user account",
	})
}
