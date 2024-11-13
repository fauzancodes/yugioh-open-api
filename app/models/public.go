package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomGormModel struct {
	ID        uint            `gorm:"type:int8;column:id;primaryKey" json:"id"`
	CreatedAt time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}

// type YOAUsedApiKey struct {
// 	CustomGormModel
// 	PublicKey    string `json:"-" gorm:"type: text"`
// 	Base64Key    string `json:"-" gorm:"type: text"`
// 	ReceivedHMAC string `json:"-" gorm:"type: text"`
// 	ExpectedHMAC string `json:"-" gorm:"type: text"`
// }

// func (YOAUsedApiKey) TableName() string {
// 	return "yoa_used_api_keys"
// }
