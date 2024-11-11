package models

type YOAUser struct {
	CustomGormModel
	Username  string         `json:"username" gorm:"type:varchar(255);column:username"`
	Password  string         `json:"-" gorm:"type:varchar(255);column:password"`
	ApiSecret string         `json:"-" gorm:"type:varchar(255);column:api_secret"`
	IsAdmin   bool           `json:"is_admin" gorm:"type:bool;column:is_admin"`
	Decks     []DeckRelation `json:"decks,omitempty" gorm:"foreignKey:UserID"`
}

func (YOAUser) TableName() string {
	return "yoa_users"
}

type UserRelation struct {
	CustomGormModel
	Username string `json:"username" gorm:"type:varchar(255);column:username"`
}

func (UserRelation) TableName() string {
	return "yoa_users"
}
