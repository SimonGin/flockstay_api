package controllers

import (
	"flockstay_api/initializers"
	"flockstay_api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(ginCtx *gin.Context) {
	var reqBody struct {
		Phone    string `json:"phone"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if ginCtx.BindJSON(&reqBody) != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}

	var existingUser models.User
	if err := initializers.DB.Where("phone = ?", reqBody.Phone).First(&existingUser).Error; err == nil {
		ginCtx.JSON(http.StatusConflict, gin.H{
			"msg": "Phone number already exists",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to hash password",
		})
		return
	}

	newUser := models.User{Phone: reqBody.Phone, Username: reqBody.Username, Password: string(hashPassword)}
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

func Login(ginCtx *gin.Context) {
	var reqBody struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	// Bind the request body
	if err := ginCtx.BindJSON(&reqBody); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}
	// Find the user by phone
	var existingUser models.User
	if err := initializers.DB.Where("phone = ?", reqBody.Phone).First(&existingUser).Error; err != nil {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"msg": "Invalid phone number",
		})
		return
	}
	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(reqBody.Password)); err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Invalid password",
		})
		return
	}
	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": existingUser.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	// Sign the token with a secret
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to generate token",
		})
		return
	}
	// Return the access token
	ginCtx.JSON(http.StatusOK, gin.H{
		"access_token": tokenStr,
	})
}
