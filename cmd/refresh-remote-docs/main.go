package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Backblaze/blazer/b2"
	"github.com/fauzancodes/yugioh-open-api/app/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	keyID := config.LoadConfig().BackblazeKeyID
	applicationKey := config.LoadConfig().BackblazeApplicationKey
	bucketName := config.LoadConfig().BackblazeBucketName
	folder := config.LoadConfig().BackblazeFolder
	ctx := context.Background()

	b2, err := b2.NewClient(ctx, keyID, applicationKey)
	if err != nil {
		fmt.Println("Failed to connect to Backblaze:", err.Error())
		return
	}

	bucket, err := b2.Bucket(ctx, bucketName)
	if err != nil {
		fmt.Println("Backblaze bucket not found:", err.Error())
		return
	}

	srcs := []string{
		"assets/html/index.html",
		"docs/Yu-Gi-Oh! Open API.postman_collection.json",
		"docs/Yu-Gi-Oh! Open API.postman_environment.json",
	}
	dsts := []string{
		folder + "/assets/html/index.html",
		folder + "/docs/Yu-Gi-Oh! Open API.postman_collection.json",
		folder + "/docs/Yu-Gi-Oh! Open API.postman_environment.json",
	}

	for i, src := range srcs {
		f, err := os.Open(src)
		if err != nil {
			fmt.Println("Failed to open local file:", err.Error())
			return
		}
		defer f.Close()

		obj := bucket.Object(dsts[i])

		if _, err := obj.Attrs(ctx); err == nil {
			if err := obj.Delete(ctx); err != nil {
				fmt.Println("Failed to delete existing file in Backblaze:", err.Error())
				return
			}
			fmt.Println("Existing file deleted successfully:", dsts[i])
		}

		w := obj.NewWriter(ctx)
		if _, err := io.Copy(w, f); err != nil {
			w.Close()
			fmt.Println("Failed to write to backblaze:", err.Error())
			return
		}

		fmt.Println("Success to write to backblaze:", dsts[i])
		w.Close()
	}
}
