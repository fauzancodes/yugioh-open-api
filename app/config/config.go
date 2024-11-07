package config

import "os"

type Config struct {
	CloudinaryFolder    string
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CLoudinaryAPISecret string
}

func LoadConfig() (config *Config) {
	cloudinaryFolder := os.Getenv("CLOUDINARY_FOLDER")
	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cLoudinaryAPISecret := os.Getenv("CLOUDINARY_API_SECRET")

	return &Config{
		CloudinaryFolder:    cloudinaryFolder,
		CloudinaryCloudName: cloudinaryCloudName,
		CloudinaryAPIKey:    cloudinaryAPIKey,
		CLoudinaryAPISecret: cLoudinaryAPISecret,
	}
}
