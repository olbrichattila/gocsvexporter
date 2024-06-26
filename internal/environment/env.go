// Package environment returns with the values for database connection
package environment

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFileName     = "./.env"
	envDBConnection = "DB_CONNECTION"
	envDBUserName   = "DB_USERNAME"
	envDBPassword   = "DB_PASSWORD"
	envDBHost       = "DB_HOST"
	envDBPort       = "DB_PORT"
	envDBDatabase   = "DB_DATABASE"
	envDBSSLMode    = "DB_SSLMODE"
)

// Enver environment variable retrival interface
type Enver interface {
	DBConnection() string
	DBUserName() string
	DBPassword() string
	DBHost() string
	DBPort() string
	DBDatabase() string
	DBSSLMode() string
}

// New creates an environment manager
func New() (Enver, error) {
	env := &env{}
	// TODO de we handle this? If no error we use os exports
	err := env.loadEnv()
	if err != nil {
		return nil, err
	}
	return env, nil
}

type env struct {
}

func (e *env) DBConnection() string {
	return os.Getenv(envDBConnection)
}

func (e *env) DBUserName() string {
	return os.Getenv(envDBUserName)
}

func (e *env) DBPassword() string {
	return os.Getenv(envDBPassword)
}

func (e *env) DBHost() string {
	return os.Getenv(envDBHost)
}

func (e *env) DBPort() string {
	return os.Getenv(envDBPort)
}

func (e *env) DBDatabase() string {
	return os.Getenv(envDBDatabase)
}

func (e *env) DBSSLMode() string {
	return os.Getenv(envDBSSLMode)
}

func (e *env) loadEnv() error {
	_, err := os.Stat(envFileName)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return godotenv.Load(envFileName)
}
