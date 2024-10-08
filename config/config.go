package config

import (
	"jwt_go_server/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection with Gorm v2
func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=stayUp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	models.Migrate(db)

	return db
}
