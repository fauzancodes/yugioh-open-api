package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CheckURLStatus(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer resp.Body.Close()

	// Cek status kode apakah 200
	return resp.StatusCode == http.StatusOK
}

func BuildPreload(db *gorm.DB, fields []string) *gorm.DB {
	if len(fields) > 0 {
		for _, field := range fields {
			db = db.Preload(field)
		}
	}

	return db
}

func GetBuildPreloadFields(c echo.Context) (fields []string) {
	raw := c.QueryParam("preload_fields")

	if raw != "" {
		fields = strings.Split(raw, ",")
	}

	return
}

func GetCurrentUserID(c echo.Context) uint {
	userID := c.Get("currentUser").(jwt.MapClaims)["id"].(float64)

	return uint(userID)
}

func RemoveDuplicatesFromStringArray(strings []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string

	for _, str := range strings {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	return result
}

func GetDuplicatesMoreThanThree(nums []uint) []uint {
	freqMap := make(map[uint]int)
	duplicates := []uint{}

	for _, num := range nums {
		freqMap[num]++
	}

	for num, count := range freqMap {
		if count > 3 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}
