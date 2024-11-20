package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/pkg/upload"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/fauzancodes/yugioh-open-api/repository"
	"gorm.io/gorm"
)

func CreateCard(request dto.CardRequest) (response models.YOACard, statusCode int, err error) {
	data := models.YOACard{
		CustomGormModel: models.CustomGormModel{ID: request.ID},
		Name:            request.Name,
		Type:            request.Type,
		Description:     request.Description,
		Race:            request.Race,
		Archetype:       request.Archetype,
		Attack:          request.Attack,
		Defense:         request.Defense,
		Level:           request.Level,
		Attribute:       request.Attribute,
		CardSets:        request.CardSets,
		ImageUrl:        request.ImageUrl,
		Rarity:          request.Rarity,
		RarityCode:      request.RarityCode,
	}

	response, err = repository.CreateCard(data)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusCreated
	return
}

func GetCardByID(id uint) (data models.YOACard, statusCode int, err error) {
	data, err = repository.GetCardByID(id)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func GetCards(cardType, race, archetype, attribute, cardsets, rarity, rarityCode string, attack, attackMarginTop, attackMarginBottom, defense, defenseMarginTop, defenseMarginBottom, level, levelMarginTop, levelMarginBottom int, param utils.PagingRequest) (response utils.PagingResponse, data []models.YOACard, statusCode int, err error) {
	baseFilter := "deleted_at IS NULL"
	filter := baseFilter
	var filterValues []any

	if cardType != "" {
		filter += " AND type = ?"
		filterValues = append(filterValues, cardType)
	}
	if race != "" {
		filter += " AND race = ?"
		filterValues = append(filterValues, race)
	}
	if archetype != "" {
		filter += " AND archetype = ?"
		filterValues = append(filterValues, archetype)
	}
	if attribute != "" {
		filter += " AND attribute = ?"
		filterValues = append(filterValues, attribute)
	}
	if cardsets != "" {
		filter += " AND card_sets ILIKE ?"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", cardsets))
	}
	if rarity != "" {
		filter += " AND rarity = ?"
		filterValues = append(filterValues, rarity)
	}
	if rarityCode != "" {
		filter += " AND rarity_code = ?"
		filterValues = append(filterValues, rarityCode)
	}
	if attack > 0 {
		filter += " AND attack = ?"
		filterValues = append(filterValues, attack)
	}
	if attackMarginTop > 0 {
		filter += " AND attack <= ?"
		filterValues = append(filterValues, attackMarginTop)
	}
	if attackMarginBottom > 0 {
		filter += " AND attack >= ?"
		filterValues = append(filterValues, attackMarginBottom)
	}
	if defense > 0 {
		filter += " AND defense = ?"
		filterValues = append(filterValues, defense)
	}
	if defenseMarginTop > 0 {
		filter += " AND defense <= ?"
		filterValues = append(filterValues, defenseMarginTop)
	}
	if defenseMarginBottom > 0 {
		filter += " AND defense >= ?"
		filterValues = append(filterValues, defenseMarginBottom)
	}
	if level > 0 {
		filter += " AND level = ?"
		filterValues = append(filterValues, level)
	}
	if levelMarginTop > 0 {
		filter += " AND level <= ?"
		filterValues = append(filterValues, levelMarginTop)
	}
	if levelMarginBottom > 0 {
		filter += " AND level >= ?"
		filterValues = append(filterValues, levelMarginBottom)
	}

	if param.Search != "" {
		filter += " AND (name ILIKE ? OR description ILIKE ?)"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
	}

	data, total, totalFiltered, err := repository.GetCards(dto.FindParameter{
		BaseFilter:   baseFilter,
		Filter:       filter,
		FilterValues: filterValues,
		Limit:        param.Limit,
		Order:        param.Order,
		Offset:       param.Offset,
	})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	response = utils.PopulateResPaging(&param, data, total, totalFiltered)

	statusCode = http.StatusOK
	return
}

func UpdateCard(id uint, request dto.CardRequest) (response models.YOACard, statusCode int, err error) {
	data, err := repository.GetCardByID(id)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if request.ID > 0 {
		data.ID = request.ID
	}
	if request.Name != "" {
		data.Name = request.Name
	}
	if request.Type != "" {
		data.Type = request.Type
	}
	if request.Description != "" {
		data.Description = request.Description
	}
	if request.Race != "" {
		data.Race = request.Race
	}
	if request.Archetype != "" {
		data.Archetype = request.Archetype
	}
	if request.Attack > 0 {
		data.Attack = request.Attack
	}
	if request.Defense > 0 {
		data.Defense = request.Defense
	}
	if request.Level > 0 {
		data.Level = request.Level
	}
	if request.Attribute != "" {
		data.Attribute = request.Attribute
	}
	if request.CardSets != "" {
		data.CardSets = request.CardSets
	}
	if request.ImageUrl != "" {
		data.ImageUrl = request.ImageUrl
	}
	if request.Rarity != "" {
		data.Rarity = request.Rarity
	}
	if request.RarityCode != "" {
		data.RarityCode = request.RarityCode
	}

	response, err = repository.UpdateCard(data)
	if err != nil {
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func DeleteCard(id uint) (statusCode int, err error) {
	data, err := repository.GetCardByID(id)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	err = repository.DeleteCard(data)
	if err != nil {
		err = errors.New("failed to delete data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func GetCardUtility(field string) (responses []string, statusCode int, err error) {
	acceptedField := []string{"type", "race", "archetype", "level", "card_sets", "rarity", "attribute", "rarity_code"}
	var accepted bool
	for _, item := range acceptedField {
		if item == field {
			accepted = true
			break
		}
	}
	if !accepted {
		statusCode = http.StatusBadRequest
		err = errors.New("accepted fields: " + strings.Join(acceptedField, ", "))

		return
	}

	responses, err = repository.GetCardUtility(field)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if field == "card_sets" {
		responsesCopy := responses
		responses = []string{}
		for _, data := range responsesCopy {
			items := strings.Split(data, ", ")
			responses = append(responses, items...)
		}
		responses = utils.RemoveDuplicatesFromStringArray(responses)
	}

	return
}

func UploadCardPicture(file *multipart.FileHeader) (responseURL string, statusCode int, err error) {
	extension := filepath.Ext(file.Filename)
	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" && extension != ".webp" {
		err = errors.New("the file extension is wrong. allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
		statusCode = http.StatusBadRequest
		return
	}

	var src multipart.File
	src, err = file.Open()
	if err != nil {
		err = errors.New("failed to open file: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}
	defer src.Close()

	responseURL, _, _, err = upload.UploadImageOrVideo(src, "", "")
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}
