package main

import (
	"fmt"
	"net/http"
	"simple-weather-alert-app/controllers"
	"simple-weather-alert-app/repositories"
	"simple-weather-alert-app/services"
)

const PORT = ":8080"

func main() {
	weatherRepo := repositories.NewWeatherRepository()
	weatherService := services.NewWeatherService(weatherRepo)
	weatherController := controllers.NewWeatherController(weatherService)

	http.HandleFunc("/", weatherController.GetWeatherMetric)
	http.Handle("/template/static/", http.StripPrefix("/template/static/", http.FileServer(http.Dir("template/static"))))
	fmt.Println("Server running on port:", PORT)
	http.ListenAndServe(PORT, nil)
}
