package models

type YOACard struct {
	CustomGormModel
	Name        string `json:"name" gorm:"column:name;type:varchar(255)"`
	Type        string `json:"type" gorm:"column:type;type:varchar(255)"`
	Description string `json:"description" gorm:"column:description;type:text"`
	Race        string `json:"race" gorm:"column:race;type:varchar(255)"`
	Archetype   string `json:"archetype" gorm:"column:archetype;type:varchar(255)"`
	Attack      int    `json:"attack" gorm:"column:attack;type:int8"`
	Defense     int    `json:"defense" gorm:"column:defense;type:int8"`
	Level       int    `json:"level" gorm:"column:level;type:int8"`
	CardSets    int    `json:"card_sets" gorm:"column:card_sets;varchar(255)"`
	ImageUrl    int    `json:"image_url" gorm:"column:image_url;varchar(255)"`
}

func (YOACard) TableName() string {
	return "yoa_cards"
}
