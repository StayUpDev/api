package models

// User model
type User struct {
    ID       uint   `gorm:"primary_key" json:"id"`
    Email    string `gorm:"unique_index" json:"email"`
    Password string `json:"password"`
}

