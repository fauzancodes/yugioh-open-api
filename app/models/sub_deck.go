package models

type YOAMainDeckCard struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:yoa_deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:yoa_card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOAMainDeckCard) TableName() string {
	return "yoa_main_deck_cards"
}

type YOAExtraDeckCard struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:yoa_deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:yoa_card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOAExtraDeckCard) TableName() string {
	return "yoa_extra_deck_cards"
}

type YOASideDeckCard struct {
	CustomGormModel
	DeckID uint     `json:"deck_id" gorm:"type:int8;column:yoa_deck_id"`
	CardID uint     `json:"card_id" gorm:"type:int8;column:yoa_card_id"`
	Card   *YOACard `json:"card,omitempty" gorm:"foreignKey:CardID"`
}

func (YOASideDeckCard) TableName() string {
	return "yoa_side_deck_cards"
}
