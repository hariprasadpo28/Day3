package Routes

import (
	"github.com/gin-gonic/gin"
	"project/Controller"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-data")
	{
		grp1.GET("student", Controller.GetUsers)
		grp1.POST("student", Controller.CreateUser)
		grp1.GET("student/:id", Controller.GetUserByID)
		grp1.PUT("student/:id", Controller.UpdateUser)
		grp1.DELETE("student/:id", Controller.DeleteUser)
		grp1.PATCH("student/:id", Controller.PatchUser)
	}
	return r
}