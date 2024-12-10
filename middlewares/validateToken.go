package middlewares

import (
	"flockstay_api/initializers"
	"flockstay_api/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(ginCtx *gin.Context) {
	authHeader := ginCtx.GetHeader("Authorization")

	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Unauthorized",
		})
		return
	}

	tokenStr := authHeader[7:]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Fatal(err)
		ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Unauthorized",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var loggedInUser models.User
		if err := initializers.DB.Where("id = ?", claims["uid"]).First(&loggedInUser).Error; err != nil {
			ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}

		ginCtx.Set("user", loggedInUser)
	} else {
		ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "Unauthorized",
		})
	}
}
