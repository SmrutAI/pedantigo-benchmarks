package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/SmrutAI/pedantigo/v2/validator"
)

// ============================================================================
// Pedantigo Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Playground)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Validate_Simple validates an existing 5-field struct (bypass)
func Benchmark_Pedantigo_Validate_Simple(b *testing.B) {
	user := ValidUserPedantigo
	_ = validator.Validate(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&user)
	}
}

// Benchmark_Pedantigo_Validate_Complex validates an existing nested struct (bypass)
func Benchmark_Pedantigo_Validate_Complex(b *testing.B) {
	order := ValidOrderPedantigo
	_ = validator.Validate(&order) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&order)
	}
}

// Benchmark_Pedantigo_Validate_Large validates an existing 20+ field struct (bypass)
func Benchmark_Pedantigo_Validate_Large(b *testing.B) {
	config := ValidConfigPedantigo
	_ = validator.Validate(&config) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Validate(&config)
	}
}

// ----------------------------------------------------------------------------
// JSONValidate (json.Unmarshal + Validate)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_JSONValidate_Simple tests stdlib json.Unmarshal + Validate
func Benchmark_Pedantigo_JSONValidate_Simple(b *testing.B) {
	var user UserPedantigo
	_ = json.Unmarshal(ValidUserJSON, &user)
	_ = validator.Validate(&user)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var u UserPedantigo
		_ = json.Unmarshal(ValidUserJSON, &u)
		_ = validator.Validate(&u)
	}
}

// Benchmark_Pedantigo_JSONValidate_Complex tests stdlib json.Unmarshal + Validate for nested
func Benchmark_Pedantigo_JSONValidate_Complex(b *testing.B) {
	var order OrderPedantigo
	_ = json.Unmarshal(ValidOrderJSON, &order)
	_ = validator.Validate(&order)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var o OrderPedantigo
		_ = json.Unmarshal(ValidOrderJSON, &o)
		_ = validator.Validate(&o)
	}
}

// ----------------------------------------------------------------------------
// Unmarshal (single-call JSON decode + defaults + validate)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Unmarshal_Simple tests pedantigo's own Unmarshal: JSON
// bytes decoded, defaults applied, and validated in one call (unlike
// Benchmark_Pedantigo_JSONValidate_Simple, which uses stdlib json.Unmarshal
// followed by a separate Validate call).
func Benchmark_Pedantigo_Unmarshal_Simple(b *testing.B) {
	if _, err := validator.Unmarshal[UserPedantigo](ValidUserJSON); err != nil { // warm cache
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := validator.Unmarshal[UserPedantigo](ValidUserJSON); err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark_Pedantigo_Unmarshal_Complex tests pedantigo's own Unmarshal for a
// nested struct: JSON bytes decoded, defaults applied, and validated in one call.
func Benchmark_Pedantigo_Unmarshal_Complex(b *testing.B) {
	if _, err := validator.Unmarshal[OrderPedantigo](ValidOrderJSON); err != nil { // warm cache
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := validator.Unmarshal[OrderPedantigo](ValidOrderJSON); err != nil {
			b.Fatal(err)
		}
	}
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_New_Simple measures validator creation for simple struct
func Benchmark_Pedantigo_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.New[UserPedantigo]()
	}
}

// Benchmark_Pedantigo_New_Complex measures validator creation for nested struct
func Benchmark_Pedantigo_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.New[OrderPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Schema Generation (Pedantigo-only feature)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Schema_Uncached measures first-call schema generation
func Benchmark_Pedantigo_Schema_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := validator.New[UserPedantigo]()
		_ = v.Schema()
	}
}

// Benchmark_Pedantigo_Schema_Cached measures cached schema retrieval
func Benchmark_Pedantigo_Schema_Cached(b *testing.B) {
	_ = validator.Schema[UserPedantigo]() // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.Schema[UserPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// OpenAPI Schema Generation
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_OpenAPI_Uncached measures first-call OpenAPI schema generation
func Benchmark_Pedantigo_OpenAPI_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := validator.New[UserPedantigo]()
		_ = v.SchemaOpenAPI()
	}
}

// Benchmark_Pedantigo_OpenAPI_Cached measures cached OpenAPI schema retrieval
func Benchmark_Pedantigo_OpenAPI_Cached(b *testing.B) {
	_ = validator.SchemaOpenAPI[UserPedantigo]() // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validator.SchemaOpenAPI[UserPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Marshal (Validate + JSON output)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Marshal_Simple measures validate + JSON marshal
func Benchmark_Pedantigo_Marshal_Simple(b *testing.B) {
	user := ValidUserPedantigo
	_, _ = validator.Marshal(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = validator.Marshal(&user)
	}
}
