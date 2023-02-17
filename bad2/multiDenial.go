package bad2

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

func init() {
	inst, _ := global.MeterProvider().Meter("bad").Int64ObservableGauge("bad")
	callback := func(ctx context.Context, o metric.Observer) error {
		return errors.New("bad")
	}
	global.MeterProvider().Meter("bad").RegisterCallback(callback, inst)
}
