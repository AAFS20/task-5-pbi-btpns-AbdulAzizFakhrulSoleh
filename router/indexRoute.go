package router

import (
	"github.com/aafs20/rakamin_golang/controllers"
	"github.com/aafs20/rakamin_golang/controllers/photos_controller"
	"github.com/aafs20/rakamin_golang/controllers/user_controller"
	"github.com/aafs20/rakamin_golang/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	//Route User Login
	route.POST("/login", controllers.LoginUser) //Untuk melakukan login user

	//Route User
	route.GET("/user", user_controller.GetAllUser)        //Menampilkan Semua user
	route.POST("/user", user_controller.InsertUser)       //Menginputkan data user baru
	route.GET("/user/:id", user_controller.GetById)       //Menampilkan user berdasarkan id user
	route.PATCH("/user/:id", user_controller.UpdateUser)  //Meng-edit user berdasarkan id user
	route.DELETE("/user/:id", user_controller.DeleteUser) //Menghapus user berdasarkan id user

	//Route Validasi
	route.GET("/validate", middleware.RequireAuth, controllers.Validate) //melakukan validasi setelah login

	//Route Photos
	route.GET("/photos", middleware.RequireAuth, photos_controller.GetAllPhotos)            //melakukan validasi pada saat Menampilkan semua foto pada user
	route.POST("/photos", middleware.RequireAuth, photos_controller.InsertPhoto)            //melakukan validasi pada saat Menambahkan data foto oleh user
	route.PATCH("/photos/:photoId", middleware.RequireAuth, photos_controller.UpdatePhoto)  //melakukan validasi pada saat Meng-edit data foto pada user id photo
	route.DELETE("/photos/:photoId", middleware.RequireAuth, photos_controller.DeletePhoto) //melakukan validasi pada saat Menghapus foto pada user berdasarkan id photo
}
