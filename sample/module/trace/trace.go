package trace

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
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

type customSampler struct {
	delegate sdktrace.Sampler
}

func (cs *customSampler) ShouldSample(p sdktrace.SamplingParameters) sdktrace.SamplingResult {
	// /metricsはサンプリングしない
	if strings.Contains(p.Name, "/metrics") {
		return sdktrace.SamplingResult{Decision: sdktrace.RecordOnly}
	}
	return cs.delegate.ShouldSample(p)
}

func (cs *customSampler) Description() string {
	return "CustomSampler"
}

func InitTraceProvider() {
	opts := []otlptracehttp.Option{
		// WithEndpointURLを使う場合はスキーマの設定も必要
		otlptracehttp.WithEndpoint(os.Getenv("TRACING_ENDPOINT")),
	}
	if os.Getenv("APP_ENV") == "local" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	exporter, err := otlptracehttp.New(
		context.Background(),
		opts...,
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
		sdktrace.WithSampler(&customSampler{
			delegate: sdktrace.TraceIDRatioBased(1.0),
		}),
	)
	otel.SetTracerProvider(tp)
}
