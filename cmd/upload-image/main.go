package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fauzancodes/yugioh-open-api/app/pkg/upload"
	_ "github.com/joho/godotenv/autoload"
)

type UploadResult struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func main() {
	// Define the path to the images folder and initialize a slice to hold the upload results
	folderPath := "images"
	var results []UploadResult
	counter := 1

	// Walk through all files in the images folder
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is an image (not a directory)
		if !info.IsDir() {
			fileName := info.Name()
			filename := fileName[:len(fileName)-len(filepath.Ext(fileName))] // Name without extension

			// Check if the image already exists in Cloudinary
			isExist, secureUrl := upload.CheckAssetExistanceByPublicID(filename)
			if isExist {
				results = append(results, UploadResult{
					ID:  filename,
					URL: secureUrl,
				})
				fmt.Printf("(%v) File %s already exists in Cloudinary, skipping upload\n", counter, fileName)
				counter++
				return nil
			}

			// Upload the image to Cloudinary if it does not exist
			secureUrl, _, _, err := upload.UploadFile(path, "", filename)
			if err != nil {
				fmt.Printf("Failed to upload %s: %v\n", fileName, err)
			} else {
				fmt.Printf("(%v) Successfully uploaded %s: %s\n", counter, fileName, secureUrl)
				counter++

				// Append the result to the results slice
				results = append(results, UploadResult{
					ID:  filename,
					URL: secureUrl,
				})
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to walk through the folder: %v", err)
	}

	// Save the results to a JSON file
	jsonFile, err := os.Create("data/cloudinary_urls.json")
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ") // Optional: format the JSON with indentation
	if err := encoder.Encode(results); err != nil {
		log.Fatalf("Failed to write JSON data: %v", err)
	}

	fmt.Println("Successfully saved upload results to cloudinary_urls.json")
}
