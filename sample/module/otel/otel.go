package otel

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sin392/db-media-sample/sample/internal/config"
	// "go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"

	// "go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	// "go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	// sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

const name = "github.com/sin392/db-media-sample"

// TODO: パッケージの位置微妙
var tracer = otel.Tracer(name)
var meter = otel.Meter(name)

// otelのbridgeにバグがありコンパイルされない
// ref: https://github.com/open-telemetry/opentelemetry-go-contrib/issues/6239
// var logger = zap.New(otelzap.NewCore(name))

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

func SetupOTelSDK(cfg *config.Config) (shutdown func(context.Context) error, err error) {
	ctx := context.Background()

	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Resourceを設定
	resource, err := NewResource(cfg.AppName, cfg.AppVersion)
	if err != nil {
		handleErr(err)
		return
	}

	// LoggerProviderを設定
	// loggerProvider, err := newLoggerProvider()
	// if err != nil {
	// 	handleErr(err)
	// 	return
	// }
	// shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	// global.SetLoggerProvider(loggerProvider)

	// MeterProviderを設定
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// TracerProviderを設定
	traceProvider, err := newTraceProvider(ctx, resource)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)
	otel.SetTracerProvider(traceProvider)

	// propagatorを設定
	propagator := newPropagator()
	otel.SetTextMapPropagator(propagator)

	return shutdown, nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func NewResource(serviceName string, serviceVersion string) (*sdkresource.Resource, error) {
	return sdkresource.Merge(
		sdkresource.Default(),
		sdkresource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		),
	)
}

// func newLoggerProvider() (*sdklog.LoggerProvider, error) {
// 	exporter, err := stdoutlog.New()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create logger exporter: %w", err)
// 	}

// 	provider := sdklog.NewLoggerProvider(
// 		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
// 	)
// 	return provider, nil
// }

func newMeterProvider() (*sdkmetric.MeterProvider, error) {
	// prometheus形式への変換はotel collectorで行うのが良さそう
	// metricExporter, err := stdoutmetric.New()
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	provider := sdkmetric.NewMeterProvider(
		// sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter,
		// 	// Default is 1m. Set to 3s for demonstrative purposes.
		// 	sdkmetric.WithInterval(3*time.Second))),
		sdkmetric.WithReader(exporter),
	)
	return provider, nil
}

func newTraceProvider(ctx context.Context, resource *sdkresource.Resource) (*sdktrace.TracerProvider, error) {
	opts := []otlptracehttp.Option{
		// WithEndpointURLを使う場合はスキーマの設定も必要
		otlptracehttp.WithEndpoint(os.Getenv("HTTP_TRACING_ENDPOINT")),
	}
	if os.Getenv("APP_ENV") == "local" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	exporter, err := otlptracehttp.New(
		ctx,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource),
		sdktrace.WithSampler(&customSampler{
			delegate: sdktrace.TraceIDRatioBased(1.0),
		}),
	)

	return provider, nil
}
