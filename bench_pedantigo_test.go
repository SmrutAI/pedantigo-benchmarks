package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/SmrutAI/pedantigo"
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
	_ = pedantigo.Validate(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&user)
	}
}

// Benchmark_Pedantigo_Validate_Complex validates an existing nested struct (bypass)
func Benchmark_Pedantigo_Validate_Complex(b *testing.B) {
	order := ValidOrderPedantigo
	_ = pedantigo.Validate(&order) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&order)
	}
}

// Benchmark_Pedantigo_Validate_Large validates an existing 20+ field struct (bypass)
func Benchmark_Pedantigo_Validate_Large(b *testing.B) {
	config := ValidConfigPedantigo
	_ = pedantigo.Validate(&config) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Validate(&config)
	}
}

// ----------------------------------------------------------------------------
// Pedantigo Unique: UnmarshalMap (JSON → map → struct + validate)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_UnmarshalMap_Simple tests Pedantigo's unique JSON→map→struct flow
func Benchmark_Pedantigo_UnmarshalMap_Simple(b *testing.B) {
	_, _ = pedantigo.Unmarshal[UserPedantigo](ValidUserJSON) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Unmarshal[UserPedantigo](ValidUserJSON)
	}
}

// Benchmark_Pedantigo_UnmarshalMap_Complex tests Pedantigo's unique JSON→map→nested flow
func Benchmark_Pedantigo_UnmarshalMap_Complex(b *testing.B) {
	_, _ = pedantigo.Unmarshal[OrderPedantigo](ValidOrderJSON) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Unmarshal[OrderPedantigo](ValidOrderJSON)
	}
}

// ----------------------------------------------------------------------------
// UnmarshalDirect (json.Unmarshal + Validate)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_UnmarshalDirect_Simple tests stdlib json.Unmarshal + Validate
func Benchmark_Pedantigo_UnmarshalDirect_Simple(b *testing.B) {
	var user UserPedantigo
	_ = json.Unmarshal(ValidUserJSON, &user)
	_ = pedantigo.Validate(&user)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var u UserPedantigo
		_ = json.Unmarshal(ValidUserJSON, &u)
		_ = pedantigo.Validate(&u)
	}
}

// Benchmark_Pedantigo_UnmarshalDirect_Complex tests stdlib json.Unmarshal + Validate for nested
func Benchmark_Pedantigo_UnmarshalDirect_Complex(b *testing.B) {
	var order OrderPedantigo
	_ = json.Unmarshal(ValidOrderJSON, &order)
	_ = pedantigo.Validate(&order)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var o OrderPedantigo
		_ = json.Unmarshal(ValidOrderJSON, &o)
		_ = pedantigo.Validate(&o)
	}
}

// ----------------------------------------------------------------------------
// Validator Creation
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_New_Simple measures validator creation for simple struct
func Benchmark_Pedantigo_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.New[UserPedantigo]()
	}
}

// Benchmark_Pedantigo_New_Complex measures validator creation for nested struct
func Benchmark_Pedantigo_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.New[OrderPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Schema Generation (Pedantigo-only feature)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Schema_Uncached measures first-call schema generation
func Benchmark_Pedantigo_Schema_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := pedantigo.New[UserPedantigo]()
		_ = v.Schema()
	}
}

// Benchmark_Pedantigo_Schema_Cached measures cached schema retrieval
func Benchmark_Pedantigo_Schema_Cached(b *testing.B) {
	_ = pedantigo.Schema[UserPedantigo]() // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.Schema[UserPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// OpenAPI Schema Generation
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_OpenAPI_Uncached measures first-call OpenAPI schema generation
func Benchmark_Pedantigo_OpenAPI_Uncached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := pedantigo.New[UserPedantigo]()
		_ = v.SchemaOpenAPI()
	}
}

// Benchmark_Pedantigo_OpenAPI_Cached measures cached OpenAPI schema retrieval
func Benchmark_Pedantigo_OpenAPI_Cached(b *testing.B) {
	_ = pedantigo.SchemaOpenAPI[UserPedantigo]() // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = pedantigo.SchemaOpenAPI[UserPedantigo]()
	}
}

// ----------------------------------------------------------------------------
// Marshal (Validate + JSON output)
// ----------------------------------------------------------------------------

// Benchmark_Pedantigo_Marshal_Simple measures validate + JSON marshal
func Benchmark_Pedantigo_Marshal_Simple(b *testing.B) {
	user := ValidUserPedantigo
	_, _ = pedantigo.Marshal(&user) // warm cache
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pedantigo.Marshal(&user)
	}
}
