package main

import (
	"fmt"
	"strconv"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/fauzancodes/yugioh-open-api/repository"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cards, err := utils.LoadCardJSON()
	if err != nil {
		return
	}

	images, err := utils.LoadCloudinaryUrlJSON()
	if err != nil {
		return
	}

	config.Database()

	counter := 1
	for _, card := range cards {
		check, _ := repository.GetCardByID(uint(card.ID))
		if check.ID != 0 {
			existingCardSets, _ := repository.GetCardSetsByCardID(check.ID)
			// for _, existingCardSet := range existingCardSets {
			// 	repository.DeleteCardSet(existingCardSet)
			// }

			for _, set := range card.CardSets {
				var setFound bool
				for _, existingCardSet := range existingCardSets {
					if existingCardSet.SetCode == set.SetCode && existingCardSet.SetName == set.SetName {
						setFound = true
						break
					}
				}

				if !setFound {
					setData := models.YOACardSet{
						SetName:       set.SetName,
						SetCode:       set.SetCode,
						SetRarity:     set.SetRarity,
						SetRarityCode: set.SetRarityCode,
						CardID:        check.ID,
					}

					_, err = repository.CreateCardSet(setData)
					if err != nil {
						fmt.Println(fmt.Sprintf("(%v) Failed to update card:", counter), check.ID)
						counter++
						return
					}
				}
			}

			fmt.Println(fmt.Sprintf("(%v) Success to update card:", counter), check.ID)
			counter++
			continue
		}

		InsertToDatabase(card, images, counter)
		counter++
		fmt.Println(fmt.Sprintf("(%v) Success to insert card:", counter), card.ID)
	}
}

func InsertToDatabase(card dto.Card, images []dto.CardImage, counter int) {
	check, _ := repository.GetCardByID(uint(card.ID))
	if check.ID == 0 {
		var selectedImageurl string
		for _, image := range images {
			id, _ := strconv.Atoi(image.ID)
			if id == card.ID {
				selectedImageurl = image.Url
				break
			}
		}

		// var sets string
		// var rarity string
		// var rarityCode string
		// for _, set := range card.CardSets {
		// 	if sets != "" {
		// 		sets += ", "
		// 	}
		// 	sets += set.SetName

		// 	if set.SetRarityCode == "(C)" {
		// 		if rarityCode == "" {
		// 			rarity = "Common"
		// 			rarityCode = "C"
		// 		}
		// 	}
		// 	if set.SetRarityCode == "(R)" {
		// 		if rarityCode == "" || rarityCode == "C" {
		// 			rarity = "Rare"
		// 			rarityCode = "R"
		// 		}
		// 	}
		// 	if set.SetRarityCode == "(SR)" {
		// 		if rarityCode == "" || rarityCode == "C" || rarityCode == "R" {
		// 			rarity = "Super Rare"
		// 			rarityCode = "SR"
		// 		}
		// 	}
		// 	if set.SetRarityCode == "(UR)" {
		// 		if rarityCode == "" || rarityCode == "C" || rarityCode == "R" || rarityCode == "SR" {
		// 			rarity = "Ultra Rare"
		// 			rarityCode = "UR"
		// 		}
		// 	}
		// }
		// sets, _ := json.Marshal(card.CardSets)

		data := models.YOACard{
			CustomGormModel: models.CustomGormModel{
				ID: uint(card.ID),
			},
			Name:        card.Name,
			Type:        card.Type,
			Description: card.Description,
			Race:        card.Race,
			Archetype:   card.Archetype,
			Attack:      card.Attack,
			Defense:     card.Defense,
			Level:       card.Level,
			Attribute:   card.Attribute,
			ImageUrl:    selectedImageurl,
			// Rarity:      rarity,
			// RarityCode:  rarityCode,
			// CardSets: string(sets),
		}

		data, err := repository.CreateCard(data)
		if err != nil {
			fmt.Println("Failed to create card:", err.Error())
			return
		}

		for _, set := range card.CardSets {
			setData := models.YOACardSet{
				SetName:       set.SetName,
				SetCode:       set.SetCode,
				SetRarity:     set.SetRarity,
				SetRarityCode: set.SetRarityCode,
				CardID:        data.ID,
			}

			_, err = repository.CreateCardSet(setData)
			if err != nil {
				fmt.Println("Failed to create cardset:", err.Error())
				return
			}
		}

		fmt.Println(fmt.Sprintf("(%v) Success to create card:", counter), data.ID)
		return
	}

	fmt.Println(fmt.Sprintf("(%v) Card already exist:", counter), check.ID)
}
