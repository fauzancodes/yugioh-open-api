package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
)

func CreateMainDeckCard(data models.YOAMainDeckCard) (models.YOAMainDeckCard, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetMainDeckCardByID(id uint) (response models.YOAMainDeckCard, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetMainDeckCards(param dto.FindParameter) (responses []models.YOAMainDeckCard, total int64, totalFiltered int64, err error) {
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

func UpdateMainDeckCard(data models.YOAMainDeckCard) (models.YOAMainDeckCard, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteMainDeckCard(data models.YOAMainDeckCard) error {
	err := config.DB.Delete(&data).Error

	return err
}

func CreateExtraDeckCard(data models.YOAExtraDeckCard) (models.YOAExtraDeckCard, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetExtraDeckCardByID(id uint) (response models.YOAExtraDeckCard, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetExtraDeckCards(param dto.FindParameter) (responses []models.YOAExtraDeckCard, total int64, totalFiltered int64, err error) {
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

func UpdateExtraDeckCard(data models.YOAExtraDeckCard) (models.YOAExtraDeckCard, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteExtraDeckCard(data models.YOAExtraDeckCard) error {
	err := config.DB.Delete(&data).Error

	return err
}

func CreateSideDeckCard(data models.YOASideDeckCard) (models.YOASideDeckCard, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetSideDeckCardByID(id uint) (response models.YOASideDeckCard, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetSideDeckCards(param dto.FindParameter) (responses []models.YOASideDeckCard, total int64, totalFiltered int64, err error) {
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

func UpdateSideDeckCard(data models.YOASideDeckCard) (models.YOASideDeckCard, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteSideDeckCard(data models.YOASideDeckCard) error {
	err := config.DB.Delete(&data).Error

	return err
}
