package benchmarks

import (
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
// Unmarshal (single-call JSON decode + validate, via DeserializeDict)
// ----------------------------------------------------------------------------

// Benchmark_Godasse_Unmarshal_Simple - JSON bytes -> validated struct in one
// call, via godasse's own DeserializeBytes (which decodes internally and then
// calls DeserializeDict, invoking each result's Validate() method - see
// vendor/github.com/pasqal-io/godasse/deserialize/deserialize.go:371-382).
// This has the same (bytes in, *T out) shape as Pedantigo.Unmarshal and
// godantic.Unmarshal, unlike the two-step decode-then-validate pattern in
// Benchmark_Playground_JSONValidate_Simple.
func Benchmark_Godasse_Unmarshal_Simple(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[UserGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	jsonData := ValidUserJSON

	if _, err := deserializer.DeserializeBytes(jsonData); err != nil { // warm
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := deserializer.DeserializeBytes(jsonData); err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark_Godasse_Unmarshal_Complex - JSON bytes -> validated struct in one
// call, for a nested struct. See Benchmark_Godasse_Unmarshal_Simple for why
// this calls DeserializeBytes directly rather than replicating its internals.
func Benchmark_Godasse_Unmarshal_Complex(b *testing.B) {
	deserializer, err := deserialize.MakeMapDeserializer[OrderGodasse](deserialize.Options{
		Unmarshaler: jsonPkg.Driver,
		MainTagName: "json",
	})
	if err != nil {
		b.Fatal(err)
	}

	jsonData := ValidOrderJSON

	if _, err := deserializer.DeserializeBytes(jsonData); err != nil { // warm
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := deserializer.DeserializeBytes(jsonData); err != nil {
			b.Fatal(err)
		}
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
