package web

import (
	"encoding/json"
	"errors"
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/infra/gateway"
	"net/http"

	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/usecase"
)

type OrderHandler struct {
	weather usecase.WeatherUseCase
}

func NewOrderHandler(
	weather usecase.WeatherUseCase,
) *OrderHandler {
	return &OrderHandler{
		weather: weather,
	}
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	output, err := h.weather.Execute(cep)
	if err != nil {
		if errors.Is(err, gateway.ErrorInvalidCEP) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, gateway.ErrorNotFoundCEP) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
