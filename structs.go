package benchmarks

// ============================================================================
// Test Structs - Identical structure, different types per library
// ============================================================================

// ----------------------------------------------------------------------------
// Simple User (5 fields, basic constraints)
// ----------------------------------------------------------------------------

// Pedantigo version
type UserPedantigo struct {
	Name     string `json:"name" pedantigo:"required,min=2,max=100"`
	Email    string `json:"email" pedantigo:"required,email"`
	Age      int    `json:"age" pedantigo:"min=0,max=150"`
	Website  string `json:"website" pedantigo:"url"`
	Username string `json:"username" pedantigo:"alphanum,min=3,max=20"`
}

// go-playground/validator version (same tags)
type UserPlayground struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"min=0,max=150"`
	Website  string `json:"website" validate:"url"`
	Username string `json:"username" validate:"alphanum,min=3,max=20"`
}

// ozzo-validation version (validation done via methods, no tags)
type UserOzzo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Website  string `json:"website"`
	Username string `json:"username"`
}

// Ozzo nested types for Complex benchmark
type AddressOzzo struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

type CustomerOzzo struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address AddressOzzo `json:"address"`
}

type OrderItemOzzo struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderOzzo struct {
	ID       string          `json:"id"`
	Customer CustomerOzzo    `json:"customer"`
	Items    []OrderItemOzzo `json:"items"`
	Total    float64         `json:"total"`
	Notes    string          `json:"notes"`
}

// Ozzo Config type for Large benchmark
type ConfigOzzo struct {
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

// huma version
type UserHuma struct {
	Name     string `json:"name" minLength:"2" maxLength:"100" required:"true"`
	Email    string `json:"email" format:"email" required:"true"`
	Age      int    `json:"age" minimum:"0" maximum:"150"`
	Website  string `json:"website" format:"uri"`
	Username string `json:"username" minLength:"3" maxLength:"20" pattern:"^[a-zA-Z0-9]+$"`
}

// ----------------------------------------------------------------------------
// Complex Order (nested structs, arrays)
// ----------------------------------------------------------------------------

// Pedantigo nested types
type AddressPedantigo struct {
	Street  string `json:"street" pedantigo:"required,min=5"`
	City    string `json:"city" pedantigo:"required"`
	Country string `json:"country" pedantigo:"required,len=2"`
	Zip     string `json:"zip" pedantigo:"required"`
}

type CustomerPedantigo struct {
	ID      string           `json:"id" pedantigo:"required,uuid"`
	Name    string           `json:"name" pedantigo:"required,min=2"`
	Email   string           `json:"email" pedantigo:"required,email"`
	Address AddressPedantigo `json:"address" pedantigo:"required"`
}

type OrderItemPedantigo struct {
	SKU      string  `json:"sku" pedantigo:"required,min=3"`
	Name     string  `json:"name" pedantigo:"required"`
	Quantity int     `json:"quantity" pedantigo:"required,min=1"`
	Price    float64 `json:"price" pedantigo:"required,gt=0"`
}

type OrderPedantigo struct {
	ID       string               `json:"id" pedantigo:"required,uuid"`
	Customer CustomerPedantigo    `json:"customer" pedantigo:"required"`
	Items    []OrderItemPedantigo `json:"items" pedantigo:"required,min=1,dive"`
	Total    float64              `json:"total" pedantigo:"required,gt=0"`
	Notes    string               `json:"notes" pedantigo:"max=500"`
}

// Playground nested types
type AddressPlayground struct {
	Street  string `json:"street" validate:"required,min=5"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required,len=2"`
	Zip     string `json:"zip" validate:"required"`
}

type CustomerPlayground struct {
	ID      string            `json:"id" validate:"required,uuid"`
	Name    string            `json:"name" validate:"required,min=2"`
	Email   string            `json:"email" validate:"required,email"`
	Address AddressPlayground `json:"address" validate:"required"`
}

