package socialize

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"Go-Rampup/constants"
)

type SocializeController struct {
	DB *gorm.DB
}

func (ctrl *SocializeController) FollowUsers(context *gin.Context) {
	var followPayload FollowPayload
	context.Bind(&followPayload)
	userIdsValidator := validator.New(validator.WithRequiredStructEnabled())
	err := userIdsValidator.Struct(followPayload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	status, createdFollowingUsersIds, err := ctrl.Follow(context.GetString(constants.GetConstants().ContextKeys.EMAIL), followPayload.UserIds)
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, gin.H{"details": fmt.Sprintf("Started following %v", createdFollowingUsersIds)})
}

func (ctrl *SocializeController) UnFollowUsers(context *gin.Context) {
	var unFollowPayload FollowPayload
	context.Bind(&unFollowPayload)
	userIdsValidator := validator.New(validator.WithRequiredStructEnabled())
	err := userIdsValidator.Struct(unFollowPayload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, unFollowedUsersIds, err := ctrl.UnFollow(context.GetString(constants.GetConstants().ContextKeys.EMAIL), unFollowPayload.UserIds)
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, gin.H{"details": fmt.Sprintf("UnFollowed %v", unFollowedUsersIds)})
}

func (ctrl *SocializeController) GetFollowers(context *gin.Context) {
	status, followers, err := ctrl.GetFollowersList(context.GetString(constants.GetConstants().ContextKeys.EMAIL))
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, followers)
}

func (ctrl *SocializeController) GetFollowings(context *gin.Context) {
	status, followings, err := ctrl.GetFollowingsList(context.GetString(constants.GetConstants().ContextKeys.EMAIL))
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, followings)
}
