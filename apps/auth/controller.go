package auth

import (
	models "Go-Rampup/db/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type UserAuthController struct {
	DB *gorm.DB
}

func (ctrl *UserAuthController) register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := ctrl.createUser(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	context.Set("email", user.Email)
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
	err = user.GetUser(ctrl.DB, [][]string{{"email = ?", loginPayload.Email}})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	ctrl.DB.Model(&user).Where("email = ?", loginPayload.Email).First(&user)
	if !VerifyPassword(loginPayload.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, nil)
		return
	}
	context.Set("email", user.Email)
	context.Next()
}

func (ctrl *UserAuthController) Detail(context *gin.Context) {
	var user UserDetails
	ctrl.DB.Model(&models.User{}).Where("email = ?", context.GetString("email")).First(&user)
	context.JSON(http.StatusOK, user)
}

func (ctrl *UserAuthController) UpdateUser(context *gin.Context) {
	var userData UserUpdatePayload
	if err := context.ShouldBindJSON(&userData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userUpdateValidator := validator.New(validator.WithRequiredStructEnabled())
	userUpdateValidator.Struct(userData)
	var user models.User
	ctrl.DB.Model(&user).Where("email = ?", context.GetString("email")).Updates(&userData)
	ctrl.DB.Model(&user).Where("email = ?", context.GetString("email")).First(&user)
	context.JSON(http.StatusOK, user)
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
	ctrl.DB.Model(&user).Where("email = ?", context.GetString("email")).First(&user)
	if !VerifyPassword(passwordUpdatePayload.OldPassword, user.Password) {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Old password does not match"})
		return
	}
	if newPassword, err := HashPassword(passwordUpdatePayload.NewPassword); err == nil {
		ctrl.DB.Model(&user).Where("email = ?", context.GetString("email")).Update("password", newPassword)
	}
	context.JSON(http.StatusOK, user)
}

func (ctrl *UserAuthController) DeleteUser(context *gin.Context) {
	var user models.User
	ctrl.DB.Model(&user).Where("email = ?", context.GetString("email")).First(&user)
	ctrl.DB.Delete(&user)
}
