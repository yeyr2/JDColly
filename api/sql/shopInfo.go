package sql

import "reptile-test-go/api/cmd"

func GetShopInfoByKey(key string, hots *[]*cmd.Hot) {
	db.Where("`key` = ?", key).Find(hots)
}

func AddShopInfo(hots *[]*cmd.Hot, key string) {
	if len(*hots) == 0 {
		return
	}

	result := db.Where("`key` = ?", key).Find(&cmd.Hot{})
	if result.RowsAffected == 0 {
		db.Create(hots)
	}
}
