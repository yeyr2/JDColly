package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/cmd"
	"reptile-test-go/logic"
	"strconv"
)

func GetComment(c *gin.Context) {
	id := c.Query("id")
	isColly, _ := strconv.Atoi(c.Query("isColly"))
	lastTime, _ := strconv.ParseInt(c.Query("lastTime"), 0, 64)
	token, _ := c.Cookie("token")

	if _, err := logic.ParseToken(token); err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	var analyze cmd.AnalyzeComment
	var jdComment cmd.JDComment
	comments := new([]cmd.Comments)

	if isColly != 1 {
		// 获取评论
		logic.GetCommentById(id, lastTime, &jdComment)
		//fmt.Println(len(jdComment.Comments))
		comments = &jdComment.Comments
	} else {
		// 从数据库中获取数据
		comments = logic.GetCommentBySql(id, lastTime)
	}

	// 分析获取的评价(总评价,评价区间)
	flag := logic.AnalyzeGetComments(comments, &analyze)
	if !flag {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "没有评论",
		})
		return
	}

	// 获取词云分析
	logic.WordCloudAnalysis(comments, &analyze, id)

	//setHeader(c)

	c.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      analyze,
	})
}
