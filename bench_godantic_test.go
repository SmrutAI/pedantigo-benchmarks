package benchmarks

import (
	"testing"

	"github.com/deepankarm/godantic/pkg/godantic"
)

// ============================================================================
// godantic Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Pedantigo/Playground)
// ----------------------------------------------------------------------------

// Benchmark_Godantic_Validate_Simple validates an existing 5-field struct
func Benchmark_Godantic_Validate_Simple(b *testing.B) {
	user := ValidUserGodantic
	validator := godantic.NewValidator[UserGodantic]()

	// warm
	_ = validator.Validate(&user)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&user)
	}
}

// Benchmark_Godantic_Validate_Complex validates nested order struct
func Benchmark_Godantic_Validate_Complex(b *testing.B) {
	order := ValidOrderGodantic
	validator := godantic.NewValidator[OrderGodantic]()

	// warm
	_ = validator.Validate(&order)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&order)
	}
}

// Benchmark_Godantic_Validate_Large validates 20-field config struct
func Benchmark_Godantic_Validate_Large(b *testing.B) {
	config := ValidConfigGodantic
	validator := godantic.NewValidator[ConfigGodantic]()

	// warm
	_ = validator.Validate(&config)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&config)
	}
}

// ----------------------------------------------------------------------------
// Pedantigo-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Godantic_UnmarshalMap_Simple - Not applicable to godantic
func Benchmark_Godantic_UnmarshalMap_Simple(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// Benchmark_Godantic_UnmarshalMap_Complex - Not applicable to godantic
func Benchmark_Godantic_UnmarshalMap_Complex(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// ----------------------------------------------------------------------------
// Playground-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Godantic_UnmarshalDirect_Simple - Not applicable to godantic
func Benchmark_Godantic_UnmarshalDirect_Simple(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// Benchmark_Godantic_UnmarshalDirect_Complex - Not applicable to godantic
func Benchmark_Godantic_UnmarshalDirect_Complex(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Godantic_New_Simple - Validator creation overhead
func Benchmark_Godantic_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = godantic.NewValidator[UserGodantic]()
	}
}

// Benchmark_Godantic_New_Complex - Validator creation overhead
func Benchmark_Godantic_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = godantic.NewValidator[OrderGodantic]()
	}
}

// ----------------------------------------------------------------------------
// Schema Generation (Not supported as standalone)
// ----------------------------------------------------------------------------

// Benchmark_Godantic_Schema_Uncached - Not supported by godantic
func Benchmark_Godantic_Schema_Uncached(b *testing.B) {
	b.Skip("godantic does not support standalone schema generation")
}

// Benchmark_Godantic_Schema_Cached - Not supported by godantic
func Benchmark_Godantic_Schema_Cached(b *testing.B) {
	b.Skip("godantic does not support standalone schema generation")
}

// ----------------------------------------------------------------------------
// Marshal (Not applicable - godantic is validation-only)
// ----------------------------------------------------------------------------

// Benchmark_Godantic_Marshal_Simple - godantic doesn't have integrated marshal
func Benchmark_Godantic_Marshal_Simple(b *testing.B) {
	b.Skip("godantic is validation-only, no integrated marshal")
}
