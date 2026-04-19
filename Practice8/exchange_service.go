package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RateResponse struct {
	Rate     float64 `json:"rate"`
	ErrorMsg string  `json:"error,omitempty"`
}

type ExchangeService struct {
	BaseURL string
	Client  *http.Client
}

func NewExchangeService(url string) *ExchangeService {
	return &ExchangeService{BaseURL: url, Client: &http.Client{}}
}

func (s *ExchangeService) GetRate(from, to string) (float64, error) {
	url := fmt.Sprintf("%s/convert?from=%s&to=%s", s.BaseURL, from, to)
	resp, err := s.Client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var res RateResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("api error: %s", res.ErrorMsg)
	}
	return res.Rate, nil
}
