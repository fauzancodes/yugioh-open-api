package models

type YOAMainDeck struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOAMainDeck) TableName() string {
	return "yoa_main_decks"
}

type YOAExtraDeck struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOAExtraDeck) TableName() string {
	return "yoa_extra_decks"
}

type YOASideDeck struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOASideDeck) TableName() string {
	return "yoa_side_decks"
}
