package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/api"
	"reptile-test-go/api/cmd"
	"strconv"
	"strings"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

func StartColly(con *gin.Context) {
	key := con.Query("key")
	isColly, _ := strconv.Atoi(con.Query("isColly"))

	if strings.TrimSpace(key) == "" {
		con.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "key不能为空",
			Value:      nil,
		})
		return
	}

	var hots []*cmd.Hot
	if isColly == 0 {
		api.GetInfoByJDKey(key, &hots)
	} else if isColly == 1 {
		api.GetInfoByJDKeyBySql(key, &hots)
	} else {
		con.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "isColly err",
		})
		return
	}

	//setHeader(con)

	con.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      hots,
	})
}
