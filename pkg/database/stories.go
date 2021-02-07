package database

import (
	"github.com/jinzhu/gorm"
	"story/models"
)

// GetMetadatasByAppToken joins tokens and stories table
// it takes app_id by app_token
// then take stories by app_id
func GetMetadatasByAppToken(db *gorm.DB, appToken string) ([]models.Story, error) {
	var stories []models.Story
	query := db.Select("stories.*")
	query = query.Joins("LEFT JOIN tokens ON stories.app_id =tokens.app_id")
	query = query.Where("token = ?", appToken)
	if err := query.Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

// CheckStoryIDExists checks story_id exist in stories table according to this :
// app_id && story_id
func CheckStoryIDExists(db *gorm.DB, appID, storyID int) (bool, error) {
	var s models.Story
	query := db.Where("app_id = ? AND story_id = ?", appID, storyID)
	if err := query.First(&s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
