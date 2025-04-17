package testdb

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	db     *sqlx.DB
	once   sync.Once
	dbErr  error
	dbUser = "testuser"
	dbPass = "testpass"
	dbName = "testdb"
)

// StartContainer starts the test database once for all integration tests.
func SetupTestDatabase() (*sqlx.DB, error) {
	once.Do(func() {
		ctx := context.Background()
		pgContainer, err := postgres.Run(ctx,
			"postgres:16",
			postgres.WithInitScripts(filepath.Join("src", "tests", "scripts", "init-db.sql")),
			postgres.WithDatabase(dbName),
			postgres.WithUsername(dbUser),
			postgres.WithPassword(dbPass),
			testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
				ContainerRequest: testcontainers.ContainerRequest{
					ExposedPorts: []string{"5433/tcp"},
				},
			}),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5*time.Second)),
		)
		defer func() {
			if err := testcontainers.TerminateContainer(pgContainer); err != nil {
				log.Printf("failed to terminate container: %s", err)
			}
		}()
		if err != nil {
			dbErr = fmt.Errorf("failed to start container: %w", err)
			return
		}

		// Wait for startup
		time.Sleep(2 * time.Second)

		connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			dbErr = fmt.Errorf("failed to get connection string: %w", err)
			return
		}

		db = sqlx.MustConnect("postgres", connStr)
	})

	return db, dbErr
}
