package auth

import (
	models "Go-Rampup/db/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"Go-Rampup/constants"
)

type UserAuthController struct {
	DB *gorm.DB
}

func (ctrl *UserAuthController) register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	status, err := ctrl.createUser(&user)
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.Set(constants.GetConstants().ContextKeys.EMAIL, user.Email)
}

func (ctrl *UserAuthController) Login(context *gin.Context) {
	var loginPayload UserLoginPayload
	context.Bind(&loginPayload)
	userValidator := validator.New(validator.WithRequiredStructEnabled())
	err := userValidator.Struct(loginPayload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	status, err := ctrl.LoginUser(&user, loginPayload)
	if err != nil{
		context.AbortWithStatusJSON(status, err.Error())
	}
	context.Set(constants.GetConstants().ContextKeys.EMAIL, user.Email)
}

func (ctrl *UserAuthController) Detail(context *gin.Context) {
	var user UserDetails
	status, err := ctrl.GetUserDetail(&user, context.GetString(constants.GetConstants().ContextKeys.EMAIL))
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, user)
}

func (ctrl *UserAuthController) UpdateUser(context *gin.Context) {
	var userData UserUpdatePayload
	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userUpdateValidator := validator.New(validator.WithRequiredStructEnabled())
	userUpdateValidator.Struct(userData)
	var user UserDetails
	status, err := ctrl.UpdateUserDetail(&userData, &user, context.GetString(constants.GetConstants().ContextKeys.EMAIL))
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, user)
}

func (ctrl *UserAuthController) UpdatePassword(context *gin.Context) {
	var passwordUpdatePayload PasswordUpdatePayload
	if err := context.ShouldBindJSON(&passwordUpdatePayload); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	passwordUpdateValidator := validator.New(validator.WithRequiredStructEnabled())
	err := passwordUpdateValidator.Struct(passwordUpdatePayload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	status, err := ctrl.UpdateUserPassword(&user, context.GetString(constants.GetConstants().ContextKeys.EMAIL), passwordUpdatePayload)
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, user)
}

func (ctrl *UserAuthController) DeleteUser(context *gin.Context) {
	status, err := ctrl.DeleteUserDetail(context.GetString(constants.GetConstants().ContextKeys.EMAIL))
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, "User Deleted Successfully")
}

func (ctrl *UserAuthController) GetAllUsers(context *gin.Context){
	var users []UsersListItem
	status, err := ctrl.GetUsersList(users)
	if err != nil {
		context.AbortWithStatusJSON(status, err.Error())
		return
	}
	context.JSON(status, users)
}
