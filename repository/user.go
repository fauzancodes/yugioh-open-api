package repository

import (
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
)

func CreateUser(data models.YOAUser) (models.YOAUser, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetUserByID(id uint, preloadFields []string) (response models.YOAUser, err error) {
	db := utils.BuildPreload(config.DB, preloadFields)

	err = db.Where("id = ?", id).First(&response).Error

	return
}

func GetUsers(param dto.FindParameter, preloadFields []string) (responses []models.YOAUser, total int64, totalFiltered int64, err error) {
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

func UpdateUser(data models.YOAUser) (models.YOAUser, error) {
	err := config.DB.Save(&data).Error

	return data, err
}

func DeleteUser(data models.YOAUser) error {
	err := config.DB.Delete(&data).Error

	return err
}
