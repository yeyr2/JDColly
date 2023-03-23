package request

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reptile-test-go/api"
	"reptile-test-go/api/cmd"
	"strconv"
)

func GetComment(c *gin.Context) {
	id := c.Query("id")
	isColly, _ := strconv.Atoi(c.Query("isColly"))
	lastTime, _ := strconv.ParseInt(c.Query("lastTime"), 0, 64)

	var analyze cmd.AnalyzeComment
	var comment cmd.JDComment

	if isColly != 1 {
		// 获取评论
		api.GetCommentById(id, lastTime, &comment)
		//fmt.Println(len(comment.Comments))

		// 分析获取的评价(总评价,评价区间)
		flag := api.AnalyzeGetComments(&comment, &analyze)
		if !flag {
			c.JSON(http.StatusOK, cmd.Response{
				StatusCode: 1,
				StatusMsg:  "没有评论",
			})
			return
		}

		// 获取词云分析
		api.WordCloudAnalysis(&comment, &analyze)
	} else {
		// 从数据库中获取数据
		api.GetWordCloudAndAnalyzeRating(&analyze, id, lastTime)
	}

	c.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      analyze,
	})
}
