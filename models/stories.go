package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

// Story represents stories table
type Story struct {
	gorm.Model
	AppID    int `json:"app_id"`
	StoryID  int
	Metadata json.RawMessage `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// StoryDTO represents StoriesHandler response object
type StoryDTO struct {
	AppID     int        `json:"app_id"`
	Ts        int64      `json:"ts"` // current timestamp in second
	Metadatas []Metadata `json:"metadata"`
}

// Metadata represents StoryDTO Metadata field
type Metadata struct {
	ID       int             `json:"id"`
	Metadata json.RawMessage `json:"metadata"`
}

// StoriesToDTO converts []Story --> StoryDTO
func StoriesToDTO(stories []Story) StoryDTO {
	var metadatas []Metadata
	for _, s := range stories {
		m := Metadata{
			ID:       s.StoryID,
			Metadata: s.Metadata,
		}
		metadatas = append(metadatas, m)
	}

	return StoryDTO{
		AppID:     stories[0].AppID,
		Ts:        time.Now().Unix(),
		Metadatas: metadatas,
	}
}
