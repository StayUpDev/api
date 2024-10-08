package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
    // AutoMigrate returns a *gorm.DB instance, and the Error field will hold any migration error
    return db.AutoMigrate(&User{}, &Participant{}, &Promoter{})
}