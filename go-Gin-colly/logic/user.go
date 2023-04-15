package logic

import (
	sql "reptile-test-go/model"
	"reptile-test-go/struct"
)

func ModifyUser(user *_struct.User, id int64) error {
	err := sql.ModifyUser(user, id)
	return err
}
