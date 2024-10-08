package models

import (
	"time"
)


type Event struct {
    ID                   uint          `gorm:"primary_key" json:"id"`
    Name                 string        `json:"name"`
    Description          string        `json:"description"`
    PromoterID           uint          `json:"promoter_id"` // Foreign key to Promoter
    EventStartDatetime   time.Time     `json:"event_start_datetime"`
    EventEndDatetime     time.Time     `json:"event_end_datetime"`
    Participants         []Participant `gorm:"many2many:event_participants;" json:"participants"` // Many-to-many relationship
    Price                float32       `json:"price"`
    Location             string        `json:"location"`
    Parking              *bool         `json:"parking"` // Nullable boolean
    DressCode            string        `json:"dress_code"`
    HotelsNearby         *bool         `json:"hotels_nearby"` // Nullable boolean
    Likes                uint64        `json:"likes"`
    Tags                 []string      `gorm:"type:text[]" json:"tags"`
}






