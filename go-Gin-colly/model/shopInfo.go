package sql

import (
	"reptile-test-go/struct"
)

func GetShopInfoByKey(key string, hots *[]*_struct.Hot) {
	db.Where("`key` = ?", key).Order("id ASC").Find(hots)
}

func AddShopInfo(hots *[]*_struct.Hot, key string) {
	if len(*hots) == 0 {
		return
	}

	result := db.Where("`key` = ?", key).Find(&_struct.Hot{})
	if result.RowsAffected == 0 {
		db.Create(hots)
	}
}
