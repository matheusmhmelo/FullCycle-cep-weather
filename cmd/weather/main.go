package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/infra/gateway"
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/infra/web"
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/infra/web/webserver"
	"github.com/matheusmhmelo/FullCycle-cep-weather/internal/usecase"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load("cmd/weather/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	gt := gateway.New(os.Getenv("API_KEY"))
	weatherUseCase := usecase.NewWeatherUseCase(gt)

	webServer := webserver.NewWebServer(os.Getenv("WEB_SERVER_PORT"))
	webOrderHandler := web.NewOrderHandler(weatherUseCase)
	webServer.AddHandler("/weather", webserver.HTTP_GET, webOrderHandler.Get)
	fmt.Println("Starting web server on port", os.Getenv("WEB_SERVER_PORT"))
	webServer.Start()
}
