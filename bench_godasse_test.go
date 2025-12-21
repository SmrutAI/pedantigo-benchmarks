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
// Validate (Not applicable - godasse requires hand-written Validate() methods)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_Validate_Simple - godasse doesn't have tag-based validation
func Benchmark_Godasse_Validate_Simple(b *testing.B) {
	b.Skip("godasse requires hand-written Validate() methods, not comparable to tag-based validation")
}

// Benchmark_Godasse_Validate_Complex - godasse doesn't have tag-based validation
func Benchmark_Godasse_Validate_Complex(b *testing.B) {
	b.Skip("godasse requires hand-written Validate() methods, not comparable to tag-based validation")
}

// Benchmark_Godasse_Validate_Large - godasse doesn't have tag-based validation
func Benchmark_Godasse_Validate_Large(b *testing.B) {
	b.Skip("godasse requires hand-written Validate() methods, not comparable to tag-based validation")
}

// ----------------------------------------------------------------------------
// JSONValidate (JSON -> map -> struct + validate)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_JSONValidate_Simple - JSON -> map -> struct + validate
// NOTE: JSON parsing is included in the timer for fair comparison with Pedantigo,
// which also parses JSON inside its Unmarshal function.
func Benchmark_Godasse_JSONValidate_Simple(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[UserGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	jsonData := ValidUserJSON

	// warm
	dict := make(jsonPkg.JSON)
	_ = json.Unmarshal(jsonData, &dict)
	_, _ = deserializer.DeserializeDict(dict)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Include JSON parsing for fair comparison - Pedantigo.Unmarshal also parses JSON
		dict := make(jsonPkg.JSON)
		_ = json.Unmarshal(jsonData, &dict)
		_, _ = deserializer.DeserializeDict(dict)
	}
}

// Benchmark_Godasse_JSONValidate_Complex - JSON -> map -> struct + validate
// NOTE: JSON parsing is included in the timer for fair comparison with Pedantigo,
// which also parses JSON inside its Unmarshal function.
func Benchmark_Godasse_JSONValidate_Complex(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[OrderGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	jsonData := ValidOrderJSON

	// warm
	dict := make(jsonPkg.JSON)
	_ = json.Unmarshal(jsonData, &dict)
	_, _ = deserializer.DeserializeDict(dict)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Include JSON parsing for fair comparison - Pedantigo.Unmarshal also parses JSON
		dict := make(jsonPkg.JSON)
		_ = json.Unmarshal(jsonData, &dict)
		_, _ = deserializer.DeserializeDict(dict)
	}
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
