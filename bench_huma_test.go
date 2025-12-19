package benchmarks

import (
	"reflect"
	"testing"

	"github.com/danielgtaylor/huma/v2"
)

// ============================================================================
// Huma Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Pedantigo/Playground)
// ----------------------------------------------------------------------------

// Benchmark_Huma_Validate_Simple validates an existing 5-field struct
func Benchmark_Huma_Validate_Simple(b *testing.B) {
	user := ValidUserHuma
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	schema := registry.Schema(reflect.TypeOf(user), true, "")
	pb := huma.NewPathBuffer([]byte{}, 0)
	res := &huma.ValidateResult{}

	// warm
	huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &user, res)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res.Reset()
		pb.Reset()
		huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &user, res)
	}
}

// Benchmark_Huma_Validate_Complex validates nested order struct
func Benchmark_Huma_Validate_Complex(b *testing.B) {
	order := ValidOrderHuma
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	schema := registry.Schema(reflect.TypeOf(order), true, "")
	pb := huma.NewPathBuffer([]byte{}, 0)
	res := &huma.ValidateResult{}

	// warm
	huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &order, res)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res.Reset()
		pb.Reset()
		huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &order, res)
	}
}

// Benchmark_Huma_Validate_Large validates 20-field config struct
func Benchmark_Huma_Validate_Large(b *testing.B) {
	config := ValidConfigHuma
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	schema := registry.Schema(reflect.TypeOf(config), true, "")
	pb := huma.NewPathBuffer([]byte{}, 0)
	res := &huma.ValidateResult{}

	// warm
	huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &config, res)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res.Reset()
		pb.Reset()
		huma.Validate(registry, schema, pb, huma.ModeWriteToServer, &config, res)
	}
}

// ----------------------------------------------------------------------------
// Pedantigo-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Huma_UnmarshalMap_Simple - Not applicable to Huma
func Benchmark_Huma_UnmarshalMap_Simple(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// Benchmark_Huma_UnmarshalMap_Complex - Not applicable to Huma
func Benchmark_Huma_UnmarshalMap_Complex(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// ----------------------------------------------------------------------------
// Playground-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Huma_UnmarshalDirect_Simple - Not applicable to Huma
func Benchmark_Huma_UnmarshalDirect_Simple(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// Benchmark_Huma_UnmarshalDirect_Complex - Not applicable to Huma
func Benchmark_Huma_UnmarshalDirect_Complex(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Huma_New_Simple - Registry + Schema creation overhead
func Benchmark_Huma_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
		_ = registry.Schema(reflect.TypeOf(UserHuma{}), true, "")
	}
}

// Benchmark_Huma_New_Complex - Registry + Schema creation overhead
func Benchmark_Huma_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
		_ = registry.Schema(reflect.TypeOf(OrderHuma{}), true, "")
	}
}

// ----------------------------------------------------------------------------
// Schema Generation
// ----------------------------------------------------------------------------

// Benchmark_Huma_Schema_Uncached - Schema generation without caching
func Benchmark_Huma_Schema_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
		_ = registry.Schema(reflect.TypeOf(UserHuma{}), true, "")
	}
}

// Benchmark_Huma_Schema_Cached - Schema reuse from registry
func Benchmark_Huma_Schema_Cached(b *testing.B) {
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	_ = registry.Schema(reflect.TypeOf(UserHuma{}), true, "") // prime cache

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = registry.Schema(reflect.TypeOf(UserHuma{}), true, "")
	}
}

// ----------------------------------------------------------------------------
// Marshal (Not applicable - Huma is API framework, not serialization)
// ----------------------------------------------------------------------------

// Benchmark_Huma_Marshal_Simple - Huma doesn't have standalone marshal
func Benchmark_Huma_Marshal_Simple(b *testing.B) {
	b.Skip("Huma is an API framework, not a serialization library")
}
