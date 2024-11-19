package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/service"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/labstack/echo/v4"
)

func CreateDeck(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	var request dto.DeckRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)
	}

	if err := request.Validate(); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)
	}

	result, statusCode, err := service.CreateDeck(userID, request)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to create",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to create",
			Data:    result,
		},
	)
}

func GetDecks(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	preloadFields := utils.GetBuildPreloadFields(c)

	param := utils.PopulatePaging(c, "")
	data, _, statusCode, err := service.GetDecks(userID, param, preloadFields)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get data",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(statusCode, data)
}

func GetDeckByID(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	preloadFields := utils.GetBuildPreloadFields(c)

	data, statusCode, err := service.GetDeckByID(uint(id), userID, preloadFields)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get data",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to get data",
			Data:    data,
		},
	)
}

func UpdateDeck(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var request dto.DeckRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)
	}

	if err := request.Validate(); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)
	}

	data, statusCode, err := service.UpdateDeck(uint(id), userID, request)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to update data",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to update data",
			Data:    data,
		},
	)
}

func DeleteDeck(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	statusCode, err := service.DeleteDeck(uint(id), userID)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to delete data",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to delete data",
		},
	)
}

func GetPublicDecks(c echo.Context) error {
	preloadFields := utils.GetBuildPreloadFields(c)

	param := utils.PopulatePaging(c, "")
	param.Custom = "true"
	data, _, statusCode, err := service.GetDecks(0, param, preloadFields)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get data",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(statusCode, data)
}

func ExportDeckByID(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	useName := strings.ToLower(c.QueryParam("identifier")) == "name"
	useGroup, _ := strconv.ParseBool(c.QueryParam("group_copy"))

	data, statusCode, err := service.ExportDeck(useName, useGroup, uint(id), userID)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get data",
				Error:   err.Error(),
			},
		)
	}

	return c.Blob(http.StatusOK, "application/octet-stream", []byte(data))
}
