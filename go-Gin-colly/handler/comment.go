package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/cmd"
	"reptile-test-go/logic"
	"strconv"
	"strings"
)

func GetComment(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 0, 64)
	productId := c.Query("productId")
	isColly, _ := strconv.Atoi(c.Query("isColly"))
	startTime, _ := strconv.ParseInt(c.Query("startTime"), 0, 64)
	lastTime, _ := strconv.ParseInt(c.Query("lastTime"), 0, 64)
	//token, _ := c.Cookie("token")
	token := c.Query("token")
	logic.Trim(&token)

	if strings.TrimSpace(productId) == "" {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "商品id为空",
		})
		return
	}

	cl, err := logic.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if cl.Id != id {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息错误",
		})
		return
	}

	var analyze cmd.AnalyzeComment
	var jdComment cmd.JDComment
	comments := new([]cmd.Comments)

	if isColly != 1 {
		// 获取评论
		logic.GetCommentById(productId, startTime, lastTime, &jdComment)
		//fmt.Println(len(jdComment.Comments))
		comments = &jdComment.Comments
	} else {
		// 从数据库中获取数据
		comments = logic.GetCommentBySql(productId, startTime, lastTime)
	}
	analyze.Count = int32(len(*comments))

	// 分析获取的评价(总评价,评价区间)
	flag := logic.AnalyzeGetComments(comments, &analyze, "Chinese NLP")
	if !flag {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "没有评论",
		})
		return
	}

	// 获取词云分析
	logic.WordCloudAnalysis(comments, &analyze, productId)

	//setHeader(c)

	c.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      analyze,
	})
}
