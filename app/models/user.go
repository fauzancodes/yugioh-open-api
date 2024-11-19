package models

type YOAUser struct {
	CustomGormModel
	Username  string    `json:"username" gorm:"type:varchar(255);column:username"`
	Password  string    `json:"-" gorm:"type:varchar(255);column:password"`
	SecretKey string    `json:"-" gorm:"type:text;column:secret_key"`
	PublicKey string    `json:"-" gorm:"type:text;column:public_key"`
	IsAdmin   bool      `json:"-" gorm:"type:bool;column:is_admin"`
	Decks     []YOADeck `json:"decks,omitempty" gorm:"foreignKey:UserID"`
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
