package server

import "testing"

const (
	CorrectLoadLondon = `{"coord":{"lon":-0.13,"lat":51.51},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"base":"stations","main":{"temp":280.39,"feels_like":274.97,"temp_min":279.15,"temp_max":281.48,"pressure":1010,"humidity":87},"visibility":10000,"wind":{"speed":6.2,"deg":200},"rain":{"1h":0.76},"clouds":{"all":75},"dt":1583783703,"sys":{"type":1,"id":1414,"country":"GB","sunrise":1583735250,"sunset":1583776480},"timezone":0,"id":2643743,"name":"London","cod":200}`
	ErrorMessageLoad  = `{"cod":401, "message": "big bad error"}`
)

func Test_validateResponseCode(t *testing.T) {
	tests := []struct {
		name string
		json string
		want bool
	}{
		{
			name: "no remote error",
			json: CorrectLoadLondon,
			want: true,
		},
		{
			name: "some remote error",
			json: ErrorMessageLoad,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateResponseCodeSuccessful(tt.json); got != tt.want {
				t.Errorf("validateResponseCodeSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}
