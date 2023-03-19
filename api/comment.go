package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetComment(c *gin.Context) {
	id := c.Query("id")
	//LastTime := c.Query("LastTime")

	var analyze AnalyzeComment
	var comment JDComment

	// 获取评论
	GetCommentById(id, &comment)
	//fmt.Println(len(comment.Comments))

	// 分析获取的评价(总评价,评价区间)
	AnalyzeGetComments(&comment, &analyze)

	// 获取词云分析
	WordCloudAnalysis(&comment, &analyze)

	c.JSON(http.StatusOK, analyze)
}
