package dto

import (
	"errors"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"slices"
)

type CardRequest struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Race        string    `json:"race"`
	Archetype   string    `json:"archetype"`
	Attack      int       `json:"attack"`
	Defense     int       `json:"defense"`
	Level       int       `json:"level"`
	Attribute   string    `json:"attribute"`
	CardSets    []CardSet `json:"card_sets"`
	ImageUrl    string    `json:"image_url"`
	// Rarity      string `json:"rarity"`
	// RarityCode  string `json:"rarity_code"`
}

func (request CardRequest) Validate() error {
	var pass bool

	var cardType []string
	err := config.DB.Raw("SELECT DISTINCT type FROM " + models.YOACard{}.TableName() + " WHERE type != ''").Scan(&cardType).Error
	if err != nil {
		return err
	}
	pass = slices.Contains(cardType, request.Type)
	if !pass {
		return errors.New("accepted types: " + strings.Join(cardType, ", "))
	}

	var rarity []string
	err = config.DB.Raw("SELECT DISTINCT set_rarity FROM " + models.YOACardSet{}.TableName() + " WHERE set_rarity != ''").Scan(&rarity).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range rarity {
		for _, set := range request.CardSets {
			if set.SetRarity == item {
				pass = true
				break
			}
		}
	}
	if !pass {
		return errors.New("accepted rarities: " + strings.Join(rarity, ", "))
	}

	var rarityCode []string
	err = config.DB.Raw("SELECT DISTINCT set_rarity_code FROM " + models.YOACardSet{}.TableName() + " WHERE set_rarity_code != ''").Scan(&rarityCode).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range rarityCode {
		for _, set := range request.CardSets {
			if set.SetRarityCode == item {
				pass = true
				break
			}
		}
	}
	if !pass {
		return errors.New("accepted rarity codes: " + strings.Join(rarityCode, ", "))
	}

	var cardsets []string
	err = config.DB.Raw("SELECT DISTINCT set_name FROM " + models.YOACardSet{}.TableName() + " WHERE set_name != ''").Scan(&cardsets).Error
	if err != nil {
		return err
	}
	pass = false
	for _, item := range rarityCode {
		for _, set := range request.CardSets {
			if set.SetName == item {
				pass = true
				break
			}
		}
	}
	if !pass {
		return errors.New("accepted set names: " + strings.Join(rarityCode, ", "))
	}

	var race []string
	err = config.DB.Raw("SELECT DISTINCT race FROM " + models.YOACard{}.TableName() + " WHERE race != ''").Scan(&race).Error
	if err != nil {
		return err
	}
	pass = slices.Contains(race, request.Race)
	if !pass {
		return errors.New("accepted races: " + strings.Join(race, ", "))
	}

	var archetype []string
	err = config.DB.Raw("SELECT DISTINCT archetype FROM " + models.YOACard{}.TableName() + " WHERE archetype != ''").Scan(&archetype).Error
	if err != nil {
		return err
	}
	pass = slices.Contains(archetype, request.Archetype)
	if !pass {
		return errors.New("accepted archetypes: " + strings.Join(archetype, ", "))
	}

	var attribute []string
	err = config.DB.Raw("SELECT DISTINCT attribute FROM " + models.YOACard{}.TableName() + " WHERE attribute != ''").Scan(&attribute).Error
	if err != nil {
		return err
	}
	pass = slices.Contains(attribute, request.Attribute)
	if !pass {
		return errors.New("accepted attributes: " + strings.Join(attribute, ", "))
	}

	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Level, validation.Min(1), validation.Max(13)),
		validation.Field(&request.ImageUrl, is.URL),
	)
}
