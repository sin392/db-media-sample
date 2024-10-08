package trace

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sin392/db-media-sample/sample/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
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

func InitTraceProvider(cfg *config.Config) {
	opts := []otlptracehttp.Option{
		// WithEndpointURLを使う場合はスキーマの設定も必要
		otlptracehttp.WithEndpoint(os.Getenv("HTTP_TRACING_ENDPOINT")),
	}
	if os.Getenv("APP_ENV") == "local" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	ctx := context.Background()
	exporter, err := otlptracehttp.New(
		ctx,
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
				semconv.ServiceNameKey.String(cfg.AppName),
				semconv.ServiceVersionKey.String(cfg.AppVersion),
			),
		),
		sdktrace.WithSampler(&customSampler{
			delegate: sdktrace.TraceIDRatioBased(1.0),
		}),
	)
	otel.SetTracerProvider(tp)

	// トレーシングコンテキストの伝搬を設定
	otel.SetTextMapPropagator(propagation.TraceContext{})
}
