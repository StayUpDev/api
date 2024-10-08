package utils

import (
	"errors"
	"jwt_go_server/models"

	"gorm.io/gorm"
)

const (
	PromoterRole    = "promoter"
	participantRole = "participant"
)

func GetUserRole(db *gorm.DB, userID string) (string, error) {

	var promoter models.Promoter
	if err := db.Where("user_id = ?", userID).First(&promoter).Error; err != nil {
		return PromoterRole, nil
	}

	var participant models.Participant
	if err := db.Where("user_id = ?", userID).First(&participant).Error; err != nil {
		return participantRole, nil
	}

	return "", errors.New("user is neither participant nor promoter")

}
