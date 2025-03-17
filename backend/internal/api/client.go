package api

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"time"

	"github.com/JhonnyPG/StockAnalyzer/internal/models"
)

// APIResponse is a struct that represents the response from the API
type APIResponse struct {
	Items    []models.Stock `json:"items"`
	NextPage string         `json:"next_page"`
}

// Client is a struct that represents the client
type Client struct {
	BaseURL    string
	AuhtToken  string
	HTTPClient *http.Client
}

func NewClient(baseURL, authToken string) *Client {
	return &Client{
		BaseURL:    baseURL,
		AuhtToken:  authToken,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetStocks(nextPage string) (*APIResponse, error) {
	url := c.BaseURL
	if nextPage != "" {
		url = fmt.Sprintf("%s?next_page=%s", url, nextPage)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)

	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuhtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned Error status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &apiResp, nil
}

func (c *Client) GetAllStock() ([]models.Stock, error) {
	var allItems []models.Stock
	nextPage := ""

	for {
		resp, err := c.GetStocks(nextPage)
		if err != nil {
			return nil, err
		}

		allItems = append(allItems, resp.Items...)

		if resp.NextPage == "" {
			break
		}

		nextPage = resp.NextPage
	}

	return allItems, nil
}
