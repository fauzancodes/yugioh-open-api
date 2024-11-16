package dto

import (
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type DeckRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MainDeckCardID  []uint `json:"main_deck_card_id"`
	ExtraDeckCardID []uint `json:"extra_deck_card_id"`
	SideDeckCardID  []uint `json:"side_deck_card_id"`
	IsPublic        bool   `json:"is_public"`
}

func (request DeckRequest) Validate() error {
	if len(request.MainDeckCardID) < 40 {
		return errors.New("minimum number of cards in the main deck: 40. current: " + strconv.Itoa(len(request.MainDeckCardID)))
	}
	if len(request.MainDeckCardID) > 60 {
		return errors.New("maximum number of cards in the main deck: 60. current: " + strconv.Itoa(len(request.MainDeckCardID)))
	}
	if len(request.ExtraDeckCardID) > 15 {
		return errors.New("maximum number of cards in the extra deck: 15. current: " + strconv.Itoa(len(request.ExtraDeckCardID)))
	}
	if len(request.SideDeckCardID) > 15 {
		return errors.New("maximum number of cards in the side deck: 15. current: " + strconv.Itoa(len(request.SideDeckCardID)))
	}

	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
	)
}
