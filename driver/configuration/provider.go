package configuration

// Provider contains provide all configs to gotham
type Provider interface {
	// Data Source Name (DSN) provides connectivity to a database through an Driver
	DSN() string
}
