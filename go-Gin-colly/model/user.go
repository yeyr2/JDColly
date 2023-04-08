package sql

import (
	"fmt"
	"reptile-test-go/cmd"
)

//username = phone_number

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
	var nickname string
	length := len(username)
	if length > 4 {
		nickname = "用户" + username[length-4:]
	} else {
		nickname = "用户" + username
	}
	user := cmd.User{
		Username:    username,
		Nickname:    nickname,
		Password:    password,
		Sex:         "未知",
		PhoneNumber: username,
		Email:       "",
	}

	result := db.Create(&user)
	if result.RowsAffected == 0 {
		return &user, fmt.Errorf("error: User registration failed")
	}

	return &user, nil
}

func FindUserById(id int64) (user *cmd.User, err error) {

	result := db.Select("id,username,nickname,sex,phone_number,email,address,emergency_contact").Where("id = ?", id).Find(&user)
	if result.RowsAffected == 0 {
		return user, fmt.Errorf("the user cannot be found")
	}
	return user, nil
}

func ModifyUser(user *cmd.User, id int64) error {
	result := db.Where("id = ?", id).Updates(&user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("the user cannot be found")
	}
	return nil
}
