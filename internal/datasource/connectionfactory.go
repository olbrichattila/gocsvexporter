package datasource

import (
	"csvdownloader/internal/environment"
	"fmt"

	// This blank import is necessary to have the driver
	_ "github.com/mattn/go-sqlite3"
	// This blank import is necessary to have the driver
	_ "github.com/lib/pq"

	// This blank import is necessary to have the driver
	_ "github.com/go-sql-driver/mysql"

	// This blank import is necessary to have the driver
	_ "github.com/nakagami/firebirdsql"
)

const (
	driverNameFirebird = "firebirdsql"
	driverNameSqLite   = "sqlite3"
	driverNameMySQL    = "mysql"
	driverNamePostgres = "postgres"

	dbConnectionTypeSqLite   = "sqlite"
	dbConnectionTypeMySQL    = "mysql"
	dbConnectionTypePgSQL    = "pgsql"
	dbConnectionTypeFirebird = "firebird"
)

// GetDbConnectionParams converts to database connection string, and quotes from .env or linux export
// returns connection string, driver name, quote character and error if failed
func GetDbConnectionParams(env environment.Enver) (string, string, string, error) {
	switch env.DBConnection() {
	case dbConnectionTypeSqLite:
		return env.DBDatabase(), driverNameSqLite, `"`, nil
	case dbConnectionTypeMySQL:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			env.DBUserName(),
			env.DBPassword(),
			env.DBHost(),
			env.DBPort(),
			env.DBDatabase(),
		), driverNameMySQL, "`", nil
	case dbConnectionTypePgSQL:
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			env.DBUserName(),
			env.DBPassword(),
			env.DBHost(),
			env.DBPort(),
			env.DBDatabase(),
			env.DBSSLMode(),
		), driverNamePostgres, `"`, nil

	case dbConnectionTypeFirebird:
		return fmt.Sprintf(
			"%s:%s@%s:%s%s",
			env.DBUserName(),
			env.DBPassword(),
			env.DBHost(),
			env.DBPort(),
			env.DBDatabase(),
		), driverNameFirebird, `"`, nil
	default:
		return "", "", "", fmt.Errorf("the %s connection not implemented", env.DBConnection())

	}
}
