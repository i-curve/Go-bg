package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int
	ModifiedOn int
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func CheckUser(username, password string) (bool, error) {
	var auth User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID <= 0 {
		return false, nil
	}
	return true, nil
}

func Register(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return false, err
	}
	err = db.Create(&User{Username: username, Password: password}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func IsExist(name string) bool {
	var user User

	err := db.Model(&user).Where("username=?", name).First(&user).Error
	if err != gorm.ErrRecordNotFound && user.ID > 0 {
		return true
	}
	return false
}
