package config

import (
	"fmt"
	"log"

	"github.com/fauzancodes/yugioh-open-api/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB {
	host := LoadConfig().DatabaseHost
	user := LoadConfig().DatabaseUsername
	password := LoadConfig().DatabasePassword
	name := LoadConfig().DatabaseName
	port := LoadConfig().DatabasePort
	schema := LoadConfig().DatabaseSchema

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s", host, user, password, name, port, schema)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if LoadConfig().EnableDatabaseAutomigration {
		go RunAutoMigration()
	}

	log.Printf("Connected to Database: %v", name)

	return DB
}

func RunAutoMigration() {
	err := DB.AutoMigrate(
		// &models.YOAUsedApiKey{},
		&models.YOACard{},
		&models.YOADeck{},
		&models.YOAUser{},
		&models.YOAMainDeck{},
		&models.YOAExtraDeck{},
		&models.YOASideDeck{},
		&models.YOACardSet{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
