package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
)

func CreateDeck(data models.YOADeck, mainDeckCards []models.YOACard, extraDeckCards []models.YOACard, sideDeckCards []models.YOACard) (models.YOADeck, error) {
	err := config.DB.Create(&data).Error
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("MainDeckCards").Append(mainDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("ExtraDeckCards").Append(extraDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("SideDeckCards").Append(sideDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	return data, err
}

func GetDeckByID(id uint, preloadFields []string) (response models.YOADeck, err error) {
	db := utils.BuildPreload(config.DB, preloadFields)

	err = db.Where("id = ?", id).First(&response).Error

	return
}

func GetDecks(param dto.FindParameter, preloadFields []string) (responses []models.YOADeck, total int64, totalFiltered int64, err error) {
	err = config.DB.Model(responses).Where(param.BaseFilter, param.BaseFilterValues...).Count(&total).Error
	if err != nil {
		return
	}

	err = config.DB.Model(responses).Where(param.Filter, param.FilterValues...).Count(&totalFiltered).Error
	if err != nil {
		return
	}

	db := utils.BuildPreload(config.DB, preloadFields)

	if param.Limit == 0 {
		err = db.Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	} else {
		err = db.Limit(param.Limit).Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	}

	return
}

func UpdateDeck(data models.YOADeck, mainDeckCards []models.YOACard, extraDeckCards []models.YOACard, sideDeckCards []models.YOACard) (models.YOADeck, error) {
	err := config.DB.Save(&data).Error
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("MainDeckCards").Clear()
	if err != nil {
		return models.YOADeck{}, err
	}
	err = config.DB.Model(&data).Association("MainDeckCards").Append(mainDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("ExtraDeckCards").Clear()
	if err != nil {
		return models.YOADeck{}, err
	}
	err = config.DB.Model(&data).Association("ExtraDeckCards").Append(extraDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	err = config.DB.Model(&data).Association("SideDeckCards").Clear()
	if err != nil {
		return models.YOADeck{}, err
	}
	err = config.DB.Model(&data).Association("SideDeckCards").Append(sideDeckCards)
	if err != nil {
		return models.YOADeck{}, err
	}

	return data, err
}

func DeleteDeck(data models.YOADeck) error {
	err := config.DB.Delete(&data).Error

	return err
}
