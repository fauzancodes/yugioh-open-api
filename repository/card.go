package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
)

func CreateCard(data models.YOACard) (models.YOACard, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetCardByID(id uint) (response models.YOACard, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetCards(param dto.FindParameter) (responses []models.YOACard, total int64, totalFiltered int64, err error) {
	err = config.DB.Model(responses).Where(param.BaseFilter, param.BaseFilterValues...).Count(&total).Error
	if err != nil {
		return
	}

	err = config.DB.Model(responses).Where(param.Filter, param.FilterValues...).Count(&totalFiltered).Error
	if err != nil {
		return
	}

	if param.Limit == 0 {
		err = config.DB.Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	} else {
		err = config.DB.Limit(param.Limit).Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	}

	return
}

func UpdateCard(data models.YOACard) (models.YOACard, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteCard(data models.YOACard) error {
	err := config.DB.Delete(&data).Error

	return err
}

func GetCardUtility(field string) (responses []string, err error) {
	nullValues := "''"
	if field == "level" {
		nullValues = "0"
	}
	err = config.DB.Raw("SELECT DISTINCT " + field + " FROM " + models.YOACard{}.TableName() + " WHERE " + field + " != " + nullValues + " AND deleted_at IS NULL").Scan(&responses).Error

	return
}
