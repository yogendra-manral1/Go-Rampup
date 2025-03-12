package auth

import (
	"github.com/gin-gonic/gin"
	"Go-Rampup/middlewares"
)

func (ctrl *UserAuthController) SetRoutes(r *gin.RouterGroup){
	r.GET("/", middlewares.Authorization, ctrl.Detail)
	r.GET("/list/", middlewares.Authorization, ctrl.GetAllUsers)
	r.DELETE("/", middlewares.Authorization, ctrl.DeleteUser)
	r.PATCH("/update/", middlewares.Authorization, ctrl.UpdateUser)
	r.POST("/login/", ctrl.Login, middlewares.CreateToken)
	r.POST("/register/", ctrl.register, middlewares.CreateToken)
	r.POST("/update-password/", middlewares.Authorization, ctrl.UpdatePassword)
}
