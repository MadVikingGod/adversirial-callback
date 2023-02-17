package bad

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
)

func init() {
	callback := func(ctx context.Context, io instrument.Int64Observer) error {
		return errors.New("bad")
	}
	global.MeterProvider().Meter("bad").Int64ObservableGauge("bad", instrument.WithInt64Callback(callback))
}