type OrderItemPlayground struct {
	SKU      string  `json:"sku" validate:"required,min=3"`
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity" validate:"required,min=1"`
	Price    float64 `json:"price" validate:"required,gt=0"`
}

type OrderPlayground struct {
	ID       string                `json:"id" validate:"required,uuid"`
	Customer CustomerPlayground    `json:"customer" validate:"required"`
	Items    []OrderItemPlayground `json:"items" validate:"required,min=1,dive"`
	Total    float64               `json:"total" validate:"required,gt=0"`
	Notes    string                `json:"notes" validate:"max=500"`
}

// ----------------------------------------------------------------------------
// Large Config (20+ fields)
// ----------------------------------------------------------------------------

type ConfigPedantigo struct {
	AppName        string `json:"app_name" pedantigo:"required,min=1,max=100"`
	Version        string `json:"version" pedantigo:"required"`
	Environment    string `json:"environment" pedantigo:"required,oneof=dev staging prod"`
	Debug          bool   `json:"debug"`
	LogLevel       string `json:"log_level" pedantigo:"oneof=debug info warn error"`
	Port           int    `json:"port" pedantigo:"required,min=1,max=65535"`
	Host           string `json:"host" pedantigo:"required"`
	DatabaseURL    string `json:"database_url" pedantigo:"required,url"`
	RedisURL       string `json:"redis_url" pedantigo:"url"`
	MaxConnections int    `json:"max_connections" pedantigo:"min=1,max=1000"`
	Timeout        int    `json:"timeout" pedantigo:"min=1,max=300"`
	RetryCount     int    `json:"retry_count" pedantigo:"min=0,max=10"`
	CacheEnabled   bool   `json:"cache_enabled"`
	CacheTTL       int    `json:"cache_ttl" pedantigo:"min=0"`
	RateLimit      int    `json:"rate_limit" pedantigo:"min=0"`
	APIKey         string `json:"api_key" pedantigo:"required,min=32"`
	SecretKey      string `json:"secret_key" pedantigo:"required,min=32"`
	AllowedOrigins string `json:"allowed_origins"`
	EnableMetrics  bool   `json:"enable_metrics"`
	MetricsPort    int    `json:"metrics_port" pedantigo:"min=1,max=65535"`
}

type ConfigPlayground struct {
	AppName        string `json:"app_name" validate:"required,min=1,max=100"`
	Version        string `json:"version" validate:"required"`
	Environment    string `json:"environment" validate:"required,oneof=dev staging prod"`
	Debug          bool   `json:"debug"`
	LogLevel       string `json:"log_level" validate:"oneof=debug info warn error"`
	Port           int    `json:"port" validate:"required,min=1,max=65535"`
	Host           string `json:"host" validate:"required"`
	DatabaseURL    string `json:"database_url" validate:"required,url"`
	RedisURL       string `json:"redis_url" validate:"url"`
	MaxConnections int    `json:"max_connections" validate:"min=1,max=1000"`
	Timeout        int    `json:"timeout" validate:"min=1,max=300"`
	RetryCount     int    `json:"retry_count" validate:"min=0,max=10"`
	CacheEnabled   bool   `json:"cache_enabled"`
	CacheTTL       int    `json:"cache_ttl" validate:"min=0"`
	RateLimit      int    `json:"rate_limit" validate:"min=0"`
	APIKey         string `json:"api_key" validate:"required,min=32"`
	SecretKey      string `json:"secret_key" validate:"required,min=32"`
	AllowedOrigins string `json:"allowed_origins"`
	EnableMetrics  bool   `json:"enable_metrics"`
	MetricsPort    int    `json:"metrics_port" validate:"min=1,max=65535"`
}

// ============================================================================
// Test Data - Valid instances
// ============================================================================

var ValidUserJSON = []byte(`{
	"name": "Alice Smith",
	"email": "alice@example.com",
	"age": 30,
	"website": "https://alice.dev",
	"username": "alice123"
}`)

