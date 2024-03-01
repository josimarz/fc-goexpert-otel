package main

import (
	"context"
	"log"

	"github.com/josimarz/fc-goexpert-otel/internal/otel/setup"
	"github.com/josimarz/fc-goexpert-otel/internal/otel/tracing"
	"github.com/josimarz/fc-goexpert-otel/internal/servicea/handler"
)

func main() {
	ctx := context.Background()
	tracer, shutdow := tracing.Start()
	defer func() {
		_ = shutdow(ctx)
	}()
	h := &handler.DefaultHandler{
		Tracer: tracer,
	}
	if err := setup.Run(h); err != nil {
		log.Fatalln(err)
	}
}
