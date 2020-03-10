package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	weatherByCityNameURL = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"
)

type OpenWeatherMapResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
	Message  string `json:"message"`
}

type APIClient struct {
	apiKey string
	Fetch  Fetcher
}

type Fetcher func(string) (string, error)
type Adapter func(fetcherFunc Fetcher) Fetcher

func NewAPIClient(apiKey string, fetcher Fetcher) *APIClient {
	return &APIClient{
		apiKey: apiKey,
		Fetch:  fetcher,
	}
}

func (c *APIClient) GetWeatherByCityName(cityName string) (*OpenWeatherMapResponse, error) {
	str, err := c.Fetch(fmt.Sprintf(weatherByCityNameURL, cityName, c.apiKey))
	if err != nil {
		return nil, err
	}

	return convertStringToOpenWeatherMapResponse(str)
}

func convertStringToOpenWeatherMapResponse(s string) (*OpenWeatherMapResponse, error) {
	a := &OpenWeatherMapResponse{}
	if err := json.NewDecoder(strings.NewReader(s)).Decode(a); err != nil {
		return nil, err
	}

	return a, nil
}

func HttpGetFetcher(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(rawBody), nil
}
