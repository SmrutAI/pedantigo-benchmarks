package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/pasqal-io/godasse/deserialize"
	jsonPkg "github.com/pasqal-io/godasse/deserialize/json"
)

// ============================================================================
// godasse Benchmarks
// ============================================================================

// ----------------------------------------------------------------------------
// Core Comparison (Apples-to-Apples with Pedantigo/Playground)
// Note: godasse is a deserializer with validation interface, not a standalone
// validator. For fair comparison we benchmark the Validate() method directly.
// ----------------------------------------------------------------------------

// Benchmark_Godasse_Validate_Simple validates an existing 5-field struct
func Benchmark_Godasse_Validate_Simple(b *testing.B) {
	user := ValidUserGodasse

	// warm
	_ = user.Validate()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = user.Validate()
	}
}

// Benchmark_Godasse_Validate_Complex validates nested order struct
func Benchmark_Godasse_Validate_Complex(b *testing.B) {
	order := ValidOrderGodasse

	// warm
	_ = order.Validate()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = order.Validate()
	}
}

// Benchmark_Godasse_Validate_Large validates 20-field config struct
func Benchmark_Godasse_Validate_Large(b *testing.B) {
	config := ValidConfigGodasse

	// warm
	_ = config.Validate()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = config.Validate()
	}
}

// ----------------------------------------------------------------------------
// godasse Unmarshal + Validate (comparable to Pedantigo UnmarshalMap)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_UnmarshalMap_Simple - JSON -> map -> struct + validate
func Benchmark_Godasse_UnmarshalMap_Simple(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[UserGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	// Pre-parse JSON to map
	jsonData := ValidUserJSON
	dict := make(jsonPkg.JSON)
	if err := json.Unmarshal(jsonData, &dict); err != nil {
		b.Fatal(err)
	}

	// warm
	_, _ = deserializer.DeserializeDict(dict)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = deserializer.DeserializeDict(dict)
	}
}

// Benchmark_Godasse_UnmarshalMap_Complex - JSON -> map -> struct + validate
func Benchmark_Godasse_UnmarshalMap_Complex(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[OrderGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	// Pre-parse JSON to map
	jsonData := ValidOrderJSON
	dict := make(jsonPkg.JSON)
	if err := json.Unmarshal(jsonData, &dict); err != nil {
		b.Fatal(err)
	}

	// warm
	_, _ = deserializer.DeserializeDict(dict)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = deserializer.DeserializeDict(dict)
	}
}

// ----------------------------------------------------------------------------
// Playground-only features (Skip)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_UnmarshalDirect_Simple - Not applicable to godasse
func Benchmark_Godasse_UnmarshalDirect_Simple(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// Benchmark_Godasse_UnmarshalDirect_Complex - Not applicable to godasse
func Benchmark_Godasse_UnmarshalDirect_Complex(b *testing.B) {
	b.Skip("UnmarshalDirect is a Playground-only pattern")
}

// ----------------------------------------------------------------------------
// Validator/Deserializer Creation
// ----------------------------------------------------------------------------

// Benchmark_Godasse_New_Simple - Deserializer creation overhead
func Benchmark_Godasse_New_Simple(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = deserialize.MakeMapDeserializer[UserGodasse](deserialize.Options{
			Unmarshaler: jsonPkg.Driver,
			MainTagName: "json",
		})
	}
}

// Benchmark_Godasse_New_Complex - Deserializer creation overhead
func Benchmark_Godasse_New_Complex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = deserialize.MakeMapDeserializer[OrderGodasse](deserialize.Options{
			Unmarshaler: jsonPkg.Driver,
			MainTagName: "json",
		})
	}
}

// ----------------------------------------------------------------------------
// Schema Generation (Not supported)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_Schema_Uncached - Not supported by godasse
func Benchmark_Godasse_Schema_Uncached(b *testing.B) {
	b.Skip("godasse does not support schema generation")
}

// Benchmark_Godasse_Schema_Cached - Not supported by godasse
func Benchmark_Godasse_Schema_Cached(b *testing.B) {
	b.Skip("godasse does not support schema generation")
}

// ----------------------------------------------------------------------------
// Marshal (Not applicable - godasse is deserialization-only)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_Marshal_Simple - godasse doesn't have marshal
func Benchmark_Godasse_Marshal_Simple(b *testing.B) {
	b.Skip("godasse is deserialization-only, no marshal support")
}
