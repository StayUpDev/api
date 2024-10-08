package models

// User model
type Promoter struct {
	ID                     uint `gorm:"primary_key" json:"id"` // Primary key
	UserID                 uint
	User                   User
	CurrentEvents          []Event       `gorm:"foreignKey:PromoterID" json:"current_events"`
	PastEvents             []Event       `gorm:"foreignKey:PromoterID" json:"past_events"`
	ScheduledEvents        []Event       `gorm:"foreignKey:PromoterID" json:"scheduled_events"`
	Address                string        `json:"address"`
	Rating                 float32       `json:"rating"`
	SubscribedParticipants []Participant `gorm:"many2many:promoter_participants;" json:"subscribed_participants"` // Many-to-Many relationship
	Likes                  uint32        `json:"likes"`
	Tags                   []string      `gorm:"type:text[]" json:"tags"`   // Store tags as an array in PostgreSQL
	Images                 []string      `gorm:"type:text[]" json:"images"` // Store multiple image URLs
}
