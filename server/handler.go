package server

import (
	"github.com/dominik-najberg/gogoapps/client"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type handlers struct {
	fetcher client.Fetcher
	appID string
}

func newHandlers(f client.Fetcher) *handlers {
	appID := os.Getenv("API_KEY")
	if appID == ""{
		log.Fatalf("API_KEY misssing from ENV variables")
	}
	return &handlers{
		fetcher: f,
		appID: appID,
	}
}

// /weather/city/?c={city names comma separated}&appid={your api key}
func (h *handlers) weatherByCitiesViewHandler(writer http.ResponseWriter, r *http.Request) {
	var (
		q         = r.URL.Query()
		cityNames = strings.Split(q.Get("cities"), ",")
		apiClient = client.NewAPIClient(h.appID, h.fetcher)
	)

	var resources []*client.OpenWeatherMapResponse

	for _, c := range cityNames {
		res, err := apiClient.GetWeatherByCityName(c)
		if err != nil {
			log.Printf("error while fetching weather data for %s: %v", c, err)
			continue
		}
		if res.Cod != 200 {
			log.Printf("remote server error for %s: %s", c, res.Message)
			continue
		}
		resources = append(resources, res)
	}


	html, err := template.ParseFiles("server/template/bootstrap.html")
	if err != nil {
		log.Fatalf("error while parsing template: %v", err)
	}
	err = html.Execute(writer, struct{Cities []*client.OpenWeatherMapResponse}{Cities:resources})
	if err != nil {
		log.Fatalf("error while parsing template: %v", err)
	}
}
