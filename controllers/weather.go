package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"simple-weather-alert-app/params"
	"simple-weather-alert-app/services"
	"text/template"
)

type WeatherController interface {
	GetWeatherMetric(w http.ResponseWriter, r *http.Request)
}

type weatherController struct {
	weatherService services.WeatherService
}

func NewWeatherController(ws services.WeatherService) WeatherController {
	return &weatherController{
		weatherService: ws,
	}
}

func (wc *weatherController) GetWeatherMetric(w http.ResponseWriter, r *http.Request) {
	var errorResponse map[string]interface{}
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("template/static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		metric := params.Weather{
			Water: uint(rand.Intn(100)),
			Wind:  uint(rand.Intn(100)),
		}

		err = wc.weatherService.SetWeatherMetric(metric)
		if err != nil {
			errorResponse["error"] = err.Error()
			jsonResp, err := json.Marshal(errorResponse)
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}

		data, err := wc.weatherService.GetWeatherMetric()
		if err != nil {
			errorResponse["error"] = err.Error()
			jsonResp, err := json.Marshal(errorResponse)
			if err != nil {
				panic(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
