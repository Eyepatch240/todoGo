package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"todoGo/database"
	"todoGo/models"
	"todoGo/security"
)

func Signup(user *models.User) error {

	result := database.DB.Where("email = ?", user.Email).First(&user)

	if result.RowsAffected > 0 {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func Login(user *models.User) ([]string, error) {

	var savedUser models.User

	result := database.DB.Where("email = ?", user.Email).First(&savedUser)

	if result.RowsAffected == 0 {
		return nil, errors.New("no user with such email exists")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	tokens := make([]string, 2)

	accessToken, err := security.GenerateToken(savedUser, "access")
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.GenerateToken(savedUser, "refresh")
	if err != nil {
		return nil, err
	}

	tokens[0] = accessToken
	tokens[1] = refreshToken

	var refreshTokenDB models.RefreshToken

	refreshTokenDB.Token = tokens[1]
	refreshTokenDB.UserID = savedUser.ID

	database.DB.Save(&refreshTokenDB)

	return tokens, nil
}
