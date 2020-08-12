package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	"go.opentelemetry.io/otel/sdk/metric/controller/pull"
	"go.opentelemetry.io/otel/sdk/resource"
)

func main() {

}

func sendDataToPrometheus() {
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

	// Create two instruments with Go SDK metric package
	counter := metric.Must(meter).NewInt64Counter(
		"a.counter",
		metric.WithDescription("Counts things"),
	)
	recorder := metric.Must(meter).NewInt64ValueRecorder(
		"a.valuerecorder",
		metric.WithDescription("Records values"),
	)

	// Add initial values to the instruments
	counter.Add(ctx, 5, kv.String("key", "value"))
	recorder.Record(ctx, 100, kv.String("key", "value"))

	// Repeatedly record values every 3 seconds
	go func() {
		for i := 1; i <= 5000; i++ {
			time.Sleep(3 * time.Second)
			value := int64(i * 100)
			fmt.Printf("%d. Recording %d in recorder and adding %d to counter\n", i, value, i)
			recorder.Record(ctx, value, kv.String("key", "value"))
			counter.Add(ctx, int64(i), kv.String("key", "value"))
		}
	}()

	// Set up an endpoint to wait for Prometheus scrapes
	fmt.Println("Server started!")
	http.Handle("/", exporter)
	http.ListenAndServe(":8888", nil)
}
