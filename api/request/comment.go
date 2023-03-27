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
	var jdComment cmd.JDComment
	comments := new([]cmd.Comments)

	if isColly != 1 {
		// 获取评论
		api.GetCommentById(id, lastTime, &jdComment)
		//fmt.Println(len(jdComment.Comments))
		comments = &jdComment.Comments
	} else {
		// 从数据库中获取数据
		comments = api.GetCommentBySql(id, lastTime)
	}

	// 分析获取的评价(总评价,评价区间)
	flag := api.AnalyzeGetComments(comments, &analyze)
	if !flag {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "没有评论",
		})
		return
	}

	// 获取词云分析
	api.WordCloudAnalysis(comments, &analyze)

	setHeader(c)

	c.JSON(http.StatusOK, cmd.Response{
		StatusCode: 0,
		StatusMsg:  "",
		Value:      analyze,
	})
}
