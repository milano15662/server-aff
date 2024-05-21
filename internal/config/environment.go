package config

import "os"

const (
	EnvRelease     = "release"
	EnvDevelopment = "development"
)

var actualEnvironment = os.Getenv("ENVIRONMENT")

// IsRelease ...
func IsRelease() bool {
	return actualEnvironment == EnvRelease
}

// IsDevelopment ...
func IsDevelopment() bool {
	return actualEnvironment == EnvDevelopment
}
