package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aafs20/rakamin_golang/database"
	"github.com/aafs20/rakamin_golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//Get cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Silahkan Login Terlebih Dahulu",
		})
		return
	}
	//Ekstrak validasi kode
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Mencari Token user
		var user models.User
		database.DB.First(&user, claims["sub"])

		if user.Id_u == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Melakukan SET untuk Request user
		c.Set("saya", user.Id_u)
		c.Set("user", user)
		//Continue
		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
