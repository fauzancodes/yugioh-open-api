package models

type YOADeck struct {
	CustomGormModel
	Name               string        `json:"name" gorm:"type:varchar(255);column:name"`
	Description        string        `json:"description" gorm:"type:text;column:description"`
	TotalDeckCard      uint          `json:"total_deck_card" gorm:"type:int8;column:total_deck_card"`
	TotalMainDeckCard  uint          `json:"total_main_deck_card" gorm:"type:int8;column:total_main_deck_card"`
	TotalExtraDeckCard uint          `json:"total_extra_deck_card" gorm:"type:int8;column:total_extra_deck_card"`
	TotalSideDeckCard  uint          `json:"total_side_deck_card" gorm:"type:int8;column:total_side_deck_card"`
	TotalMonsterCard   uint          `json:"total_monster_card" gorm:"type:int8;column:total_monster_card"`
	TotalSpellCard     uint          `json:"total_spell_card" gorm:"type:int8;column:total_spell_card"`
	TotalTrapCard      uint          `json:"total_trap_card" gorm:"type:int8;column:total_trap_card"`
	MainDeckCards      []YOACard     `json:"main_deck_cards" gorm:"many2many:yoa_main_deck_cards"`
	ExtraDeckCards     []YOACard     `json:"extra_deck_cards" gorm:"many2many:yoa_extra_deck_cards"`
	SideDeckCards      []YOACard     `json:"side_deck_cards" gorm:"many2many:yoa_side_deck_cards"`
	IsPublic           bool          `json:"is_public" gorm:"type:bool;column:is_public"`
	UserID             uint          `json:"user_id" gorm:"type:int8;column:user_id"`
	User               *UserRelation `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (YOADeck) TableName() string {
	return "yoa_decks"
}

type DeckRelation struct {
	CustomGormModel
	Name               string    `json:"name" gorm:"type:varchar(255);column:name"`
	Description        string    `json:"description" gorm:"type:text;column:description"`
	TotalDeckCard      uint      `json:"total_deck_card" gorm:"type:int8;column:total_deck_card"`
	TotalMainDeckCard  uint      `json:"total_main_deck_card" gorm:"type:int8;column:total_main_deck_card"`
	TotalExtraDeckCard uint      `json:"total_extra_deck_card" gorm:"type:int8;column:total_extra_deck_card"`
	TotalSideDeckCard  uint      `json:"total_side_deck_card" gorm:"type:int8;column:total_side_deck_card"`
	TotalMonsterCard   uint      `json:"total_monster_card" gorm:"type:int8;column:total_monster_card"`
	TotalSpellCard     uint      `json:"total_spell_card" gorm:"type:int8;column:total_spell_card"`
	TotalTrapCard      uint      `json:"total_monster_trap_card" gorm:"type:int8;column:total_monster_trap_card"`
	MainDeckCards      []YOACard `json:"main_deck_cards" gorm:"many2many:yoa_main_deck_cards"`
	ExtraDeckCards     []YOACard `json:"extra_deck_cards" gorm:"many2many:yoa_extra_deck_cards"`
	SideDeckCards      []YOACard `json:"side_deck_cards" gorm:"many2many:yoa_side_deck_cards"`
	UserID             uint      `json:"-" gorm:"column:user_id"`
}

func (DeckRelation) TableName() string {
	return "yoa_decks"
}
