package trace

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

// TODO: パッケージの位置微妙
var tracer = otel.Tracer("github.com/sin392/db-media-sample")

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return tracer.Start(ctx, name)
}

func InitTraceProvider() {
	exporter, err := zipkin.New(
		os.Getenv("ZIPKIN_ENDPOINT"),
	)
	if err != nil {
		fmt.Println("failed to create exporter:", err)
		return
	}
	// BatchSpanProcessorを設定
	batcher := sdktrace.NewBatchSpanProcessor(exporter)
	// TracerProviderを設定
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		// Other options
		sdktrace.WithResource(
			sdkresource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("sample"),
				semconv.ServiceVersionKey.String("v1.0.0"),
			),
		),
	)
	otel.SetTracerProvider(tp)
}
