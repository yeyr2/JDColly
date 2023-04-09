package sql

import (
	"fmt"
	"reptile-test-go/cmd"
	"time"
)

func GetSearchByClaimsId(cl *cmd.Claims) *[]cmd.Search {
	var search []cmd.Search
	id := cl.Id

	db.Select("`key`,update_time").Where("user_id = ?", id).Limit(15).Find(&search)

	for i := range search {
		search[i].Time = time.Unix(search[i].UpdateTime, 0).Format("2006-01-02 15:04:05")
	}

	return &search
}

func SetSearch(key string, id int64) error {
	search := cmd.Search{
		Key:        key,
		Id:         id,
		UpdateTime: time.Now().Unix(),
	}

	result := db.Where("user_id = ? and `key` = ?", id, key).Find(&search)
	if result.RowsAffected != 0 {
		return UpdateSearch(&search)
	}

	search.CreateTime = time.Now().Unix()
	result = db.Create(&search)

	if result.RowsAffected == 0 {
		return fmt.Errorf("error : No search record")
	}

	return nil
}

func UpdateSearch(search *cmd.Search) error {
	search.UpdateTime = time.Now().Unix()

	result := db.Model(&cmd.Search{}).Where("`key` = ? and user_id = ?", search.Key, search.Id).Updates(&search)

	if result.RowsAffected == 0 {
		return fmt.Errorf("更新搜索记录失败")
	}

	return nil
}
