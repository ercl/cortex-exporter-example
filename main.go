package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	"go.opentelemetry.io/otel/sdk/metric/controller/pull"
	"go.opentelemetry.io/otel/sdk/resource"
)

func main() {
	// Setup a new Prometheus Exporter
	exporter, err := prometheus.NewExportPipeline(
		prometheus.Config{},
		pull.WithResource(resource.New(kv.String("R", "V"))),
	)
	if err != nil {
		panic(err)
	}
	meter := exporter.Provider().Meter("example")
	ctx := context.Background()

	// Read config file.
	config, err := readConfig("data_config.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, config := range config.InstrumentConfigs {
		if config.Type == "COUNTER" {
			counter := metric.Must(meter).NewInt64Counter(
				config.Label,
				metric.WithDescription(config.Description),
			)
			config.instrument = counter

			go func(counter metric.Int64Counter) {
				for i := 1; i <= config.DataPointCount; i++ {
					time.Sleep(time.Duration(config.RecordInterval) * time.Second)
					counter.Add(ctx, 1, kv.String("key", "value"))
					fmt.Printf("%d. Adding 1 to counter\n", i)
				}
			}(counter)
		}
		if config.Type == "VALUERECORDER" {
			recorder := metric.Must(meter).NewInt64ValueRecorder(
				config.Label,
				metric.WithDescription(config.Description),
			)
			config.instrument = recorder

			go func(recorder metric.Int64ValueRecorder) {
				for i := 1; i <= config.DataPointCount; i++ {
					time.Sleep(time.Duration(config.RecordInterval) * time.Second)
					recorder.Record(ctx, 1, kv.String("key", "value"))
					fmt.Printf("%d. Recording %d in recorder\n", i, 1)
				}
			}(recorder)
		}
	}

	// Set up an endpoint to wait for Prometheus scrapes
	fmt.Println("Server started!")
	http.Handle("/", exporter)
	http.ListenAndServe(":8888", nil)
}
