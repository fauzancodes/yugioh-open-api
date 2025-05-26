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
	var cardType []string
	cardTypeParam := c.QueryParam("card_type")
	if cardTypeParam != "" {
		cardType = strings.Split(strings.ReplaceAll(cardTypeParam, ", ", ","), ",")
	}

	var race []string
	raceParam := c.QueryParam("race")
	if raceParam != "" {
		race = strings.Split(strings.ReplaceAll(raceParam, ", ", ","), ",")
	}

	var archetype []string
	archetypeParam := c.QueryParam("archetype")
	if archetypeParam != "" {
		archetype = strings.Split(strings.ReplaceAll(archetypeParam, ", ", ","), ",")
	}

	var attribute []string
	attributeParam := c.QueryParam("attribute")
	if attributeParam != "" {
		attribute = strings.Split(strings.ReplaceAll(attributeParam, ", ", ","), ",")
	}

	var cardsets []string
	cardsetsParam := c.QueryParam("card_type")
	if cardsetsParam != "" {
		cardsets = strings.Split(strings.ReplaceAll(cardsetsParam, ", ", ","), ",")
	}

	var rarity []string
	rarityParam := c.QueryParam("rarity")
	if rarityParam != "" {
		rarity = strings.Split(strings.ReplaceAll(rarityParam, ", ", ","), ",")
	}

	var rarityCode []string
	rarityCodeParam := c.QueryParam("rarity_code")
	if rarityCodeParam != "" {
		rarityCode = strings.Split(strings.ReplaceAll(rarityCodeParam, ", ", ","), ",")
	}

	attack, _ := strconv.Atoi(c.QueryParam("attack"))
	attackMarginTop, _ := strconv.Atoi(c.QueryParam("attack_margin_top"))
	attackMarginBottom, _ := strconv.Atoi(c.QueryParam("attack_margin_bottom"))
	defense, _ := strconv.Atoi(c.QueryParam("defense"))
	defenseMarginTop, _ := strconv.Atoi(c.QueryParam("defense_margin_top"))
	defenseMarginBottom, _ := strconv.Atoi(c.QueryParam("defense_margin_bottom"))
	levelMarginTop, _ := strconv.Atoi(c.QueryParam("level_margin_top"))
	levelMarginBottom, _ := strconv.Atoi(c.QueryParam("level_margin_bottom"))

	var level []int
	levelParam := c.QueryParam("level")
	if levelParam != "" {
		levelParamManipulated := strings.Split(strings.ReplaceAll(levelParam, ", ", ","), ",")
		for _, item := range levelParamManipulated {
			levelInt, _ := strconv.Atoi(item)
			level = append(level, levelInt)
		}
	}

	param := utils.PopulatePaging(c, "")
	data, _, statusCode, err := service.GetCards(cardType, race, archetype, attribute, cardsets, rarity, rarityCode, level, attack, attackMarginTop, attackMarginBottom, defense, defenseMarginTop, defenseMarginBottom, levelMarginTop, levelMarginBottom, param)
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

func UploadCardPicture(c echo.Context) error {
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

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  500,
				Message: "Failed to get file from form",
				Error:   err.Error(),
			},
		)
	}

	responseURL, statusCode, err := service.UploadCardPicture(file)
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to upload card picture",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to upload",
			Data:    responseURL,
		},
	)
}
