package controllers

import (
	"github.com/HugoJBello/go-heartbeat-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetActivity(c *gin.Context) {
	query := c.Request.URL.Query()

	limitStr := query["limit"][0]
	skipStr := query["skip"][0]

	if limitStr == "" {
		limitStr = "10000"
	}

	if skipStr == "" {
		skipStr = "0"
	}
	limit, _ := strconv.ParseInt(limitStr, 10, 64)
	skip, _ := strconv.ParseInt(skipStr, 10, 64)

	var lastActivity []models.ServiceActivity
	models.DB.Order("date desc").Offset(skip).Limit(limit).Find(&lastActivity)

	c.JSON(http.StatusOK, gin.H{"data": lastActivity})
}



func CreateActivity(c *gin.Context) {
	var input models.CreateServiceActivity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := models.ServiceActivity{Date: input.Date, ServiceId: input.ServiceId, ServiceName: input.ServiceName, ActivityContent: input.ActivityContent, ActivityContentType: input.ActivityContentType}
	models.DB.Create(&activity)

	c.JSON(http.StatusOK, gin.H{"data": []models.ServiceActivity{activity}})
}