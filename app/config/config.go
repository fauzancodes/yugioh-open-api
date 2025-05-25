package config

import (
	"os"
	"strconv"
)

type Config struct {
	SecretKey                   string
	Port                        string
	BaseUrl                     string
	CloudinaryFolder            string
	CloudinaryCloudName         string
	CloudinaryAPIKey            string
	CLoudinaryAPISecret         string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	DatabaseSchema              string
	EnableDatabaseAutomigration bool
	EnableAPIKey                bool
	SpecialApiKey               string
	SecretApiKey                string
	PublicApiKey                string
	BackblazeFolder             string
	BackblazeBucketName         string
	BackblazeKeyID              string
	BackblazeApplicationKey     string
}

func LoadConfig() (config *Config) {
	secretKey := os.Getenv("SECRET_KEY")
	port := os.Getenv("PORT")
	baseUrl := os.Getenv("BASE_URL")
	cloudinaryFolder := os.Getenv("CLOUDINARY_FOLDER")
	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cLoudinaryAPISecret := os.Getenv("CLOUDINARY_API_SECRET")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseSchema := os.Getenv("DATABASE_SCHEMA")
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	enableApiKey, _ := strconv.ParseBool(os.Getenv("ENABLE_API_KEY"))
	specialApiKey := os.Getenv("SPECIAL_API_KEY")
	secretApiKey := os.Getenv("SECRET_API_KEY")
	publicApiKey := os.Getenv("PUBLIC_API_KEY")
	backblazeFolder := os.Getenv("BACKBLAZE_FOLDER")
	backblazeBucketName := os.Getenv("BACKBLAZE_BUCKET_NAME")
	backblazeKeyID := os.Getenv("BACKBLAZE_KEY_ID")
	backblazeApplicationKey := os.Getenv("BACKBLAZE_APPLICATION_KEY")

	return &Config{
		SecretKey:                   secretKey,
		Port:                        port,
		BaseUrl:                     baseUrl,
		CloudinaryFolder:            cloudinaryFolder,
		CloudinaryCloudName:         cloudinaryCloudName,
		CloudinaryAPIKey:            cloudinaryAPIKey,
		CLoudinaryAPISecret:         cLoudinaryAPISecret,
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		DatabaseSchema:              databaseSchema,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		EnableAPIKey:                enableApiKey,
		SpecialApiKey:               specialApiKey,
		SecretApiKey:                secretApiKey,
		PublicApiKey:                publicApiKey,
		BackblazeFolder:             backblazeFolder,
		BackblazeBucketName:         backblazeBucketName,
		BackblazeKeyID:              backblazeKeyID,
		BackblazeApplicationKey:     backblazeApplicationKey,
	}
}
