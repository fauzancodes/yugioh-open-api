package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/xuri/excelize/v2"
)

func main() {
	cards, err := utils.LoadCardJSON()
	if err != nil {
		return
	}

	// Create a new Excel file
	f := excelize.NewFile()

	// Define headers
	headers := []string{
		"id", "name", "type", "description", "race", "archetype",
		"attack", "defense", "level", "attribute", "card_sets", "image_url",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Sheet1", cell, header)
	}

	// Populate the Excel file with data
	for i, card := range cards {
		row := i + 2 // Starting from row 2 because row 1 has headers

		// Card data
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), card.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), card.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), card.Type)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), card.Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), card.Race)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), card.Archetype)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), card.Attack)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), card.Defense)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), card.Level)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), card.Attribute)

		// Format card_sets
		var cardSets []string
		for _, set := range card.CardSets {
			cardSets = append(cardSets, fmt.Sprintf("%s %s %s", set.SetName, set.SetRarity, set.SetRarityCode))
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), strings.Join(cardSets, ", "))

		// Get the first image URL from card_images
		if len(card.CardImages) > 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), card.CardImages[0].ImageURL)
		}
	}

	// Save the Excel file
	if err := f.SaveAs("data/output.xlsx"); err != nil {
		log.Fatalf("Failed to save Excel file: %v", err)
	}

	fmt.Println("Data successfully written to output.xlsx")
}
