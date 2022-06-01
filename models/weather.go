package models

type Status struct {
	Water uint `json:"water"`
	Wind  uint `json:"wind"`
}

type Weather struct {
	Status Status `json:"status"`
}
