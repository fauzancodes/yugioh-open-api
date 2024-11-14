package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/service"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/labstack/echo/v4"
)

func CreateCard(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	user, _, _ := service.GetUserByID(userID, []string{})
	if !user.IsAdmin {
		return c.JSON(
			http.StatusForbidden,
			dto.Response{
				Status:  http.StatusForbidden,
				Message: "Only admins are allowed",
			},
		)
	}

	var request dto.CardRequest
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
				Message: "Invalid request value. Please see the card utility at the /v1/card/utility endpoint.",
				Error:   err.Error(),
			},
		)
	}

	result, statusCode, err := service.CreateCard(request)
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

func GetCards(c echo.Context) error {
	cardType := c.QueryParam("card_type")
	race := c.QueryParam("race")
	archetype := c.QueryParam("archetype")
	attribute := c.QueryParam("attribute")
	cardsets := c.QueryParam("cardsets")
	rarity := c.QueryParam("rarity")
	rarityCode := c.QueryParam("rarity_code")
	attack, _ := strconv.Atoi(c.QueryParam("attack"))
	attackMarginTop, _ := strconv.Atoi(c.QueryParam("attack_margin_top"))
	attackMarginBottom, _ := strconv.Atoi(c.QueryParam("attack_margin_bottom"))
	defense, _ := strconv.Atoi(c.QueryParam("defense"))
	defenseMarginTop, _ := strconv.Atoi(c.QueryParam("defense_margin_top"))
	defenseMarginBottom, _ := strconv.Atoi(c.QueryParam("defense_margin_bottom"))
	level, _ := strconv.Atoi(c.QueryParam("level"))
	levelMarginTop, _ := strconv.Atoi(c.QueryParam("level_margin_top"))
	levelMarginBottom, _ := strconv.Atoi(c.QueryParam("level_margin_bottom"))

	param := utils.PopulatePaging(c, "")
	data, _, statusCode, err := service.GetCards(cardType, race, archetype, attribute, cardsets, rarity, rarityCode, attack, attackMarginTop, attackMarginBottom, defense, defenseMarginTop, defenseMarginBottom, level, levelMarginTop, levelMarginBottom, param)
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

func GetCardByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, statusCode, err := service.GetCardByID(uint(id))
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

func UpdateCard(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	user, _, _ := service.GetUserByID(userID, []string{})
	if !user.IsAdmin {
		return c.JSON(
			http.StatusForbidden,
			dto.Response{
				Status:  http.StatusForbidden,
				Message: "Only admins are allowed",
			},
		)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var request dto.CardRequest
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
				Message: "Invalid request value. Please see the card utility at the /v1/card/utility endpoint.",
				Error:   err.Error(),
			},
		)
	}

	data, statusCode, err := service.UpdateCard(uint(id), request)
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

func DeleteCard(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	user, _, _ := service.GetUserByID(userID, []string{})
	if !user.IsAdmin {
		return c.JSON(
			http.StatusForbidden,
			dto.Response{
				Status:  http.StatusForbidden,
				Message: "Only admins are allowed",
			},
		)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	statusCode, err := service.DeleteCard(uint(id))
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

func GetCardUtility(c echo.Context) error {
	field := c.QueryParam("field")

	data, statusCode, err := service.GetCardUtility(field)
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
