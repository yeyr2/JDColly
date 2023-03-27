package sql

import (
	"log"
	"reptile-test-go/api/cmd"
)

type sqlComment struct {
	ProductId int64  `gorm:"column:product_id"`
	Context   string `gorm:"column:context"`
	EnContext string `gorm:"column:en_context"`
	OldScore  int    `gorm:"column:old_score"`
}

func (s sqlComment) TableName() string {
	return "comments"
}

func SaveComment(comments cmd.JDComment) {
	if len(comments.Comments) == 0 {
		return
	}
	sqlCom := make([]*sqlComment, 0, len(comments.Comments))

	for _, comment := range comments.Comments {
		if comment.UsefulVoteCount == 0 {
			continue
		}
		tmp := &sqlComment{
			ProductId: comments.ProductCommentSummary.ProductID,
			Context:   comment.Content,
			EnContext: comment.EnContent,
			OldScore:  comment.Score,
		}
		sqlCom = append(sqlCom, tmp)
	}
	res := db.Create(sqlCom)
	if res.RowsAffected == 0 {
		log.Println(res.Error)
	}
}
