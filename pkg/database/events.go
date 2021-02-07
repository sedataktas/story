package database

import (
	"github.com/jinzhu/gorm"
	"story/models"
)

// GetEvent return event row according to this:
// app_id && story_id && type
func GetEvent(db *gorm.DB, eventFromBody models.Event) (models.Event, error) {
	var event models.Event
	query := db.Where("app_id = ? AND story_id = ? AND type = ?",
		eventFromBody.AppID,
		eventFromBody.StoryID,
		eventFromBody.Type)
	if err := query.First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return event, nil
		}
		return event, err
	}
	return event, nil
}
