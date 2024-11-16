package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/fauzancodes/yugioh-open-api/repository"
	"gorm.io/gorm"
)

func CreateDeck(userID uint, request dto.DeckRequest) (response models.YOADeck, statusCode int, err error) {
	mainDeckCards, extraDeckCards, sideDeckCards, totalSpellCard, totalTrapCard, totalMonsterCard, totalDeckCard, totalMainDeckCard, totalExtraDeckCard, totalSideDeckCard, err := AdjustDeckCards(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		return
	}

	data := models.YOADeck{
		Name:               request.Name,
		Description:        request.Description,
		UserID:             userID,
		TotalDeckCard:      totalDeckCard,
		TotalMainDeckCard:  totalMainDeckCard,
		TotalExtraDeckCard: totalExtraDeckCard,
		TotalSideDeckCard:  totalSideDeckCard,
		TotalMonsterCard:   totalMonsterCard,
		TotalSpellCard:     totalSpellCard,
		TotalTrapCard:      totalTrapCard,
		IsPublic:           request.IsPublic,
	}

	response, err = repository.CreateDeck(data, mainDeckCards, extraDeckCards, sideDeckCards)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusCreated
	return
}

func AdjustDeckCards(request dto.DeckRequest) (mainDeckCards []models.YOACard, extraDeckCards []models.YOACard, sideDeckCards []models.YOACard, totalSpellCard uint, totalTrapCard uint, totalMonsterCard uint, totalDeckCard uint, totalMainDeckCard uint, totalExtraDeckCard uint, totalSideDeckCard uint, err error) {
	var allCards []uint
	allCards = append(allCards, request.MainDeckCardID...)
	allCards = append(allCards, request.ExtraDeckCardID...)
	allCards = append(allCards, request.SideDeckCardID...)
	moreThanThreeCards := utils.GetDuplicatesMoreThanThree(allCards)
	if len(moreThanThreeCards) > 0 {
		var str []string
		for _, item := range moreThanThreeCards {
			str = append(str, strconv.FormatUint(uint64(item), 10))
		}
		err = errors.New("only 3 copies of a card are allowed. cards that have more than 3 copies: " + strings.Join(str, ", "))

		return
	}

	if len(request.MainDeckCardID) > 0 {
		for _, item := range request.MainDeckCardID {
			card, _ := repository.GetCardByID(item)
			if card.ID == 0 {
				err = errors.New("card not found. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "token") {
				err = errors.New("token cards should not be in deck. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "fusion") {
				err = errors.New("fusion cards should be in the extra deck. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "synchro") {
				err = errors.New("synchro cards should be in the extra deck. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "xyz") {
				err = errors.New("xyz cards should be in the extra deck. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "link") {
				err = errors.New("link cards should be in the extra deck. card_id: " + strconv.Itoa(int(item)))
				return
			}

			if strings.ToLower(card.Type) == "spell card" {
				totalSpellCard++
			} else if strings.ToLower(card.Type) == "trap card" {
				totalTrapCard++
			} else {
				totalMonsterCard++
			}

			totalDeckCard++
			totalMainDeckCard++
			mainDeckCards = append(mainDeckCards, card)
		}
	}

	if len(request.ExtraDeckCardID) > 0 {
		for _, item := range request.ExtraDeckCardID {
			card, _ := repository.GetCardByID(item)
			if card.ID == 0 {
				err = errors.New("card not found. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if !strings.Contains(strings.ToLower(card.Type), "fusion") && !strings.Contains(strings.ToLower(card.Type), "synchro") && !strings.Contains(strings.ToLower(card.Type), "xyz") && !strings.Contains(strings.ToLower(card.Type), "link") {
				err = errors.New("besides fusion, synchro, xyz or link cards should be in the main deck or side deck. card_id: " + strconv.Itoa(int(item)))
				return
			}

			totalMonsterCard++
			totalDeckCard++
			totalExtraDeckCard++
			extraDeckCards = append(extraDeckCards, card)
		}
	}

	if len(request.SideDeckCardID) > 0 {
		for _, item := range request.SideDeckCardID {
			card, _ := repository.GetCardByID(item)
			if card.ID == 0 {
				err = errors.New("card not found. card_id: " + strconv.Itoa(int(item)))
				return
			}
			if strings.Contains(strings.ToLower(card.Type), "token") {
				err = errors.New("token cards should not be in deck. card_id: " + strconv.Itoa(int(item)))
				return
			}

			if strings.ToLower(card.Type) == "spell card" {
				totalSpellCard++
			} else if strings.ToLower(card.Type) == "trap card" {
				totalTrapCard++
			} else {
				totalMonsterCard++
			}

			totalDeckCard++
			totalSideDeckCard++
			sideDeckCards = append(sideDeckCards, card)
		}
	}

	return
}

func GetDeckByID(id uint, preloadFields []string) (data models.YOADeck, statusCode int, err error) {
	data, err = repository.GetDeckByID(id, preloadFields)
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

func GetDecks(userID uint, param utils.PagingRequest, preloadFields []string) (response utils.PagingResponse, data []models.YOADeck, statusCode int, err error) {
	baseFilter := "deleted_at IS NULL"
	var baseFilterValues []any
	if userID > 0 {
		baseFilter += " AND user_id = ?"
		baseFilterValues = append(baseFilterValues, userID)
	}
	filter := baseFilter
	filterValues := baseFilterValues

	if param.Custom != "" {
		filter += " AND is_public = ?"
		filterValues = append(filterValues, param.Custom)
	}
	if param.Search != "" {
		filter += " AND (name ILIKE ? OR description ILIKE ?)"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
	}

	data, total, totalFiltered, err := repository.GetDecks(dto.FindParameter{
		BaseFilter:       baseFilter,
		BaseFilterValues: baseFilterValues,
		Filter:           filter,
		FilterValues:     filterValues,
		Limit:            param.Limit,
		Order:            param.Order,
		Offset:           param.Offset,
	}, preloadFields)
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

func UpdateDeck(id uint, request dto.DeckRequest) (response models.YOADeck, statusCode int, err error) {
	data, err := repository.GetDeckByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if request.Name != "" {
		data.Name = request.Name
	}
	if request.Description != "" {
		data.Description = request.Description
	}
	data.IsPublic = request.IsPublic

	if len(request.MainDeckCardID) == 0 {
		for _, item := range data.MainDeckCards {
			request.MainDeckCardID = append(request.MainDeckCardID, item.ID)
		}
	}
	if len(request.ExtraDeckCardID) == 0 {
		for _, item := range data.ExtraDeckCards {
			request.ExtraDeckCardID = append(request.ExtraDeckCardID, item.ID)
		}
	}
	if len(request.SideDeckCardID) == 0 {
		for _, item := range data.SideDeckCards {
			request.SideDeckCardID = append(request.SideDeckCardID, item.ID)
		}
	}

	mainDeckCards, extraDeckCards, sideDeckCards, totalSpellCard, totalTrapCard, totalMonsterCard, totalDeckCard, totalMainDeckCard, totalExtraDeckCard, totalSideDeckCard, err := AdjustDeckCards(request)
	if err != nil {
		statusCode = http.StatusBadRequest
		return
	}

	data.TotalDeckCard = totalDeckCard
	data.TotalExtraDeckCard = totalExtraDeckCard
	data.TotalMainDeckCard = totalMainDeckCard
	data.TotalMonsterCard = totalMonsterCard
	data.TotalSideDeckCard = totalSideDeckCard
	data.TotalSpellCard = totalSpellCard
	data.TotalTrapCard = totalTrapCard

	response, err = repository.UpdateDeck(data, mainDeckCards, extraDeckCards, sideDeckCards)
	if err != nil {
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func DeleteDeck(id uint) (statusCode int, err error) {
	data, err := repository.GetDeckByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	err = repository.DeleteDeck(data)
	if err != nil {
		err = errors.New("failed to delete data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}
