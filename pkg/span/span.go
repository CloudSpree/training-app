package span

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func WithEnvironment(ctx context.Context, tracer trace.Tracer, environmentName string, name string) (context.Context, trace.Span) {
	spanCtx, span := tracer.Start(context.Background(), name)
	span.SetAttributes(attribute.String("environment", environmentName))
	return spanCtx, span
}
