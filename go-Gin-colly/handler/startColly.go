package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reptile-test-go/cmd"
	"reptile-test-go/logic"
	sql "reptile-test-go/model"
	"strconv"
	"strings"
)

func StartColly(con *gin.Context) {
	key := con.Query("key")
	isColly, _ := strconv.Atoi(con.Query("isColly"))
	token, _ := con.Cookie("token")
	cl, err := logic.ParseToken(token)
	log.Println(token)

	if err != nil {
		con.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if strings.TrimSpace(key) == "" {
		con.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "key不能为空",
		})
		return
	}

	var hots []*cmd.Hot
	if isColly == 0 {
		logic.GetInfoByJDKey(key, &hots)
	} else if isColly == 1 {
		logic.GetInfoByJDKeyBySql(key, &hots)
	} else {
		con.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "isColly error",
		})
		return
	}
	err = sql.SetSearch(key, cl.Id)
	str := ""
	if err != nil {
		str = err.Error()
	}

	con.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  str,
		Value:      hots,
	})
}
