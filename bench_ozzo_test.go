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

// Benchmark_Ozzo_Validate_Complex validates nested order struct
func Benchmark_Ozzo_Validate_Complex(b *testing.B) {
	order := ValidOrderOzzo
	validateOrder := func(o OrderOzzo) error {
		// Validate Address
		if err := validation.ValidateStruct(&o.Customer.Address,
			validation.Field(&o.Customer.Address.Street, validation.Required, validation.Length(5, 0)),
			validation.Field(&o.Customer.Address.City, validation.Required),
			validation.Field(&o.Customer.Address.Country, validation.Required, validation.Length(2, 2)),
			validation.Field(&o.Customer.Address.Zip, validation.Required),
		); err != nil {
			return err
		}
		// Validate Customer
		if err := validation.ValidateStruct(&o.Customer,
			validation.Field(&o.Customer.ID, validation.Required, is.UUID),
			validation.Field(&o.Customer.Name, validation.Required, validation.Length(2, 0)),
			validation.Field(&o.Customer.Email, validation.Required, is.Email),
		); err != nil {
			return err
		}
		// Validate Items
		for i := range o.Items {
			if err := validation.ValidateStruct(&o.Items[i],
				validation.Field(&o.Items[i].SKU, validation.Required, validation.Length(3, 0)),
				validation.Field(&o.Items[i].Name, validation.Required),
				validation.Field(&o.Items[i].Quantity, validation.Required, validation.Min(1)),
				validation.Field(&o.Items[i].Price, validation.Required, validation.Min(0.01)),
			); err != nil {
				return err
			}
		}
		// Validate Order
		return validation.ValidateStruct(&o,
			validation.Field(&o.ID, validation.Required, is.UUID),
			validation.Field(&o.Total, validation.Required, validation.Min(0.01)),
			validation.Field(&o.Notes, validation.Length(0, 500)),
		)
	}
	_ = validateOrder(order) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validateOrder(order)
	}
}

// Benchmark_Ozzo_Validate_Large validates 20-field config struct
func Benchmark_Ozzo_Validate_Large(b *testing.B) {
	config := ValidConfigOzzo
	validateConfig := func(c ConfigOzzo) error {
		return validation.ValidateStruct(&c,
			validation.Field(&c.AppName, validation.Required, validation.Length(1, 100)),
			validation.Field(&c.Version, validation.Required),
			validation.Field(&c.Environment, validation.Required, validation.In("dev", "staging", "prod")),
			validation.Field(&c.LogLevel, validation.In("debug", "info", "warn", "error")),
			validation.Field(&c.Port, validation.Required, validation.Min(1), validation.Max(65535)),
			validation.Field(&c.Host, validation.Required),
			validation.Field(&c.DatabaseURL, validation.Required, is.URL),
			validation.Field(&c.RedisURL, is.URL),
			validation.Field(&c.MaxConnections, validation.Min(1), validation.Max(1000)),
			validation.Field(&c.Timeout, validation.Min(1), validation.Max(300)),
			validation.Field(&c.RetryCount, validation.Min(0), validation.Max(10)),
			validation.Field(&c.CacheTTL, validation.Min(0)),
			validation.Field(&c.RateLimit, validation.Min(0)),
			validation.Field(&c.APIKey, validation.Required, validation.Length(32, 0)),
			validation.Field(&c.SecretKey, validation.Required, validation.Length(32, 0)),
			validation.Field(&c.MetricsPort, validation.Min(1), validation.Max(65535)),
		)
	}
	_ = validateConfig(config) // warm
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = validateConfig(config)
	}
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
