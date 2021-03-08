package controllers

import "github.com/gin-gonic/gin"

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/activity/last", GetActivity)
	r.POST("/activity/new", CreateActivity)

	return r
}
