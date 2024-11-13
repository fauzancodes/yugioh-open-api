package middlewares

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/fauzancodes/yugioh-open-api/app/models"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Secure() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            0,
		ContentSecurityPolicy: "",
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/docs")
		},
	})
}

func StripHTMLMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for key, values := range c.QueryParams() {
			for i, value := range values {
				sanitizedValue := template.HTMLEscapeString(value)
				sanitizedValue = strings.ReplaceAll(sanitizedValue, "=", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, "<", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, ">", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, "*", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " AND ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " OR ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " and ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " or ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " || ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, " && ", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, "'", "")
				sanitizedValue = strings.ReplaceAll(sanitizedValue, "&#39;", "")
				values[i] = strip.StripTags(sanitizedValue)
			}
			c.QueryParams()[key] = values
		}

		return next(c)
	}
}

func CheckAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if config.LoadConfig().EnableAPIKey {
			apiKey := c.Request().Header.Get("X-API-KEY")
			if apiKey == "" {
				fmt.Println("Failed to check api key: no api key in header")
				return c.JSON(http.StatusForbidden, dto.Response{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
			}
			if apiKey == config.LoadConfig().SpecialApiKey {
				return next(c)
			}

			publicKey, receivedHMAC, err := DecodeAPIKeyBase64(apiKey)
			if err != nil {
				fmt.Println("Failed to check api key: ", err.Error())
				return c.JSON(http.StatusForbidden, dto.Response{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
			}

			var user models.YOAUser
			err = config.DB.Where("public_key = ?", publicKey).First(&user).Error
			if err != nil {
				fmt.Println("Failed to get user by public_key:", err.Error())
				return c.JSON(http.StatusForbidden, dto.Response{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
			}

			hmacVerified, _, err := VerifyAPIKeyHMAC(publicKey, receivedHMAC, user.SecretKey)
			if err != nil {
				fmt.Println("Failed to check api key: ", err.Error())
				return c.JSON(http.StatusForbidden, dto.Response{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
			}

			if !hmacVerified {
				fmt.Println("Failed to check api key: failed to verify hmac")
				return c.JSON(http.StatusForbidden, dto.Response{
					Status:  http.StatusForbidden,
					Message: "Forbidden",
				})
			}
		}

		return next(c)
	}
}

func ComputeAPIKeyHMAC(publicKey, secretKey string) (response string, err error) {
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err = h.Write([]byte(publicKey))
	response = hex.EncodeToString(h.Sum(nil))

	return
}

func DecodeAPIKeyBase64(encodedKey string) (publicKey string, hmacSignature string, err error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		err = errors.New("invalid encoded key")
		return
	}

	publicKey = parts[0]
	hmacSignature = parts[1]

	return
}

func VerifyAPIKeyHMAC(publicKey, receivedHMAC, secretKey string) (response bool, expectedHMAC string, err error) {
	expectedHMAC, err = ComputeAPIKeyHMAC(publicKey, secretKey)
	if err != nil {
		return
	}
	response = hmac.Equal([]byte(expectedHMAC), []byte(receivedHMAC))

	return
}

func GenerateRSAKeys() (publicKey string, secretKey string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	secretKey = base64.StdEncoding.EncodeToString(privateKeyPEM)

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	publicKey = base64.StdEncoding.EncodeToString(publicKeyPEM)

	return publicKey, secretKey, nil
}
