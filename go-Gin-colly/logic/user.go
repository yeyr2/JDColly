package logic

import (
	"reptile-test-go/cmd"
	sql "reptile-test-go/model"
)

func ModifyUser(user *cmd.User, id int64) error {
	err := sql.ModifyUser(user, id)
	return err
}
