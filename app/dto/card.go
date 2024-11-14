package dto

import (
	"errors"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CardRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Race        string `json:"race"`
	Archetype   string `json:"archetype"`
	Attack      int    `json:"attack"`
	Defense     int    `json:"defense"`
	Level       int    `json:"level"`
	Attribute   string `json:"attribute"`
	CardSets    string `json:"card_sets"`
	ImageUrl    string `json:"image_url"`
	Rarity      string `json:"rarity"`
	RarityCode  string `json:"rarity_code"`
}

func (request CardRequest) Validate() error {
	var pass bool

	var cardType []string
	err := config.DB.Raw("SELECT DISTINCT type FROM " + models.YOACard{}.TableName() + " WHERE type != ''").Scan(&cardType).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range cardType {
		if request.Type == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted types: " + strings.Join(cardType, ", "))
	}

	var rarity []string
	err = config.DB.Raw("SELECT DISTINCT rarity FROM " + models.YOACard{}.TableName() + " WHERE rarity != ''").Scan(&rarity).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range rarity {
		if request.Rarity == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted rarities: " + strings.Join(rarity, ", "))
	}

	var rarityCode []string
	err = config.DB.Raw("SELECT DISTINCT rarity_code FROM " + models.YOACard{}.TableName() + " WHERE rarity_code != ''").Scan(&rarityCode).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range rarityCode {
		if request.RarityCode == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted rarity codes: " + strings.Join(rarityCode, ", "))
	}

	var race []string
	err = config.DB.Raw("SELECT DISTINCT race FROM " + models.YOACard{}.TableName() + " WHERE race != ''").Scan(&race).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range race {
		if request.Race == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted races: " + strings.Join(race, ", "))
	}

	var archetype []string
	err = config.DB.Raw("SELECT DISTINCT archetype FROM " + models.YOACard{}.TableName() + " WHERE archetype != ''").Scan(&archetype).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range archetype {
		if request.Archetype == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted archetypes: " + strings.Join(archetype, ", "))
	}

	var attribute []string
	err = config.DB.Raw("SELECT DISTINCT attribute FROM " + models.YOACard{}.TableName() + " WHERE attribute != ''").Scan(&attribute).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range attribute {
		if request.Attribute == item {
			pass = true
			break
		}
	}
	if !pass {
		return errors.New("accepted attributes: " + strings.Join(attribute, ", "))
	}

	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Level, validation.Min(1), validation.Max(13)),
		validation.Field(&request.ImageUrl, is.URL),
	)
}
