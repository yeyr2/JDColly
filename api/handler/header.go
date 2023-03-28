package handler

import (
	"github.com/gin-gonic/gin"
)

func setHeader(con *gin.Context) {
	origin := con.GetHeader("Origin")
	if origin != "" {
		con.Header("Access-Control-Allow-Origin", origin)
	}

	con.Header("Access-Control-Allow-Methods", "*")

	con.Header("Access-Control-Max-Age", "3600")

	con.Header("Access-Control-Allow-Credentials", "true")

	header := con.Request.Header.Get("Access-Control-Allow-Headers")
	if header == "" {
		con.Header("Access-Control-Allow-Headers", "header")
	}
}
