package benchmarks

import (
	"github.com/deepankarm/godantic/pkg/godantic"
)

// ============================================================================
// godantic Types (Method-based validation)
// ============================================================================

// ----------------------------------------------------------------------------
// Simple User (5 fields)
// ----------------------------------------------------------------------------

type UserGodantic struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Website  string `json:"website"`
	Username string `json:"username"`
}

func (u *UserGodantic) FieldName() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(2),
		godantic.MaxLen(100),
	)
}

func (u *UserGodantic) FieldEmail() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.Email(),
	)
}

func (u *UserGodantic) FieldAge() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Min(0),
		godantic.Max(150),
	)
}

func (u *UserGodantic) FieldWebsite() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.URL(),
	)
}

func (u *UserGodantic) FieldUsername() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.MinLen(3),
		godantic.MaxLen(20),
		godantic.Regex(`^[a-zA-Z0-9]+$`),
	)
}

// ----------------------------------------------------------------------------
// Complex Order (nested structs, arrays)
// ----------------------------------------------------------------------------

type AddressGodantic struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

func (a *AddressGodantic) FieldStreet() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(5),
	)
}

func (a *AddressGodantic) FieldCity() godantic.FieldOptions[string] {
	return godantic.Field(godantic.Required[string]())
}

func (a *AddressGodantic) FieldCountry() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(2),
		godantic.MaxLen(2),
	)
}

func (a *AddressGodantic) FieldZip() godantic.FieldOptions[string] {
	return godantic.Field(godantic.Required[string]())
}

type CustomerGodantic struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Address AddressGodantic `json:"address"`
}

func (c *CustomerGodantic) FieldID() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.Format[string]("uuid"),
	)
}

func (c *CustomerGodantic) FieldName() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(2),
	)
}

func (c *CustomerGodantic) FieldEmail() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.Email(),
	)
}

type OrderItemGodantic struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (i *OrderItemGodantic) FieldSKU() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(3),
	)
}

func (i *OrderItemGodantic) FieldName() godantic.FieldOptions[string] {
	return godantic.Field(godantic.Required[string]())
}

func (i *OrderItemGodantic) FieldQuantity() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Required[int](),
		godantic.Min(1),
	)
}

func (i *OrderItemGodantic) FieldPrice() godantic.FieldOptions[float64] {
	return godantic.Field(
		godantic.Required[float64](),
		godantic.ExclusiveMin(0.0),
	)
}

type OrderGodantic struct {
	ID       string              `json:"id"`
	Customer CustomerGodantic    `json:"customer"`
	Items    []OrderItemGodantic `json:"items"`
	Total    float64             `json:"total"`
	Notes    string              `json:"notes"`
}

func (o *OrderGodantic) FieldID() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.Format[string]("uuid"),
	)
}

func (o *OrderGodantic) FieldTotal() godantic.FieldOptions[float64] {
	return godantic.Field(
		godantic.Required[float64](),
		godantic.ExclusiveMin(0.0),
	)
}

func (o *OrderGodantic) FieldNotes() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.MaxLen(500),
	)
}

func (o *OrderGodantic) FieldItems() godantic.FieldOptions[[]OrderItemGodantic] {
	return godantic.Field(
		godantic.Required[[]OrderItemGodantic](),
		godantic.MinItems[OrderItemGodantic](1),
	)
}

// ----------------------------------------------------------------------------
// Large Config (20+ fields)
// ----------------------------------------------------------------------------

type ConfigGodantic struct {
	AppName        string `json:"app_name"`
	Version        string `json:"version"`
	Environment    string `json:"environment"`
	Debug          bool   `json:"debug"`
	LogLevel       string `json:"log_level"`
	Port           int    `json:"port"`
	Host           string `json:"host"`
	DatabaseURL    string `json:"database_url"`
	RedisURL       string `json:"redis_url"`
	MaxConnections int    `json:"max_connections"`
	Timeout        int    `json:"timeout"`
	RetryCount     int    `json:"retry_count"`
	CacheEnabled   bool   `json:"cache_enabled"`
	CacheTTL       int    `json:"cache_ttl"`
	RateLimit      int    `json:"rate_limit"`
	APIKey         string `json:"api_key"`
	SecretKey      string `json:"secret_key"`
	AllowedOrigins string `json:"allowed_origins"`
	EnableMetrics  bool   `json:"enable_metrics"`
	MetricsPort    int    `json:"metrics_port"`
}

