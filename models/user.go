package models

// User model
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique_index" json:"username"`
	Email    string `gorm:"unique_index" json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"default:participant" json:"role"`
}