var ValidUserPedantigo = UserPedantigo{
	Name:     "Alice Smith",
	Email:    "alice@example.com",
	Age:      30,
	Website:  "https://alice.dev",
	Username: "alice123",
}

var ValidUserPlayground = UserPlayground{
	Name:     "Alice Smith",
	Email:    "alice@example.com",
	Age:      30,
	Website:  "https://alice.dev",
	Username: "alice123",
}

var ValidUserOzzo = UserOzzo{
	Name:     "Alice Smith",
	Email:    "alice@example.com",
	Age:      30,
	Website:  "https://alice.dev",
	Username: "alice123",
}

var ValidOrderOzzo = OrderOzzo{
	ID: "550e8400-e29b-41d4-a716-446655440000",
	Customer: CustomerOzzo{
		ID:    "550e8400-e29b-41d4-a716-446655440001",
		Name:  "John Doe",
		Email: "john@example.com",
		Address: AddressOzzo{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "US",
			Zip:     "10001",
		},
	},
	Items: []OrderItemOzzo{
		{SKU: "PROD-001", Name: "Widget", Quantity: 2, Price: 29.99},
		{SKU: "PROD-002", Name: "Gadget", Quantity: 1, Price: 49.99},
	},
	Total: 109.97,
	Notes: "Please deliver before 5pm",
}

var ValidConfigOzzo = ConfigOzzo{
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

var ValidUserMap = map[string]any{
	"name":     "Alice Smith",
	"email":    "alice@example.com",
	"age":      30,
	"website":  "https://alice.dev",
	"username": "alice123",
}

var ValidOrderJSON = []byte(`{
	"id": "550e8400-e29b-41d4-a716-446655440000",
	"customer": {
		"id": "550e8400-e29b-41d4-a716-446655440001",
		"name": "John Doe",
		"email": "john@example.com",
		"address": {
			"street": "123 Main Street",
			"city": "New York",
			"country": "US",
			"zip": "10001"
		}
	},
	"items": [
		{"sku": "PROD-001", "name": "Widget", "quantity": 2, "price": 29.99},
		{"sku": "PROD-002", "name": "Gadget", "quantity": 1, "price": 49.99}
	],
	"total": 109.97,
	"notes": "Please deliver before 5pm"
}`)

var ValidOrderPedantigo = OrderPedantigo{
	ID: "550e8400-e29b-41d4-a716-446655440000",
	Customer: CustomerPedantigo{
		ID:    "550e8400-e29b-41d4-a716-446655440001",
		Name:  "John Doe",
		Email: "john@example.com",
		Address: AddressPedantigo{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "US",
			Zip:     "10001",
		},
	},
	Items: []OrderItemPedantigo{
		{SKU: "PROD-001", Name: "Widget", Quantity: 2, Price: 29.99},
		{SKU: "PROD-002", Name: "Gadget", Quantity: 1, Price: 49.99},
	},
	Total: 109.97,
	Notes: "Please deliver before 5pm",
}

var ValidOrderPlayground = OrderPlayground{
	ID: "550e8400-e29b-41d4-a716-446655440000",
	Customer: CustomerPlayground{
		ID:    "550e8400-e29b-41d4-a716-446655440001",
		Name:  "John Doe",
		Email: "john@example.com",
		Address: AddressPlayground{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "US",
			Zip:     "10001",
		},
	},
	Items: []OrderItemPlayground{
		{SKU: "PROD-001", Name: "Widget", Quantity: 2, Price: 29.99},
		{SKU: "PROD-002", Name: "Gadget", Quantity: 1, Price: 49.99},
	},
	Total: 109.97,
	Notes: "Please deliver before 5pm",
}

var ValidConfigPedantigo = ConfigPedantigo{
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

var ValidConfigPlayground = ConfigPlayground{
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
