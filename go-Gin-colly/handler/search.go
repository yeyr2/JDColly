package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/logic"
	"reptile-test-go/model"
	"reptile-test-go/struct"
	"strconv"
)

func SearchInfo(c *gin.Context) {
	//token, _ := c.Cookie("token")
	id, _ := strconv.ParseInt(c.Query("id"), 0, 64)
	token := c.Query("token")
	logic.Trim(&token)

	cl, err := logic.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, _struct.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if cl.Id != id {
		c.JSON(http.StatusOK, _struct.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息错误",
		})
		return
	}

	search := sql.GetSearchByClaimsId(cl)

	c.JSON(http.StatusOK, _struct.Response{
		StatusCode: 0,
		Value:      search,
	})
}
