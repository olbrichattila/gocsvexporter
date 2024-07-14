package main

import "fmt"

func displayHelp() {
	fmt.Print(`Usage: "csvexporter . <tablebane> <csvfilename> <separator>"

where the Separator is optional, if not set then it defaults to (,)

The database connection can be set by .env.csvexporter file, or if not exists, the application tries to get it from linux environment variables what you can set:
Example:

------------------------------------------------
export DB_CONNECTION=sqlite
export DB_DATABASE=./data/database.sqlite
------------------------------------------------


Possible .env.csvexporter settings (please see .env.* examples as well)


### SqLite:
------------------------------------------------
DB_CONNECTION=sqlite
DB_DATABASE=./data/database.sqlite
------------------------------------------------

### MySql:
------------------------------------------------
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=migrator
DB_USERNAME=migrator
DB_PASSWORD=H8E7kU8Y
------------------------------------------------

### PostgresQl:
------------------------------------------------
DB_CONNECTION=pgsql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_SSLMODE=disable
------------------------------------------------

### FirebirdSQL
------------------------------------------------
DB_CONNECTION=firebird
DB_HOST=127.0.0.1
DB_PORT=3050
DB_DATABASE=/firebird/data/employee.fdb
DB_USERNAME=SYSDBA
DB_PASSWORD=masterkey
------------------------------------------------
`)
}
