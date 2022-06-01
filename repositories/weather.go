package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"simple-weather-alert-app/models"
)

type WeatherRepository interface {
	SetWeatherMetric(metric models.Weather) error
	GetWeatherMetric() (*models.Weather, error)
}

type weatherRepository struct {
	weather models.Weather
}

func NewWeatherRepository() WeatherRepository {
	return &weatherRepository{}
}

func (wr *weatherRepository) SetWeatherMetric(metric models.Weather) error {
	json, err := json.Marshal(&metric)
	if err != nil {
		return err
	}

	fmt.Println(string(json))

	err = ioutil.WriteFile("template/static/data.json", json, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

func (wr *weatherRepository) GetWeatherMetric() (*models.Weather, error) {
	file, err := ioutil.ReadFile("template/static/data.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &wr.weather)
	if err != nil {
		return nil, err
	}

	return &wr.weather, nil
}
