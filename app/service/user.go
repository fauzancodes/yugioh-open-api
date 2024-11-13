package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	"github.com/fauzancodes/yugioh-open-api/app/pkg/bcrypt"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/fauzancodes/yugioh-open-api/repository"
	"gorm.io/gorm"
)

func CreateUser(request dto.UserRequest) (response models.YOAUser, statusCode int, err error) {
	data := models.YOAUser{
		Username: request.Username,
		Password: bcrypt.HashPassword(request.Password),
	}

	response, err = repository.CreateUser(data)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusCreated
	return
}

func GetUserByID(id uint, preloadFields []string) (data models.YOAUser, statusCode int, err error) {
	data, err = repository.GetUserByID(id, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func GetUsers(username string, param utils.PagingRequest, preloadFields []string) (response utils.PagingResponse, data []models.YOAUser, statusCode int, err error) {
	baseFilter := "deleted_at IS NULL"
	filter := baseFilter
	var filterValues []any

	if username != "" {
		filter += " AND username = ?"
		filterValues = append(filterValues, username)
	}
	if param.Search != "" {
		filter += " AND (first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?)"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
	}

	data, total, totalFiltered, err := repository.GetUsers(dto.FindParameter{
		BaseFilter:   baseFilter,
		Filter:       filter,
		FilterValues: filterValues,
		Limit:        param.Limit,
		Order:        param.Order,
		Offset:       param.Offset,
	}, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	response = utils.PopulateResPaging(&param, data, total, totalFiltered)

	statusCode = http.StatusOK
	return
}

func UpdateUser(id uint, request dto.UserRequest) (response models.YOAUser, statusCode int, err error) {
	data, err := repository.GetUserByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	if request.Username != "" {
		data.Username = request.Username
	}
	if request.Password != "" {
		data.Password = bcrypt.HashPassword(request.Password)
	}

	response, err = repository.UpdateUser(data)
	if err != nil {
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func DeleteUser(id uint) (statusCode int, err error) {
	data, err := repository.GetUserByID(id, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	err = repository.DeleteUser(data)
	if err != nil {
		err = errors.New("failed to delete data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}
