package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"spy-cats/internal/models"
)

const baseURL = "https://api.thecatapi.com/v1/breeds"

func IsValidBreed(breedName string) (bool, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("x-api-key", os.Getenv("API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("error fetching breeds: %s", resp.Status)
	}

	var breeds []models.CatBreed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return false, err
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return true, nil
		}
	}

	return false, nil
}
