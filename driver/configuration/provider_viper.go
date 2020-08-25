package configuration

import (
	"log"

	"github.com/spf13/viper"
)

// ViperProvider implements Provider interface
type ViperProvider struct{}

var _ Provider = new(ViperProvider)

const (
	ViperKeyDSN = "dsn"
	FallBackDSN = ""
)

// NewViperProvider creates a ViperProvider instance
func NewViperProvider() *ViperProvider{
	return &ViperProvider{}
}

// DSN returns the DSN from viper config or the fallback value
func (p *ViperProvider) DSN() string {
	// TODO: Refactor it: consider to move this logic to external package
	var dsn = viper.GetString(ViperKeyDSN)
	if len(dsn) == 0 {
		return FallBackDSN
	}

	if len(dsn) > 0 {
		return dsn
	}

	log.Fatal("dsn must be set")
	return ""
}
