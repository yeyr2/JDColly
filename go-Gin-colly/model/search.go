package sql

import (
	"reptile-test-go/cmd"
	"time"
)

func GetSearchByClaimsId(cl *cmd.Claims) *[]cmd.Search {
	var search []cmd.Search
	id := cl.Id

	db.Select("`key`").Where("user_id = ?", id).Find(&search)

	return &search
}

func SetSearch(key string, id int64) error {
	search := cmd.Search{
		Key:        key,
		Id:         id,
		CreateTime: time.Now().Unix(),
	}

	result := db.Where("user_id = ? and `key` = ?", id, key).Find(&cmd.Search{})
	if result.RowsAffected != 0 {
		return nil
	}

	result = db.Create(&search)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
