package sql

import (
	"fmt"
	"reptile-test-go/cmd"
)

func CheckLogin(username string, password string) (*cmd.User, error) {
	var user cmd.User

	result := db.Select("id").Where("username = ? and password = ?", username, password).Find(&user)

	if result.RowsAffected == 0 {
		return &user, fmt.Errorf("error: username does not exist")
	}

	return &user, nil
}

func CheckUserExist(username string) bool {
	var user cmd.User

	result := db.Select("id").Where("username = ?", username).Find(&user)
	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func CreateUser(username string, password string) (*cmd.User, error) {
	user := cmd.User{
		Username:    username,
		Password:    password,
		Sex:         "未知",
		PhoneNumber: "",
		Email:       "",
	}

	result := db.Create(&user)
	if result.RowsAffected == 0 {
		return &user, fmt.Errorf("error: User registration failed")
	}

	return &user, nil
}

func FindUserById(id int64) (user *cmd.User, err error) {

	result := db.Select("id,username,sex,phoneNumber,email,address,emergencyContact").Where("id = ?", id).Find(&user)
	if result.RowsAffected == 0 {
		return user, result.Error
	}
	return user, nil
}
