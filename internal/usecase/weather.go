package usecase

import (
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/infra/gateway"
)

type Weather struct {
	Fahrenheit float64 `json:"temp_F"`
	Celsius    float64 `json:"temp_C"`
	Kelvin     float64 `json:"temp_k"`
}

type WeatherUseCase interface {
	Execute(cep string) (*Weather, error)
}

type weatherUseCaseImpl struct {
	Gateway gateway.WeatherGatewayInterface
}

func NewWeatherUseCase(
	Gateway gateway.WeatherGatewayInterface,
) WeatherUseCase {
	return &weatherUseCaseImpl{
		Gateway: Gateway,
	}
}

func (w *weatherUseCaseImpl) Execute(cep string) (*Weather, error) {
	err := w.Gateway.ValidateLocation(cep)
	if err != nil {
		return nil, err
	}

	weatherC, err := w.Gateway.GetWeather()
	if err != nil {
		return nil, err
	}

	return &Weather{
		Fahrenheit: (weatherC * 1.8) + 32,
		Celsius:    weatherC,
		Kelvin:     weatherC + 273,
	}, nil
}
