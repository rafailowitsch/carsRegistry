package integration

import (
	"carsRegistry/internal/domain"
	"carsRegistry/pkg/logg"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CarsInfoClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCarsInfoClient(baseURL string) *CarsInfoClient {
	return &CarsInfoClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *CarsInfoClient) FetchCarInfo(regNum string) (*domain.CarsInfo, error) {
	logg.LogInfo("integration:FetchCarInfo", "Fetching car info", logrus.Fields{"regNum": regNum})
	if c.baseURL == "" || regNum == "" {
		return &domain.CarsInfo{}, fmt.Errorf("baseURL or regNum is empty")
	}

	url := c.baseURL + "/info?regNum=" + regNum
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var carsInfo domain.CarsInfo
	if err := json.NewDecoder(resp.Body).Decode(&carsInfo); err != nil {
		return nil, err
	}
	return &carsInfo, nil
}
