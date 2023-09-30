package photos_controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aafs20/rakamin_golang/database"
	"github.com/aafs20/rakamin_golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllPhotos(ctx *gin.Context) {
	photo := new([]models.Photos)
	user, _ := ctx.Get("saya")
	//database.DB.Find(&users)
	err := database.DB.Table("photos").Where("user_id = ?", user).Find(&photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"Message": "internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"Message": "Data trasmitted.",
		"Data":    photo,
	})
}

func InsertPhoto(ctx *gin.Context) {
	user, _ := ctx.Get("saya")
	//Get the email/pass off req body
	var bodys struct {
		Title     *string `json:"title"`
		Photo_url string  `json:"photoUrl"`
		Caption   string  `json:"caption"`
		User_id   string  `json:"userid"`
	}
	user_id := fmt.Sprintf("%d", user)
	if ctx.Bind(&bodys) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}
	bodys.User_id = user_id
	// Menambahkan Foto dan Kolom update User
	result := database.DB.Table("photos").Create(&bodys)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed create photo",
		})
	}
	userHasil := models.User{UpdatedAt: time.Now()}
	errUpdate := database.DB.Table("users").Where("id_u = ?", user_id).Updates(&userHasil).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed update user",
		})
	}
	//Respond
	ctx.JSON(http.StatusOK, gin.H{
		"Data":    bodys,
		"User":    bodys.User_id,
		"Message": "Success Data Transmitted",
	})
}

func UpdatePhoto(ctx *gin.Context) {
	user, _ := ctx.Get("saya")
	var photo struct {
		Title     *string `json:"title"`
		Photo_url string  `json:"photoUrl"`
		Caption   string  `json:"caption"`
		User_id   string  `json:"userid"`
	}
	user_id := fmt.Sprintf("%d", user)
	idp := ctx.Param("photoId")

	if errReq := ctx.ShouldBind(&photo); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	photo.User_id = user_id
	photos := new(models.Photos)
	errDb := database.DB.Table("photos").Where("idp = ?", idp).Find(&photos).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
		})
	}
	if photos.Idp == nil {
		ctx.JSON(404, gin.H{
			"message": "Data not found.",
		})
	}

	errUpdatePhoto := database.DB.Table("photos").Where("idp = ?", idp).Updates(&photo).Error
	if errUpdatePhoto != nil {
		ctx.JSON(500, gin.H{
			"message": "Can't Update Data.",
		})
		return
	}
	userHasil := models.User{UpdatedAt: time.Now()}
	errUpdateUser := database.DB.Table("users").Where("id_u = ?", user_id).Updates(&userHasil).Error
	if errUpdateUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed update user",
		})
	}

	ctx.JSON(200, gin.H{
		"message":  "Data Succesfully.",
		"Data":     photo,
		"Id Photo": photos.Idp,
		"Id User":  photo.User_id,
	})
}
func DeletePhoto(ctx *gin.Context) {
	user, _ := ctx.Get("saya")
	user_id := fmt.Sprintf("%d", user)
	idp := ctx.Param("photoId")
	photo := new(models.Photos)
	errFind := database.DB.Table("photos").Where("idp = ?", idp).Find(&photo).Error
	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
		})
		return
	}
	if photo.Idp == nil {
		ctx.JSON(404, gin.H{
			"message": "Data not Found.",
		})
		return
	}

	errDb := database.DB.Table("photos").Unscoped().Where("idp = ?", idp).Delete(&models.Photos{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error.",
			"error":   errDb.Error(),
		})
		return
	}
	userHasil := models.User{UpdatedAt: time.Now()}
	errUpdateUser := database.DB.Table("users").Where("id_u = ?", user_id).Updates(&userHasil).Error
	if errUpdateUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed update user",
		})
	}

	ctx.JSON(200, gin.H{
		"message": "data deleted successfully.",
	})
}
