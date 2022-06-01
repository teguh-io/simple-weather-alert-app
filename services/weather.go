package services

import (
	"simple-weather-alert-app/models"
	"simple-weather-alert-app/params"
	"simple-weather-alert-app/repositories"
)

type WeatherService interface {
	GetWeatherMetric() (*params.Weather, error)
	SetWeatherMetric(request params.Weather) error
}

type weatherService struct {
	weatherRepo repositories.WeatherRepository
}

func NewWeatherService(wr repositories.WeatherRepository) WeatherService {
	return &weatherService{
		weatherRepo: wr,
	}
}

func (ws *weatherService) GetWeatherMetric() (*params.Weather, error) {
	res, err := ws.weatherRepo.GetWeatherMetric()
	if err != nil {
		return nil, err
	}

	weather := params.Weather{
		Wind:  res.Status.Wind,
		Water: res.Status.Water,
	}

	return &weather, nil

}

func (ws *weatherService) SetWeatherMetric(request params.Weather) error {
	weatherModel := models.Weather{
		Status: models.Status{
			Water: request.Water,
			Wind:  request.Wind,
		},
	}

	return ws.weatherRepo.SetWeatherMetric(weatherModel)
}
