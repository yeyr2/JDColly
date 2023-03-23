package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetComment(c *gin.Context) {
	id := c.Query("id")
	isColly, _ := strconv.Atoi(c.Query("isColly"))
	lastTime, _ := strconv.ParseInt(c.Query("lastTime"), 0, 64)

	var analyze AnalyzeComment
	var comment JDComment

	if isColly != 1 {
		// 获取评论
		GetCommentById(id, lastTime, &comment)
		//fmt.Println(len(comment.Comments))

		// 分析获取的评价(总评价,评价区间)
		AnalyzeGetComments(&comment, &analyze)

		// 获取词云分析
		WordCloudAnalysis(&comment, &analyze)
	} else {
		// 从数据库中获取数据
		GetWordCloudAndAnalyzeRating(&analyze, id, lastTime)
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      analyze,
	})
}
