package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/middlewares"
	"github.com/fauzancodes/yugioh-open-api/app/pkg/bcrypt"
	webToken "github.com/fauzancodes/yugioh-open-api/app/pkg/jwt"
	"github.com/fauzancodes/yugioh-open-api/app/service"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
	"github.com/fauzancodes/yugioh-open-api/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var request dto.UserRequest
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

	param := utils.PopulatePaging(c, "")
	_, check, _, _ := service.GetUsers(request.Username, param, []string{})
	if len(check) > 0 {
		return c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Username has been taken",
				Error:   "",
			},
		)
	}

	result, statusCode, err := service.CreateUser(dto.UserRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to register",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to register",
			Data:    result,
		},
	)
}

func Login(c echo.Context) error {
	var request dto.UserRequest
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

	param := utils.PopulatePaging(c, "")
	_, user, statusCode, _ := service.GetUsers(request.Username, param, []string{})
	if len(user) == 0 {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Username not found",
			},
		)
	}

	err := bcrypt.VerifyPassword(request.Password, user[0].Password)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Failed to verify password",
				Error:   err.Error(),
			},
		)
	}

	claims := jwt.MapClaims{}
	claims["id"] = user[0].ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := webToken.GenerateToken(&claims)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			dto.Response{
				Status:  401,
				Message: "Failed to generate jwt token",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to login",
			Data:    token,
		},
	)
}

func GetCurrentUser(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	data, statusCode, err := service.GetUserByID(userID, []string{
		"Decks",
		"Decks.Cards",
	})
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Data not found",
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

func UpdateProfile(c echo.Context) error {
	var request dto.UserRequest
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

	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	data, statusCode, err := service.UpdateUser(userID, request)
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
		http.StatusOK,
		dto.Response{
			Status:  200,
			Message: "Success to update data",
			Data:    data,
		},
	)
}

func RemoveAccount(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)
	log.Printf("Current user ID: %v", userID)

	statusCode, err := service.DeleteUser(userID)
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

func GenerateApiKey(c echo.Context) error {
	userID := utils.GetCurrentUserID(c)

	user, statusCode, err := service.GetUserByID(userID, []string{})
	if err != nil {
		return c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get user",
				Error:   err.Error(),
			},
		)
	}

	publicKey, secretKey, err := middlewares.GenerateRSAKeys()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to generate api key",
				Error:   err.Error(),
			},
		)
	}

	user.SecretKey = secretKey
	user.PublicKey = publicKey

	user, err = repository.UpdateUser(user)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update user",
				Error:   err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		dto.Response{
			Status:  http.StatusOK,
			Message: "Success to generate api key. How to use: 1) Calculate the HMAC signature using the SHA256 method from the public_key and secret_key. 2) Pair the public_key and signature with the pattern public_key:signature. 3) Endcode the pattern using the base64 method. 4) Put the encoded string into the http header x-api-key. Make sure you save this public_key and secret_key in a safe place, because you will not be able to see it again. Never share your secret_key with anyone.",
			Data: echo.Map{
				"public_key": user.PublicKey,
				"secret_key": user.SecretKey,
			},
		},
	)
}
