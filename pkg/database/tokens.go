package database

import (
	"github.com/jinzhu/gorm"
	"story/models"
)

// In the first implementation this function used, but now not used for better performance
func CheckAppTokenExists(db *gorm.DB, appToken string) (bool, error) {
	var t models.Token
	query := db.Where("token = ?", appToken)
	if err := query.First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetAppIDByAppToken return app_id by app_token
func GetAppIDByAppToken(db *gorm.DB, appToken string) (int, error) {
	var t models.Token
	query := db.Where("token = ?", appToken)
	if err := query.First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return t.AppID, nil
}
