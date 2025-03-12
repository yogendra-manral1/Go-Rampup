package socialize

import (
	"github.com/gin-gonic/gin"
	"Go-Rampup/middlewares"
)

func (ctrl *SocializeController) SetRoutes(r *gin.RouterGroup){
	r.POST("/follow/", middlewares.Authorization, ctrl.FollowUsers)
	r.POST("/un-follow/", middlewares.Authorization, ctrl.UnFollowUsers)
	r.GET("/followers", middlewares.Authorization, ctrl.GetFollowers)
	r.GET("/followings", middlewares.Authorization, ctrl.GetFollowings)
}
