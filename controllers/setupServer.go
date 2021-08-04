package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET(GetActivityRoute, GetActivity)
	r.POST(CreateActivityRoute, CreateActivity)

	return r
}
