package sql

import "reptile-test-go/api/cmd"

func GetShopInfoByKey(key string, hots *[]*cmd.Hot) {
	db.Where("`key` = ?", key).Find(hots)
}

func AddShopInfo(hots *[]*cmd.Hot) {
	if len(*hots) == 0 {
		return
	}

	key := (*(*hots)[0]).Key

	result := db.Where("`key` = ?", key).Find(&cmd.Hot{})
	if result.RowsAffected == 0 {
		db.Create(hots)
	}
}
