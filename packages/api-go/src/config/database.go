package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
)

var db *sqlx.DB

func initializeDatabase() {
	host := GetString("database.relational.host")
	port := GetInt("database.relational.port")
	dbName := GetString("database.relational.database-name")
	user := GetString("database.relational.auth.user")
	password := GetString("database.relational.auth.password")

	connDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db = otelsqlx.MustOpen("postgres", connDsn)
}

func GetDb() *sqlx.DB {
	return db
}
