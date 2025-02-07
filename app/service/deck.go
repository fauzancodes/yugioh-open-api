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

	response, err = repository.CreateDeck(data)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	mainDeckChan := make([]chan models.YOAMainDeck, len(mainDeckCards))
	mainDeckErrChan := make([]chan error, len(mainDeckCards))
	for i, item := range mainDeckCards {
		mainDeckChan[i] = make(chan models.YOAMainDeck)
		mainDeckErrChan[i] = make(chan error)
		go CreateMainDeckCard(models.YOAMainDeck{
			DeckID: response.ID,
			CardID: item.ID,
		}, mainDeckChan[i], mainDeckErrChan[i])
	}

	for i := range mainDeckCards {
		select {
		case err = <-mainDeckErrChan[i]:
			if err != nil {
				err = errors.New("failed to create main deck: " + err.Error())
				statusCode = http.StatusInternalServerError
				return
			}
		case mainDeck := <-mainDeckChan[i]:
			response.MainDeck = append(response.MainDeck, mainDeck)
		}
	}

	extraDeckChan := make([]chan models.YOAExtraDeck, len(extraDeckCards))
	extraDeckErrChan := make([]chan error, len(extraDeckCards))
	for i, item := range extraDeckCards {
		extraDeckChan[i] = make(chan models.YOAExtraDeck)
		extraDeckErrChan[i] = make(chan error)
		go CreateExtraDeckCard(models.YOAExtraDeck{
			DeckID: response.ID,
			CardID: item.ID,
		}, extraDeckChan[i], extraDeckErrChan[i])
	}

	for i := range extraDeckCards {
		select {
		case err = <-extraDeckErrChan[i]:
			if err != nil {
				err = errors.New("failed to create extra deck: " + err.Error())
				statusCode = http.StatusInternalServerError
				return
			}
		case extraDeck := <-extraDeckChan[i]:
			response.ExtraDeck = append(response.ExtraDeck, extraDeck)
		}
	}

	sideDeckChan := make([]chan models.YOASideDeck, len(sideDeckCards))
	sideDeckErrChan := make([]chan error, len(sideDeckCards))
	for i, item := range sideDeckCards {
		sideDeckChan[i] = make(chan models.YOASideDeck)
		sideDeckErrChan[i] = make(chan error)
		go CreateSideDeckCard(models.YOASideDeck{
			DeckID: response.ID,
			CardID: item.ID,
		}, sideDeckChan[i], sideDeckErrChan[i])
	}

	for i := range sideDeckCards {
		select {
		case err = <-sideDeckErrChan[i]:
			if err != nil {
				err = errors.New("failed to create side deck: " + err.Error())
				statusCode = http.StatusInternalServerError
				return
			}
		case sideDeck := <-sideDeckChan[i]:
			response.SideDeck = append(response.SideDeck, sideDeck)
		}
	}

	statusCode = http.StatusCreated
	return
}

func CreateMainDeckCard(data models.YOAMainDeck, result chan models.YOAMainDeck, errChan chan error) {
	var err error
	data, err = repository.CreateMainDeck(data)
	if err != nil {
		err = errors.New("failed to create main deck card: " + err.Error())
		errChan <- err
		return
	}

	result <- data
}

func CreateExtraDeckCard(data models.YOAExtraDeck, result chan models.YOAExtraDeck, errChan chan error) {
	var err error
	data, err = repository.CreateExtraDeck(data)
	if err != nil {
		err = errors.New("failed to create extra deck card: " + err.Error())
		errChan <- err
		return
	}

	result <- data
}

