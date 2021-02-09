package databases

import "cmd/src/main.go/lib/services/databases/postgresdb"

type Databases struct {
	PostgresDb postgresdb.PostgresDb
}

func StartAllDatabase() Databases {
	var databases Databases

	databases.PostgresDb = postgresdb.StartPostgresDB()

	return databases
}
