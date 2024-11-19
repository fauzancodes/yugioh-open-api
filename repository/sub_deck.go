package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
)

func CreateMainDeck(data models.YOAMainDeck) (models.YOAMainDeck, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetMainDeckByID(id uint) (response models.YOAMainDeck, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetMainDecks(param dto.FindParameter) (responses []models.YOAMainDeck, total int64, totalFiltered int64, err error) {
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

func UpdateMainDeck(data models.YOAMainDeck) (models.YOAMainDeck, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteMainDeck(data models.YOAMainDeck) error {
	err := config.DB.Delete(&data).Error

	return err
}

func CreateExtraDeck(data models.YOAExtraDeck) (models.YOAExtraDeck, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetExtraDeckByID(id uint) (response models.YOAExtraDeck, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetExtraDecks(param dto.FindParameter) (responses []models.YOAExtraDeck, total int64, totalFiltered int64, err error) {
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

func UpdateExtraDeck(data models.YOAExtraDeck) (models.YOAExtraDeck, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteExtraDeck(data models.YOAExtraDeck) error {
	err := config.DB.Delete(&data).Error

	return err
}

func CreateSideDeck(data models.YOASideDeck) (models.YOASideDeck, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetSideDeckByID(id uint) (response models.YOASideDeck, err error) {
	err = config.DB.Where("id = ?", id).First(&response).Error

	return
}

func GetSideDecks(param dto.FindParameter) (responses []models.YOASideDeck, total int64, totalFiltered int64, err error) {
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

func UpdateSideDeck(data models.YOASideDeck) (models.YOASideDeck, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteSideDeck(data models.YOASideDeck) error {
	err := config.DB.Delete(&data).Error

	return err
}
