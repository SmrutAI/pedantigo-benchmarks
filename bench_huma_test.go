package benchmarks

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/danielgtaylor/huma/v2"
)

// ============================================================================
// Huma Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Validate (Skip - Huma only validates maps, not typed structs)
// ----------------------------------------------------------------------------

// Benchmark_Huma_Validate_Simple - Huma cannot validate typed structs
func Benchmark_Huma_Validate_Simple(b *testing.B) {
	b.Skip("Huma validates map[string]any from JSON, not typed structs - see UnmarshalMap benchmarks")
}

// Benchmark_Huma_Validate_Complex - Huma cannot validate typed structs
func Benchmark_Huma_Validate_Complex(b *testing.B) {
	b.Skip("Huma validates map[string]any from JSON, not typed structs - see UnmarshalMap benchmarks")
}

// Benchmark_Huma_Validate_Large - Huma cannot validate typed structs
func Benchmark_Huma_Validate_Large(b *testing.B) {
	b.Skip("Huma validates map[string]any from JSON, not typed structs - see UnmarshalMap benchmarks")
}

// ----------------------------------------------------------------------------
// JSONValidate (JSON → map → validate) - Huma's actual flow
// ----------------------------------------------------------------------------

// Benchmark_Huma_JSONValidate_Simple tests JSON→map→validate (Huma's real API flow)
// NOTE: Huma only validates the map - it does NOT convert to a typed struct.
// This is less work than Pedantigo which parses JSON, validates, AND outputs a typed struct.
func Benchmark_Huma_JSONValidate_Simple(b *testing.B) {
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	schema := registry.Schema(reflect.TypeOf(UserHuma{}), true, "")
	pb := huma.NewPathBuffer([]byte{}, 0)
	res := &huma.ValidateResult{}

	// warm
	var parsed any
	json.Unmarshal(ValidUserJSON, &parsed)
	huma.Validate(registry, schema, pb, huma.ModeWriteToServer, parsed, res)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p any
		json.Unmarshal(ValidUserJSON, &p)
		res.Reset()
		pb.Reset()
		huma.Validate(registry, schema, pb, huma.ModeWriteToServer, p, res)
	}
}

// Benchmark_Huma_JSONValidate_Complex tests JSON→map→validate for nested structs
// NOTE: Huma only validates the map - it does NOT convert to a typed struct.
// This is less work than Pedantigo which parses JSON, validates, AND outputs a typed struct.
func Benchmark_Huma_JSONValidate_Complex(b *testing.B) {
	registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
	schema := registry.Schema(reflect.TypeOf(OrderHuma{}), true, "")
	pb := huma.NewPathBuffer([]byte{}, 0)
	res := &huma.ValidateResult{}

	// warm
	var parsed any
	json.Unmarshal(ValidOrderJSON, &parsed)
	huma.Validate(registry, schema, pb, huma.ModeWriteToServer, parsed, res)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var p any
		json.Unmarshal(ValidOrderJSON, &p)
		res.Reset()
		pb.Reset()
		huma.Validate(registry, schema, pb, huma.ModeWriteToServer, p, res)
	}
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
// OpenAPI Schema Generation (Huma uses same Schema() for OpenAPI)
// ----------------------------------------------------------------------------

// Benchmark_Huma_OpenAPI_Uncached - OpenAPI schema generation (same as Schema)
func Benchmark_Huma_OpenAPI_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		registry := huma.NewMapRegistry("#/components/schemas/", huma.DefaultSchemaNamer)
		_ = registry.Schema(reflect.TypeOf(UserHuma{}), true, "")
	}
}

// Benchmark_Huma_OpenAPI_Cached - OpenAPI schema reuse from registry
func Benchmark_Huma_OpenAPI_Cached(b *testing.B) {
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