package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New(`record not found`)
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

type UserRepo struct {
	*gorm.DB
}

func (u *UserRepo) GetUserByEmail(email string) (user *User, err error) {
	err = u.DB.Where("email = ?", email).First(&user).Error
	// if err not found, return ErrUserNotFound
	return
}

func (u *UserRepo) InsertUser(email, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: log
		return err
	}
	user := &User{
		Email:        email,
		PasswordHash: string(hashedPass),
	}
	err = u.DB.Create(user).Error
	if err != nil {
		// TODO: log errors.New("user already exists")
		return err
	}
	return err
}
