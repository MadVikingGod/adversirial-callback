package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/sdk/metric"
	// _ "main.local/bad"
	// _ "main.local/bad2"
)

func main() {
	rdr := metric.NewManualReader()
	provider := metric.NewMeterProvider(metric.WithReader(rdr))
	global.SetMeterProvider(provider)

	ctr, _ := provider.Meter("test").Int64Counter("hello")
	ctr.Add(context.Background(), 1)

	provider.Meter("test").Int64ObservableCounter("helloObs", instrument.WithInt64Callback(func(ctx context.Context, io instrument.Int64Observer) error {
		io.Observe(1)
		return nil
	}))

	rm, _ := rdr.Collect(context.Background())

	fmt.Println(rm)
}
