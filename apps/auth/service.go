package auth

import (
	"Go-Rampup/db/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (userAuthCtrl UserAuthController) createUser(user *models.User) error {
	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	result := userAuthCtrl.DB.Create(&user)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey){
		return errors.New("user with this email already exists")
	}
	return nil
}
