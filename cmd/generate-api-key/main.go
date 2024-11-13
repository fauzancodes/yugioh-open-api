package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	secretKey := config.LoadConfig().SecretApiKey
	publicKey := config.LoadConfig().PublicApiKey
	fmt.Println("secretKey:", secretKey)
	fmt.Println("publicKey:", publicKey)

	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := h.Write([]byte(publicKey))
	if err != nil {
		fmt.Println("Failed to write signature:", err.Error())
	}
	signature := hex.EncodeToString(h.Sum(nil))
	fmt.Println("signature:", signature)

	pattern := fmt.Sprintf("%s:%s", publicKey, signature)
	fmt.Println("pattern:", pattern)

	encoded := base64.StdEncoding.EncodeToString([]byte(pattern))

	fmt.Println("x-api-key:", encoded)
}
