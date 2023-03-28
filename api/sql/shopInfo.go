package sql

import "reptile-test-go/api/cmd"

func GetShopInfoByKey(key string, hots *[]*cmd.Hot) {
	db.Where("`key` = ?", key).Find(hots)
}

func AddShopInfo(hots *[]*cmd.Hot) {
	db.Create(hots)
}