func CreateSideDeckCard(data models.YOASideDeck, result chan models.YOASideDeck, errChan chan error) {
	var err error
	data, err = repository.CreateSideDeck(data)
	if err != nil {
		err = errors.New("failed to create side deck card: " + err.Error())
		errChan <- err
		return
	}

	result <- data
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

func GetDeckByID(id, userID uint, preloadFields []string) (data models.YOADeck, statusCode int, err error) {
	data, err = repository.GetDeckByID(id, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if data.UserID != userID {
		if !data.IsPublic {
			statusCode = http.StatusNotFound
			err = errors.New("data not found")
		}

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
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
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

func UpdateDeck(id, userID uint, request dto.DeckRequest) (response models.YOADeck, statusCode int, err error) {
	data, err := repository.GetDeckByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if data.UserID != userID {
		statusCode = http.StatusNotFound
		err = errors.New("data not found")

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
		for _, item := range data.MainDeck {
			request.MainDeckCardID = append(request.MainDeckCardID, item.ID)
		}
	}
	if len(request.ExtraDeckCardID) == 0 {
		for _, item := range data.ExtraDeck {
			request.ExtraDeckCardID = append(request.ExtraDeckCardID, item.ID)
		}
	}
	if len(request.SideDeckCardID) == 0 {
		for _, item := range data.SideDeck {
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

	response, err = repository.UpdateDeck(data)
	if err != nil {
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	if len(request.MainDeckCardID) > 0 {
		data, _, _, _ := repository.GetMainDecks(dto.FindParameter{
			Filter:       "deleted_at IS NULL AND deck_id = ?",
			FilterValues: []any{response.ID},
		})
		if len(data) > 0 {
			for _, item := range data {
				go repository.DeleteMainDeck(item)
			}
		}

		mainDeckChan := make([]chan models.YOAMainDeck, len(mainDeckCards))
		mainDeckErrChan := make([]chan error, len(mainDeckCards))
		for i, item := range mainDeckCards {
			mainDeckChan[i] = make(chan models.YOAMainDeck)
			mainDeckErrChan[i] = make(chan error)
			go CreateMainDeckCard(models.YOAMainDeck{
				DeckID: response.ID,
				CardID: item.ID,
			}, mainDeckChan[i], mainDeckErrChan[i])
		}

		for i := range mainDeckCards {
			select {
			case err = <-mainDeckErrChan[i]:
				if err != nil {
					err = errors.New("failed to create main deck: " + err.Error())
					statusCode = http.StatusInternalServerError
					return
				}
			case mainDeck := <-mainDeckChan[i]:
				response.MainDeck = append(response.MainDeck, mainDeck)
			}
		}
	}

	if len(request.ExtraDeckCardID) > 0 {
		data, _, _, _ := repository.GetExtraDecks(dto.FindParameter{
			Filter:       "deleted_at IS NULL AND deck_id = ?",
			FilterValues: []any{response.ID},
		})
		if len(data) > 0 {
			for _, item := range data {
				go repository.DeleteExtraDeck(item)
			}
		}

		extraDeckChan := make([]chan models.YOAExtraDeck, len(extraDeckCards))
		extraDeckErrChan := make([]chan error, len(extraDeckCards))
		for i, item := range extraDeckCards {
			extraDeckChan[i] = make(chan models.YOAExtraDeck)
			extraDeckErrChan[i] = make(chan error)
			go CreateExtraDeckCard(models.YOAExtraDeck{
				DeckID: response.ID,
				CardID: item.ID,
			}, extraDeckChan[i], extraDeckErrChan[i])
		}

		for i := range extraDeckCards {
			select {
			case err = <-extraDeckErrChan[i]:
				if err != nil {
					err = errors.New("failed to create extra deck: " + err.Error())
					statusCode = http.StatusInternalServerError
					return
				}
			case extraDeck := <-extraDeckChan[i]:
				response.ExtraDeck = append(response.ExtraDeck, extraDeck)
			}
		}
	}

	if len(request.SideDeckCardID) > 0 {
		data, _, _, _ := repository.GetSideDecks(dto.FindParameter{
			Filter:       "deleted_at IS NULL AND deck_id = ?",
			FilterValues: []any{response.ID},
		})
		if len(data) > 0 {
			for _, item := range data {
				go repository.DeleteSideDeck(item)
			}
		}

		sideDeckChan := make([]chan models.YOASideDeck, len(sideDeckCards))
		sideDeckErrChan := make([]chan error, len(sideDeckCards))
		for i, item := range sideDeckCards {
			sideDeckChan[i] = make(chan models.YOASideDeck)
			sideDeckErrChan[i] = make(chan error)
			go CreateSideDeckCard(models.YOASideDeck{
				DeckID: response.ID,
				CardID: item.ID,
			}, sideDeckChan[i], sideDeckErrChan[i])
		}

		for i := range sideDeckCards {
			select {
			case err = <-sideDeckErrChan[i]:
				if err != nil {
					err = errors.New("failed to create side deck: " + err.Error())
					statusCode = http.StatusInternalServerError
					return
				}
			case sideDeck := <-sideDeckChan[i]:
				response.SideDeck = append(response.SideDeck, sideDeck)
			}
		}
	}

	statusCode = http.StatusOK
	return
}

func DeleteDeck(id, userID uint) (statusCode int, err error) {
	data, err := repository.GetDeckByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}
	if data.UserID != userID {
		statusCode = http.StatusNotFound
		err = errors.New("data not found")

		return
	}

	mainDeck, _, _, _ := repository.GetMainDecks(dto.FindParameter{
		Filter:       "deleted_at IS NULL AND deck_id = ?",
		FilterValues: []any{data.ID},
	})
	if len(mainDeck) > 0 {
		for _, item := range mainDeck {
			go repository.DeleteMainDeck(item)
		}
	}

	extraDeck, _, _, _ := repository.GetExtraDecks(dto.FindParameter{
		Filter:       "deleted_at IS NULL AND deck_id = ?",
		FilterValues: []any{data.ID},
	})
	if len(extraDeck) > 0 {
		for _, item := range extraDeck {
			go repository.DeleteExtraDeck(item)
		}
	}

	sideDeck, _, _, _ := repository.GetSideDecks(dto.FindParameter{
		Filter:       "deleted_at IS NULL AND deck_id = ?",
		FilterValues: []any{data.ID},
	})
	if len(sideDeck) > 0 {
		for _, item := range sideDeck {
			go repository.DeleteSideDeck(item)
		}
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

func ExportDeck(useName, useGroup bool, deckID, userID uint) (file string, statusCode int, err error) {
	deck, statusCode, err := GetDeckByID(deckID, userID, []string{"MainDeck", "ExtraDeck", "SideDeck", "MainDeck.Card", "ExtraDeck.Card", "SideDeck.Card"})
	if err != nil {
		return
	}

	file = ConvertToYDK(deck, useName, useGroup)

	return
}

func ConvertToYDK(deck models.YOADeck, useName, useGroup bool) string {
	var sb strings.Builder

	formatCard := func(cardNameOrID string, count int) string {
		if useGroup {
			return fmt.Sprintf("%dx %s\n", count, cardNameOrID)
		}

		return strings.Repeat(fmt.Sprintf("%s\n", cardNameOrID), count)
	}

	addCardsToYDK := func(cards []models.YOACard) {
		cardCount := make(map[string]int)
		for _, card := range cards {
			identifier := strconv.FormatUint(uint64(card.ID), 10)
			if useName {
				identifier = card.Name
			}
			cardCount[identifier]++
		}

		for card, count := range cardCount {
			sb.WriteString(formatCard(card, count))
		}
	}

	// Main Deck
	var mainDeckCards []models.YOACard
	for _, item := range deck.MainDeck {
		mainDeckCards = append(mainDeckCards, *item.Card)
	}
	sb.WriteString("#main\n")
	addCardsToYDK(mainDeckCards)

	// Extra Deck
	var extraDeckCards []models.YOACard
	for _, item := range deck.ExtraDeck {
		extraDeckCards = append(extraDeckCards, *item.Card)
	}
	sb.WriteString("\n#extra\n")
	addCardsToYDK(extraDeckCards)

	// Side Deck
	var sideDeckCards []models.YOACard
	for _, item := range deck.SideDeck {
		sideDeckCards = append(sideDeckCards, *item.Card)
	}
	sb.WriteString("\n#side\n")
	addCardsToYDK(sideDeckCards)

	return sb.String()
}
