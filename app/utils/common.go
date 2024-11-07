package utils

import (
	"fmt"
	"net/http"
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
