// internal/api/cat_api.go
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spy-cats/internal/models"
)

type CatApi struct {
	baseURL string
	apiKey  string
}

func NewCatAPI(baseURL, apiKey string) *CatApi {
	return &CatApi{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c *CatApi) IsValidBreed(breedName string) (bool, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", c.baseURL, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("error fetching breeds: %s", resp.Status)
	}

	var breeds []models.CatBreed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return true, nil
		}
	}

	return false, nil
}
