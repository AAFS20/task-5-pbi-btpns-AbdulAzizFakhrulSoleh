package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/aafs20/rakamin_golang/database"
	"github.com/aafs20/rakamin_golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string
}

func LoginUser(c *gin.Context) {
	var body struct {
		Id_u     *int    `json:"id_u"`
		Email    *string `json:"email"`
		Username *string `json:"username"`
		Password string  `json:"password"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	//Look up requested user
	var user models.User
	database.DB.Table("users").Where("username = ?", body.Username).Find(&user)
	if user.Id_u == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Username",
		})
		return
	}
	//Compare sent in pass with saved user pass hash memband
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":     "Invalid password",
			"data":      err.Error(),
			"Password":  body.Password,
			"Password2": user.Password,
		})
		return
	}
	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id_u,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		//"exp": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid create token",
		})
		return
	}
	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	//Mengambil data set user dari middleware
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcom User",
		"Data":    user,
	})
}
