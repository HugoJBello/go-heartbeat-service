package managers

import (
	"errors"
	"github.com/HugoJBello/go-heartbeat-service/models"
)

func GetActivity(skip int64, limit int64) *[]models.ServiceActivity {
	var lastActivity []models.ServiceActivity
	GetDb().Order("date desc").Offset(skip).Limit(limit).Find(&lastActivity)
	return &lastActivity
}

func CreateActivity(activity *models.ServiceActivity) error {
	result := GetDb().Create(&activity)
	if result != nil {
		return errors.New("no data")
	}
	return nil
}
