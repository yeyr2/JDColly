package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/logic"
	sql "reptile-test-go/model"
	"reptile-test-go/struct"
	"strconv"
	"strings"
)

func StartColly(con *gin.Context) {
	id, _ := strconv.ParseInt(con.Query("id"), 0, 64)
	key := con.Query("key")
	isColly, _ := strconv.Atoi(con.Query("isColly"))
	//token, _ := con.Cookie("token")
	token := con.Query("token")
	logic.Trim(&token)
	cl, err := logic.ParseToken(token)

	if err != nil {
		con.JSON(http.StatusOK, _struct.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if cl.Id != id {
		con.JSON(http.StatusOK, _struct.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息错误",
		})
		return
	}

	if strings.TrimSpace(key) == "" {
		con.JSON(http.StatusOK, _struct.Response{
			StatusCode: 1,
			StatusMsg:  "key不能为空",
		})
		return
	}

	var hots []*_struct.Hot
	if isColly == 0 {
		logic.GetInfoByJDKey(key, &hots)
	} else if isColly == 1 {
		logic.GetInfoByJDKeyBySql(key, &hots)
	} else {
		con.JSON(http.StatusOK, _struct.Response{
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

	con.JSON(http.StatusOK, _struct.Response{
		StatusCode: 0,
		StatusMsg:  str,
		Value:      hots,
	})
}
