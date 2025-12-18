package benchmarks

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ============================================================================
// Ozzo Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Pedantigo/Playground)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_Validate_Simple validates an existing 5-field struct
func Benchmark_Ozzo_Validate_Simple(b *testing.B) {
	user := ValidUserOzzo
	// Ozzo uses method-based validation (no struct tags)
	validateUser := func(u UserOzzo) error {
		return validation.ValidateStruct(&u,
			validation.Field(&u.Name, validation.Required, validation.Length(2, 100)),
			validation.Field(&u.Email, validation.Required, is.Email),
			validation.Field(&u.Age, validation.Min(0), validation.Max(150)),
			validation.Field(&u.Website, is.URL),
			validation.Field(&u.Username, is.Alphanumeric, validation.Length(3, 20)),
		)
	}
	_ = validateUser(user) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validateUser(user)
	}
}

// Benchmark_Ozzo_Validate_Complex - Ozzo requires excessive boilerplate for nested structs
func Benchmark_Ozzo_Validate_Complex(b *testing.B) {
	b.Skip("Ozzo complex validation skipped - requires excessive boilerplate for nested structs")
}

// Benchmark_Ozzo_Validate_Large - Ozzo requires excessive boilerplate for large structs
func Benchmark_Ozzo_Validate_Large(b *testing.B) {
	b.Skip("Ozzo large validation skipped - requires excessive boilerplate for 20+ fields")
}

// ----------------------------------------------------------------------------
// Pedantigo-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_UnmarshalMap_Simple - Not applicable to Ozzo
func Benchmark_Ozzo_UnmarshalMap_Simple(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// Benchmark_Ozzo_UnmarshalMap_Complex - Not applicable to Ozzo
func Benchmark_Ozzo_UnmarshalMap_Complex(b *testing.B) {
	b.Skip("UnmarshalMap is a Pedantigo-only feature")
}

// ----------------------------------------------------------------------------
// Playground-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_UnmarshalDirect_Simple - Not applicable to Ozzo
func Benchmark_Ozzo_UnmarshalDirect_Simple(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// Benchmark_Ozzo_UnmarshalDirect_Complex - Not applicable to Ozzo
func Benchmark_Ozzo_UnmarshalDirect_Complex(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// ----------------------------------------------------------------------------
// Validator Creation (Not applicable - Ozzo uses inline validation)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_New_Simple - Ozzo doesn't have a validator object
func Benchmark_Ozzo_New_Simple(b *testing.B) {
	b.Skip("Ozzo uses inline validation, no validator object to create")
}

// Benchmark_Ozzo_New_Complex - Ozzo doesn't have a validator object
func Benchmark_Ozzo_New_Complex(b *testing.B) {
	b.Skip("Ozzo uses inline validation, no validator object to create")
}

// ----------------------------------------------------------------------------
// Schema Generation (Not supported)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_Schema_Uncached - Not supported by Ozzo
func Benchmark_Ozzo_Schema_Uncached(b *testing.B) {
	b.Skip("Ozzo does not support schema generation")
}

// Benchmark_Ozzo_Schema_Cached - Not supported by Ozzo
func Benchmark_Ozzo_Schema_Cached(b *testing.B) {
	b.Skip("Ozzo does not support schema generation")
}

// ----------------------------------------------------------------------------
// Marshal (Not applicable - Ozzo is validation-only)
// ----------------------------------------------------------------------------

// Benchmark_Ozzo_Marshal_Simple - Ozzo doesn't have integrated marshal
func Benchmark_Ozzo_Marshal_Simple(b *testing.B) {
	b.Skip("Ozzo is validation-only, no integrated marshal")
}
