package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
)

func LoadJSON() (cards []dto.Card, err error) {
	jsonFile, err := os.Open("data/cardinfo.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)

		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var data dto.Data
	json.Unmarshal(byteValue, &data)

	cards = data.Data

	return
}