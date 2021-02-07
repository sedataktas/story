package models

import "github.com/jinzhu/gorm"

// Event represents events table
type Event struct {
	gorm.Model
	AppID   int
	StoryID int    `json:"story_id"`
	Type    string `json:"event_type"`
	Count   int
}
