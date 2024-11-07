package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fauzancodes/yugioh-open-api/app/utils"
)

// DownloadImage downloads an image from the given URL and saves it to the specified filepath
func DownloadImage(url, filepath string) error {
	// Send an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Create a local file to save the downloaded image
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the response body data to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

// DownloadImagesWithRateLimit downloads images from a list of URLs with a rate limit
func DownloadImagesWithRateLimit(urls []string, rateLimit int) {
	ticker := time.NewTicker(time.Second / time.Duration(rateLimit)) // Create a ticker to allow only `rateLimit` downloads per second
	defer ticker.Stop()

	for i, url := range urls {
		<-ticker.C // Wait for the ticker before starting the next download

		go func(url string, i int) {
			urlSplited := strings.Split(url, "/")
			imageName := urlSplited[len(urlSplited)-1]
			filepath := fmt.Sprintf("images/%v", imageName) // Set file path with index to avoid overwriting

			fmt.Printf("Starting download for: %s\n", url)
			err := DownloadImage(url, filepath)
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", url, err)
			} else {
				fmt.Printf("Successfully downloaded: %s\n", filepath)
			}
		}(url, i)
	}

	// Allow time for all goroutines to finish
	time.Sleep(time.Duration(len(urls)/rateLimit+1) * time.Second)
}

func main() {
	cards, err := utils.LoadJSON()
	if err != nil {
		return
	}

	// List of URLs to download
	var urls []string
	for _, item := range cards {
		urls = append(urls, item.CardImages[0].ImageURL)
	}

	fmt.Println("Starting batch download with rate limit...")

	// Download images with a rate limit of 15 per second
	DownloadImagesWithRateLimit(urls, 15)
}
