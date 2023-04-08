package sql

import (
	"fmt"
	"reptile-test-go/cmd"
	"time"
)

func GetSearchByClaimsId(cl *cmd.Claims) *[]cmd.Search {
	var search []cmd.Search
	id := cl.Id

	db.Select("`key`,create_time").Where("user_id = ?", id).Limit(15).Find(&search)

	return &search
}

func SetSearch(key string, id int64) error {
	search := cmd.Search{
		Key:        key,
		Id:         id,
		CreateTime: time.Now().Unix(),
	}

	result := db.Where("user_id = ? and `key` = ?", id, key).Find(&search)
	if result.RowsAffected != 0 {
		return UpdateSearch(key, id, &search)
	}

	result = db.Create(&search)

	if result.RowsAffected == 0 {
		return fmt.Errorf("error : No search record")
	}

	return nil
}

func UpdateSearch(key string, id int64, search *cmd.Search) error {
	search.UpdateTime = time.Now().Unix()

	//result := db.Model()

	return nil
}
