package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/cmd"
	"reptile-test-go/logic"
	"reptile-test-go/model"
)

func SearchInfo(c *gin.Context) {
	token, _ := c.Cookie("token")

	cl, err := logic.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	search := sql.GetSearchByClaimsId(cl)

	c.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		Value:      search,
	})
}
