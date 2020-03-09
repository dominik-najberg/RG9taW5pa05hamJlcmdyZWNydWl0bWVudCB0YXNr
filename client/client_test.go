package client

import (
	"errors"
	"reflect"
	"testing"
)

const (
	CorrectLoadLondon = `{"coord":{"lon":-0.13,"lat":51.51},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"base":"stations","main":{"temp":280.39,"feels_like":274.97,"temp_min":279.15,"temp_max":281.48,"pressure":1010,"humidity":87},"visibility":10000,"wind":{"speed":6.2,"deg":200},"rain":{"1h":0.76},"clouds":{"all":75},"dt":1583783703,"sys":{"type":1,"id":1414,"country":"GB","sunrise":1583735250,"sunset":1583776480},"timezone":0,"id":2643743,"name":"London","cod":200}`
)

type fetcherMock struct {
	Response string
	Err      error
}

func (f *fetcherMock) Fetch(_ string) (string, error) {
	return f.Response, f.Err
}

func TestAPIClient_GetWeatherByCityName(t *testing.T) {
	tests := []struct {
		name       string
		fmResponse string
		fmError    error
		want       *OpenWeatherMapResponse
		wantErr    bool
	}{
		{
			name:       "happy path",
			fmResponse: CorrectLoadLondon,
			fmError:    nil,
			want:       generateCorrectOpenWeatherMapResponse(),
			wantErr:    false,
		},
		{
			name:       "error on fetch",
			fmResponse: "",
			fmError:    errors.New("some error"),
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "decoder error",
			fmResponse: "---bad json---",
			fmError:    nil,
			want:       nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := fetcherMock{
				Response: tt.fmResponse,
				Err:      tt.fmError,
			}
			c := NewAPIClient("test-key",fm.Fetch)
			got, err := c.GetWeatherByCityName("city")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeatherByCityName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeatherByCityName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateCorrectOpenWeatherMapResponse() *OpenWeatherMapResponse {
	return &OpenWeatherMapResponse{
		Coord: struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		}{
			Lon: -0.13,
			Lat: 51.51,
		},
		Weather: []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		}{
			{
				ID:          500,
				Main:        "Rain",
				Description: "light rain",
				Icon:        "10n",
			},
		},
		Base: "stations",
		Main: struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
		}{
			Temp:      280.39,
			FeelsLike: 274.97,
			TempMin:   279.15,
			TempMax:   281.48,
			Pressure:  1010,
			Humidity:  87,
		},
		Visibility: 10000,
		Wind: struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		}{
			Speed: 6.2,
			Deg:   200,
		},
		Clouds: struct {
			All int `json:"all"`
		}{
			All: 75,
		},
		Dt: 1583783703,
		Sys: struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
		}{
			Type:    1,
			ID:      1414,
			Message: 0,
			Country: "GB",
			Sunrise: 1583735250,
			Sunset:  1583776480,
		},
		Timezone: 0,
		ID:       2643743,
		Name:     "London",
		Cod:      200,
		Message:  "",
	}
}
