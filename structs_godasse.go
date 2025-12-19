package benchmarks

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

// ============================================================================
// godasse Types (Interface-based validation via Validate() method)
// ============================================================================

// Helper validators
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var alphanumRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

// ----------------------------------------------------------------------------
// Simple User (5 fields)
// ----------------------------------------------------------------------------

type UserGodasse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Website  string `json:"website"`
	Username string `json:"username"`
}

func (u *UserGodasse) Validate() error {
	// Name: required, min=2, max=100
	if u.Name == "" {
		return errors.New("name is required")
	}
	if len(u.Name) < 2 || len(u.Name) > 100 {
		return errors.New("name must be between 2 and 100 characters")
	}

	// Email: required, email format
	if u.Email == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	// Age: min=0, max=150
	if u.Age < 0 || u.Age > 150 {
		return errors.New("age must be between 0 and 150")
	}

	// Website: URL format (optional)
	if u.Website != "" {
		if _, err := url.ParseRequestURI(u.Website); err != nil {
			return errors.New("invalid website URL")
		}
	}

	// Username: alphanum, min=3, max=20
	if u.Username != "" {
		if len(u.Username) < 3 || len(u.Username) > 20 {
			return errors.New("username must be between 3 and 20 characters")
		}
		if !alphanumRegex.MatchString(u.Username) {
			return errors.New("username must be alphanumeric")
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
// Complex Order (nested structs, arrays)
// ----------------------------------------------------------------------------

type AddressGodasse struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

func (a *AddressGodasse) Validate() error {
	if a.Street == "" {
		return errors.New("street is required")
	}
	if len(a.Street) < 5 {
		return errors.New("street must be at least 5 characters")
	}
	if a.City == "" {
		return errors.New("city is required")
	}
	if a.Country == "" {
		return errors.New("country is required")
	}
	if len(a.Country) != 2 {
		return errors.New("country must be exactly 2 characters")
	}
	if a.Zip == "" {
		return errors.New("zip is required")
	}
	return nil
}

type CustomerGodasse struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Address AddressGodasse `json:"address"`
}

func (c *CustomerGodasse) Validate() error {
	if c.ID == "" {
		return errors.New("customer id is required")
	}
	if !uuidRegex.MatchString(strings.ToLower(c.ID)) {
		return errors.New("customer id must be a valid UUID")
	}
	if c.Name == "" {
		return errors.New("customer name is required")
	}
	if len(c.Name) < 2 {
		return errors.New("customer name must be at least 2 characters")
	}
	if c.Email == "" {
		return errors.New("customer email is required")
	}
	if !emailRegex.MatchString(c.Email) {
		return errors.New("invalid customer email format")
	}
	if err := c.Address.Validate(); err != nil {
		return err
	}
	return nil
}

type OrderItemGodasse struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (i *OrderItemGodasse) Validate() error {
	if i.SKU == "" {
		return errors.New("item sku is required")
	}
	if len(i.SKU) < 3 {
		return errors.New("item sku must be at least 3 characters")
	}
	if i.Name == "" {
		return errors.New("item name is required")
	}
	if i.Quantity < 1 {
		return errors.New("item quantity must be at least 1")
	}
	if i.Price <= 0 {
		return errors.New("item price must be greater than 0")
	}
	return nil
}

type OrderGodasse struct {
	ID       string             `json:"id"`
	Customer CustomerGodasse    `json:"customer"`
	Items    []OrderItemGodasse `json:"items"`
	Total    float64            `json:"total"`
	Notes    string             `json:"notes"`
}

func (o *OrderGodasse) Validate() error {
	if o.ID == "" {
		return errors.New("order id is required")
	}
	if !uuidRegex.MatchString(strings.ToLower(o.ID)) {
		return errors.New("order id must be a valid UUID")
	}
	if err := o.Customer.Validate(); err != nil {
		return err
	}
	if len(o.Items) < 1 {
		return errors.New("order must have at least 1 item")
	}
	for i := range o.Items {
		if err := o.Items[i].Validate(); err != nil {
			return err
		}
	}
	if o.Total <= 0 {
		return errors.New("order total must be greater than 0")
	}
	if len(o.Notes) > 500 {
		return errors.New("notes must be at most 500 characters")
	}
	return nil
}

// ----------------------------------------------------------------------------
// Large Config (20+ fields)
// ----------------------------------------------------------------------------

type ConfigGodasse struct {
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

func (c *ConfigGodasse) Validate() error {
	// AppName: required, min=1, max=100
	if c.AppName == "" {
		return errors.New("app_name is required")
	}
	if len(c.AppName) < 1 || len(c.AppName) > 100 {
		return errors.New("app_name must be between 1 and 100 characters")
	}

	// Version: required
	if c.Version == "" {
		return errors.New("version is required")
	}

	// Environment: required, oneof=dev,staging,prod
	if c.Environment == "" {
		return errors.New("environment is required")
	}
	switch c.Environment {
	case "dev", "staging", "prod":
	default:
		return errors.New("environment must be dev, staging, or prod")
	}

	// LogLevel: oneof=debug,info,warn,error
	if c.LogLevel != "" {
		switch c.LogLevel {
		case "debug", "info", "warn", "error":
		default:
			return errors.New("log_level must be debug, info, warn, or error")
		}
	}

	// Port: required, min=1, max=65535
	if c.Port < 1 || c.Port > 65535 {
		return errors.New("port must be between 1 and 65535")
	}

	// Host: required
	if c.Host == "" {
		return errors.New("host is required")
	}

	// DatabaseURL: required, URL format
	if c.DatabaseURL == "" {
		return errors.New("database_url is required")
	}
	if _, err := url.ParseRequestURI(c.DatabaseURL); err != nil {
		return errors.New("invalid database_url format")
	}

	// RedisURL: optional URL format
	if c.RedisURL != "" {
		if _, err := url.ParseRequestURI(c.RedisURL); err != nil {
			return errors.New("invalid redis_url format")
		}
	}

	// MaxConnections: min=1, max=1000
	if c.MaxConnections != 0 && (c.MaxConnections < 1 || c.MaxConnections > 1000) {
		return errors.New("max_connections must be between 1 and 1000")
	}

	// Timeout: min=1, max=300
	if c.Timeout != 0 && (c.Timeout < 1 || c.Timeout > 300) {
		return errors.New("timeout must be between 1 and 300")
	}

	// RetryCount: min=0, max=10
	if c.RetryCount < 0 || c.RetryCount > 10 {
		return errors.New("retry_count must be between 0 and 10")
	}

	// CacheTTL: min=0
	if c.CacheTTL < 0 {
		return errors.New("cache_ttl must be non-negative")
	}

	// RateLimit: min=0
	if c.RateLimit < 0 {
		return errors.New("rate_limit must be non-negative")
	}

	// APIKey: required, min=32
	if c.APIKey == "" {
		return errors.New("api_key is required")
	}
	if len(c.APIKey) < 32 {
		return errors.New("api_key must be at least 32 characters")
	}

	// SecretKey: required, min=32
	if c.SecretKey == "" {
		return errors.New("secret_key is required")
	}
	if len(c.SecretKey) < 32 {
		return errors.New("secret_key must be at least 32 characters")
	}

	// MetricsPort: min=1, max=65535
	if c.MetricsPort != 0 && (c.MetricsPort < 1 || c.MetricsPort > 65535) {
		return errors.New("metrics_port must be between 1 and 65535")
	}

	return nil
}

// ============================================================================
// godasse Test Data
// ============================================================================

var ValidUserGodasse = UserGodasse{
	Name:     "Alice Smith",
	Email:    "alice@example.com",
	Age:      30,
	Website:  "https://alice.dev",
	Username: "alice123",
}

var ValidOrderGodasse = OrderGodasse{
	ID: "550e8400-e29b-41d4-a716-446655440000",
	Customer: CustomerGodasse{
		ID:    "550e8400-e29b-41d4-a716-446655440001",
		Name:  "John Doe",
		Email: "john@example.com",
		Address: AddressGodasse{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "US",
			Zip:     "10001",
		},
	},
	Items: []OrderItemGodasse{
		{SKU: "PROD-001", Name: "Widget", Quantity: 2, Price: 29.99},
		{SKU: "PROD-002", Name: "Gadget", Quantity: 1, Price: 49.99},
	},
	Total: 109.97,
	Notes: "Please deliver before 5pm",
}

var ValidConfigGodasse = ConfigGodasse{
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
