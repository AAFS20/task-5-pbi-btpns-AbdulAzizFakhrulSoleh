package user_controller

import (
	"net/http"
	"time"

	"github.com/aafs20/rakamin_golang/database"
	"github.com/aafs20/rakamin_golang/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)
	//database.DB.Find(&users)
	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"Data": users,
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	errDb := database.DB.Table("users").Where("id_u = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		return
	}
	if user.Id_u == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found.",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"Message": "Data trasmitted.",
		"data":    user,
	})
}

func InsertUser(ctx *gin.Context) {
	//Get username, email dan passwd
	var bodys struct {
		Username string
		Email    string
		Password string
	}
	if ctx.Bind(&bodys) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}
	//Pengecekan email dan username unique
	userExist := new(models.User)
	resultUsername := database.DB.Where("username = ?", bodys.Username).First(&userExist)
	resultEmail := database.DB.Where("email = ?", bodys.Email).First(&userExist)

	if resultUsername.Error != gorm.ErrRecordNotFound || resultEmail.Error != gorm.ErrRecordNotFound {
		if resultUsername.Error != gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{
				"Error": "Username Sudah Digunakan",
			})
		}
		if resultEmail.Error != gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{
				"Error": "Email Sudah digunakan",
			})
		}
		return
	}
	//end email

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(bodys.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}

	//Create the user
	user := models.User{Username: &bodys.Username, Email: &bodys.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed create user",
		})
	}

	//Respond
	ctx.JSON(http.StatusOK, gin.H{
		"Data": user,
	})
}

func UpdateUser(ctx *gin.Context) {
	var userReq struct {
		Username string `json:"username" form:"username"  binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
	}
	id := ctx.Param("id")
	user := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("users").Where("id_u = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
		})
	}
	if user.Id_u == 0 {
		ctx.JSON(404, gin.H{
			"message": "Data not found.",
		})
	}

	//

	userExist := new(models.User)
	resultUsername := database.DB.Where("username = ?", userReq.Username).First(&userExist)
	resultEmail := database.DB.Where("email = ?", userReq.Email).First(&userExist)

	if resultUsername.Error != gorm.ErrRecordNotFound || resultEmail.Error != gorm.ErrRecordNotFound {
		if resultUsername.Error != gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{
				"Error": "Username Sudah Digunakan",
				"data":  userReq.Email,
			})
		}
		if resultEmail.Error != gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{
				"Error": "Email Sudah digunakan",
			})
		}
		return
	}

	//

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}

	//
	userHasil := models.User{Username: &userReq.Username, Email: &userReq.Email, Password: string(hash), UpdatedAt: time.Now()}
	errUpdate := database.DB.Table("users").Where("id_u = ?", id).Updates(&userHasil).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "Can't Update Data.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":   "Data Succesfully.",
		"data":      userHasil,
		"Email":     userReq.Email,
		"useername": userReq.Username,
	})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	errFind := database.DB.Table("users").Where("id_u = ?", id).Find(&user).Error
	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
		})
		return
	}
	if user.Id_u == 0 {
		ctx.JSON(404, gin.H{
			"message": "Data not Found.",
		})
		return
	}
	errDb := database.DB.Table("users").Unscoped().Where("id_u = ?", id).Delete(&models.User{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
			"error":   errDb.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data deleted successfully.",
	})
}
