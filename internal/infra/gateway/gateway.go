package gateway

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var (
	ErrorInvalidCEP  = errors.New("invalid zipcode")
	ErrorNotFoundCEP = errors.New("can not find zipcode")
)

var getFunc = http.Get

type WeatherGatewayInterface interface {
	ValidateLocation(cep string) error
	GetWeather() (float64, error)
}

type cepResponse struct {
	Error    string `json:"erro"`
	Location string `json:"localidade"`
}

type weatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type weatherGatewayImpl struct {
	apiKey   string
	location string
}

func New(apiKey string) WeatherGatewayInterface {
	return &weatherGatewayImpl{
		apiKey: apiKey,
	}
}

func (w *weatherGatewayImpl) ValidateLocation(cep string) error {
	if len(cep) != 8 {
		return ErrorInvalidCEP
	}

	resp, err := getFunc("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusBadRequest {
		return ErrorNotFoundCEP
	}

	var content cepResponse
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return err
	}

	if content.Error != "" {
		return ErrorNotFoundCEP
	}
	w.location = content.Location
	return nil
}

func (w *weatherGatewayImpl) GetWeather() (float64, error) {
	u, err := url.Parse("http://api.weatherapi.com/v1/current.json")
	if err != nil {
		return 0, err
	}

	query := url.Values{}
	query.Set("key", w.apiKey)
	query.Set("q", w.location)
	u.RawQuery = query.Encode()

	resp, err := getFunc(u.String())
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("invalid status received")
	}

	var content weatherResponse
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return 0, err
	}
	return content.Current.TempC, nil
}
