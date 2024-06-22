package provider

import (
	"context"

	of "github.com/open-feature/go-sdk/openfeature"
)

type P struct {
	md of.Metadata
	h  []of.Hook
}

// P implements the of.FeatureProvider interface
var _ of.FeatureProvider = &P{}

// Metadata returns the provider metadata struct
func (p *P) Metadata() of.Metadata {
	return p.md
}

// Hooks returns a copy of the provider hook slice
func (p *P) Hooks() []of.Hook {
	// exit early on empty hooks
	if len(p.h) == 0 {
		return []of.Hook{}
	}

	// otherwise make a copy and return it.
	out := make([]of.Hook, len(p.h))
	copy(out, p.h)
	return out
}

// BooleanEvaluation performs a flag eval in a boolean context
func (p *P) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx of.FlattenedContext) of.BoolResolutionDetail {
	return of.BoolResolutionDetail{}
}

// StringEvaluation performs a flag eval in a string context
func (p *P) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx of.FlattenedContext) of.StringResolutionDetail {
	return of.StringResolutionDetail{}
}

// FloatEvaluation performs a flag eval in a float context
func (p *P) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx of.FlattenedContext) of.FloatResolutionDetail {
	return of.FloatResolutionDetail{}
}

// IntEvaluation performs a flag eval in an integer context
func (p *P) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx of.FlattenedContext) of.IntResolutionDetail {
	return of.IntResolutionDetail{}
}

// ObjectEvaluation  performs a flag eval in a struct context
func (p *P) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx of.FlattenedContext) of.InterfaceResolutionDetail {
	return of.InterfaceResolutionDetail{}
}
