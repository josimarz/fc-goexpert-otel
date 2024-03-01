package main

import (
	"context"
	"log"
	"os"

	"github.com/josimarz/fc-goexpert-otel/internal/otel/setup"
	"github.com/josimarz/fc-goexpert-otel/internal/otel/tracing"
	"github.com/josimarz/fc-goexpert-otel/internal/serviceb/handler"
	"github.com/josimarz/fc-goexpert-otel/internal/serviceb/viacep"
	"github.com/josimarz/fc-goexpert-otel/internal/serviceb/weather"
)

func main() {
	ctx := context.Background()
	tracer, shutdow := tracing.Start()
	defer func() {
		_ = shutdow(ctx)
	}()
	viaCepApi := &viacep.ViaCepApi{}
	weatherApi := &weather.WeatherApi{
		Key: os.Getenv("API_KEY"),
	}
	h := &handler.DefaultHandler{
		ViaCepApi:  viaCepApi,
		WeatherApi: weatherApi,
		Tracer:     tracer,
	}
	if err := setup.Run(h); err != nil {
		log.Fatalln(err)
	}
}
