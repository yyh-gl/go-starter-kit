package app

import "os"

// IsDev returns true in develop environment
func IsDev() bool {
	return os.Getenv("APP_ENV") == "develop"
}

// IsTest returns true in test environment
func IsTest() bool {
	return os.Getenv("APP_ENV") == "test"
}

// IsStg returns true in staging environment
func IsStg() bool {
	return os.Getenv("APP_ENV") == "staging"
}

// IsPrd returns true in production environment
func IsPrd() bool {
	return os.Getenv("APP_ENV") == "production"
}
