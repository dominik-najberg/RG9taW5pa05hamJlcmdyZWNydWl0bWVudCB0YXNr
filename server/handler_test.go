package server

import (
	"errors"
	"github.com/dominik-najberg/RG9taW5pa05hamJlcmdyZWNydWl0bWVudCB0YXNr/client"
	"github.com/dominik-najberg/RG9taW5pa05hamJlcmdyZWNydWl0bWVudCB0YXNr/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// we expect to display only correct OpenWeatherAPI data
func Test_handlers_weatherByCitiesViewHandler(t *testing.T) {
	var (
		rw    = httptest.NewRecorder()
		appID = "test-app-id"
	)
	req, err := http.NewRequest("GET", "http://localhost/weather/city/?cities=London", &strings.Reader{})
	if err != nil {
		t.Fatalf("error while creating request: %v", err)
	}

	det := DisplayEngineTester{t: t}

	tests := []struct {
		name       string
		fmResponse string
		fmError    error
		want       int
	}{
		{
			name:       "happy path - expected one response struct",
			fmResponse: CorrectLoadLondon,
			fmError:    nil,
			want:       1,
		},
		{
			name:       "remote server error response - nil response expected",
			fmResponse: ErrorMessageLoad,
			fmError:    nil,
			want:       0,
		},
		{
			name:       "error on decoding",
			fmResponse: CorrectLoadLondon,
			fmError:    errors.New("error on fetch - nil response expected"),
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetcher := mocks.FetcherMock{}
			fetcher.On("Fetch", mock.Anything).Return(tt.fmResponse, tt.fmError)

			h := handlers{
				fetcher:    fetcher.Fetch,
				appID:      appID,
				displayEng: det.DisplayEngineFunc,
			}

			det.expected = tt.want

			h.weatherByCitiesViewHandler(rw, req)
		})
	}
}

type DisplayEngineTester struct {
	t        *testing.T
	expected int
}

func (t *DisplayEngineTester) DisplayEngineFunc(_ http.ResponseWriter, actual []*client.OpenWeatherMapResponse) {
	assert.Equal(t.t, t.expected, len(actual))
}
