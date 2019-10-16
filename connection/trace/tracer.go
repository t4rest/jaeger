package trace

import (
	"fmt"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

// InitJaeger .
func InitJaeger(serviceName, jaegerEndpointAddr string) (*jaeger.Exporter, error) {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: jaegerEndpointAddr,
		Process: jaeger.Process{
			ServiceName: serviceName,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("new exporter: %s", err)
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	return exporter, nil
}
