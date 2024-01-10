package container

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/metadata"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

type (
	// contextKey is an unexported type used as key for items stored in the
	// Context object
	contextKey struct{}

	// propagator implements the custom context propagator
	propagator struct{}
)

// PropagateKey is the key used to store the value in the Context object
var PropagateKey = contextKey{}

// propagationKey is the key used by the propagator to pass values through the
// Temporal server headers
const propagationKey = "custom-header"

// newContextPropagator returns a context propagator that propagates a set of
// string key-value pairs across a workflow
func newContextPropagator() workflow.ContextPropagator {
	return &propagator{}
}

// Inject injects values from context into headers for propagation
func (s *propagator) Inject(ctx context.Context, writer workflow.HeaderWriter) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return fmt.Errorf("unable to extract metadata from context")
	}

	payload, err := converter.GetDefaultDataConverter().ToPayload(md)
	if err != nil {
		return err
	}

	writer.Set(propagationKey, payload)
	return nil
}

// InjectFromWorkflow injects values from context into headers for propagation
func (s *propagator) InjectFromWorkflow(ctx workflow.Context, writer workflow.HeaderWriter) error {
	value := ctx.Value(PropagateKey)
	payload, err := converter.GetDefaultDataConverter().ToPayload(value)
	if err != nil {
		return err
	}
	writer.Set(propagationKey, payload)
	return nil
}

// Extract extracts values from headers and puts them into context
func (s *propagator) Extract(ctx context.Context, reader workflow.HeaderReader) (context.Context, error) {
	if value, ok := reader.Get(propagationKey); ok {
		var values metadata.Metadata
		if err := converter.GetDefaultDataConverter().FromPayload(value, &values); err != nil {
			return ctx, nil
		}
		ctx = appendMetadata(ctx, values)
	}

	return ctx, nil
}

func appendMetadata(ctx context.Context, md metadata.Metadata) context.Context {
	old, ok := metadata.FromContext(ctx)
	if !ok {
		if md == nil {
			md = metadata.Metadata{}
		}

		return metadata.NewContext(ctx, md)
	}

	for k, v := range md {
		old[k] = v
	}

	return ctx
}

// ExtractToWorkflow extracts values from headers and puts them into context
func (s *propagator) ExtractToWorkflow(ctx workflow.Context, reader workflow.HeaderReader) (workflow.Context, error) {
	if value, ok := reader.Get(propagationKey); ok {
		var values metadata.Metadata
		if err := converter.GetDefaultDataConverter().FromPayload(value, &values); err != nil {
			return ctx, nil
		}
		ctx = workflow.WithValue(ctx, PropagateKey, values)
	}
	return ctx, nil
}
