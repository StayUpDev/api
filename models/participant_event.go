package models

type ParticipantEvent struct {
	ID            uint `gorm:"primary_key" json:"id"`
	IdParticipant uint
	IdEvent       uint
	IsLiked       bool
	Participant   Participant `gorm:"foreignKey:IdParticipant" json:"participant"`
	Event         Event       `gorm:"foreignKey:IdEvent" json:"event"`
}