func (c *ConfigGodantic) FieldAppName() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(1),
		godantic.MaxLen(100),
	)
}

func (c *ConfigGodantic) FieldVersion() godantic.FieldOptions[string] {
	return godantic.Field(godantic.Required[string]())
}

func (c *ConfigGodantic) FieldEnvironment() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.OneOf("dev", "staging", "prod"),
	)
}

func (c *ConfigGodantic) FieldLogLevel() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.OneOf("debug", "info", "warn", "error"),
	)
}

func (c *ConfigGodantic) FieldPort() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Required[int](),
		godantic.Min(1),
		godantic.Max(65535),
	)
}

func (c *ConfigGodantic) FieldHost() godantic.FieldOptions[string] {
	return godantic.Field(godantic.Required[string]())
}

func (c *ConfigGodantic) FieldDatabaseURL() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.URL(),
	)
}

func (c *ConfigGodantic) FieldRedisURL() godantic.FieldOptions[string] {
	return godantic.Field(godantic.URL())
}

func (c *ConfigGodantic) FieldMaxConnections() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Min(1),
		godantic.Max(1000),
	)
}

func (c *ConfigGodantic) FieldTimeout() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Min(1),
		godantic.Max(300),
	)
}

func (c *ConfigGodantic) FieldRetryCount() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Min(0),
		godantic.Max(10),
	)
}

func (c *ConfigGodantic) FieldCacheTTL() godantic.FieldOptions[int] {
	return godantic.Field(godantic.Min(0))
}

func (c *ConfigGodantic) FieldRateLimit() godantic.FieldOptions[int] {
	return godantic.Field(godantic.Min(0))
}

func (c *ConfigGodantic) FieldAPIKey() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(32),
	)
}

func (c *ConfigGodantic) FieldSecretKey() godantic.FieldOptions[string] {
	return godantic.Field(
		godantic.Required[string](),
		godantic.MinLen(32),
	)
}

func (c *ConfigGodantic) FieldMetricsPort() godantic.FieldOptions[int] {
	return godantic.Field(
		godantic.Min(1),
		godantic.Max(65535),
	)
}

// ============================================================================
// godantic Test Data
// ============================================================================

var ValidUserGodantic = UserGodantic{
	Name:     "Alice Smith",
	Email:    "alice@example.com",
	Age:      30,
	Website:  "https://alice.dev",
	Username: "alice123",
}

var ValidOrderGodantic = OrderGodantic{
	ID: "550e8400-e29b-41d4-a716-446655440000",
	Customer: CustomerGodantic{
		ID:    "550e8400-e29b-41d4-a716-446655440001",
		Name:  "John Doe",
		Email: "john@example.com",
		Address: AddressGodantic{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "US",
			Zip:     "10001",
		},
	},
	Items: []OrderItemGodantic{
		{SKU: "PROD-001", Name: "Widget", Quantity: 2, Price: 29.99},
		{SKU: "PROD-002", Name: "Gadget", Quantity: 1, Price: 49.99},
	},
	Total: 109.97,
	Notes: "Please deliver before 5pm",
}

var ValidConfigGodantic = ConfigGodantic{
	AppName:        "MyApp",
	Version:        "1.0.0",
	Environment:    "prod",
	Debug:          false,
	LogLevel:       "info",
	Port:           8080,
	Host:           "0.0.0.0",
	DatabaseURL:    "https://db.example.com/mydb",
	RedisURL:       "https://redis.example.com",
	MaxConnections: 100,
	Timeout:        30,
	RetryCount:     3,
	CacheEnabled:   true,
	CacheTTL:       3600,
	RateLimit:      1000,
	APIKey:         "abcdefghijklmnopqrstuvwxyz123456",
	SecretKey:      "123456abcdefghijklmnopqrstuvwxyz",
	AllowedOrigins: "*",
	EnableMetrics:  true,
	MetricsPort:    9090,
}
