package postgresdb

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"cmd/src/main.go/lib/utils"

	"github.com/jackc/pgx/v4"
)

type PostgresDb struct {
	*pgx.Conn
}

func StartPostgresDB() PostgresDb {
	config, _ := pgx.ParseConfig("")
	config.Database = os.Getenv("DATABASE_NAME")
	config.Password = os.Getenv("DB_PASSWORD")
	config.User = os.Getenv("DB_USER")
	config.Host = os.Getenv("POSTGRES_HOST")
	port, _ := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 16)
	config.Port = uint16(port)
	utils.SuccesLogger(fmt.Sprintf("Iniciando conexão com o banco Postgres na porta %v", os.Getenv("DATABASE_NAME")))
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	utils.SuccesLogger("Conexão feita com sucesso")
	return PostgresDb{conn}
}
