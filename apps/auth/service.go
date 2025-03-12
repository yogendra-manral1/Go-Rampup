package auth

import (
	"Go-Rampup/db/models"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (userAuthCtrl UserAuthController) createUser(user *models.User) (int, error) {
	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	result := userAuthCtrl.DB.Create(&user)
	// if errors.Is(result.Error, gorm.ErrDuplicatedKey){
	// 	return errors.New("user with this email already exists")
	// }

	// Using error not nil as there is some bug in errors.Is
	if result.Error != nil {
		log.Println("Error in creating user ", result.Error.Error())
		return http.StatusBadRequest, errors.New("user with this email already exists")
	}
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) LoginUser(user *models.User, loginPayload UserLoginPayload) (int, error) {
	err := user.GetUser(userAuthCtrl.DB, [][]string{{"email = ?", loginPayload.Email}})
	if err != nil {
		return http.StatusNotFound, err
	}
	userAuthCtrl.DB.Model(&user).Where("email = ?", loginPayload.Email).First(&user)
	if !VerifyPassword(loginPayload.Password, user.Password) {
		return http.StatusUnauthorized, errors.New("incorrect Password")
	}
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) GetUserDetail(user *UserDetails, email string) (int, error) {
	result := userAuthCtrl.DB.Model(&models.User{}).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return http.StatusNotFound, result.Error
	}
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) UpdateUserDetail(updatedData *UserUpdatePayload, user *UserDetails, email string) (int, error) {
	result := userAuthCtrl.DB.Model(&user).Where("email = ?", email).Updates(&updatedData)
	if result.Error != nil {
		return http.StatusNotFound, errors.New("user does not exists")
	}
	userAuthCtrl.DB.Model(&user).Where("email = ?", email).First(&user)
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) UpdateUserPassword(user *models.User, email string, passwordUpdatePayload PasswordUpdatePayload) (int, error) {
	userAuthCtrl.DB.Model(&user).Where("email = ?", email).First(&user)
	if !VerifyPassword(passwordUpdatePayload.OldPassword, user.Password) {
		return http.StatusBadRequest, errors.New("old password does not match")
	}
	if newPassword, err := HashPassword(passwordUpdatePayload.NewPassword); err == nil {
		userAuthCtrl.DB.Model(&user).Where("email = ?", email).Update("password", newPassword)
	} else {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) DeleteUserDetail(email string) (int, error) {
	result := userAuthCtrl.DB.Where("email = ?", email).Delete(&models.User{})
	if result.Error != nil {
		return http.StatusNotFound, errors.New("user does not exist")
	}
	return http.StatusOK, nil
}

func (userAuthCtrl UserAuthController) GetUsersList(users []UsersListItem) (int, error) {
	result := userAuthCtrl.DB.Model(&models.User{}).Select("id", "email").Find(&users)
	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}
	return http.StatusOK, nil
}
