package cockroach

import (
	"context"
	"errors"
	"fmt"

	"transaction/internal/helper/di"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do"
)

// NewCockroachConnection creates and returns a new connection pool to CockroachDB
func NewCockroachConnectionPool(i *do.Injector) (*pgxpool.Pool, error) {
	var cfg CRDBConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}

	// Build the connection string
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
	)

	// Create a pool configuration
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Apply connection pool settings
	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolConfig.HealthCheckPeriod = cfg.HealthCheckPeriod
	poolConfig.ConnConfig.ConnectTimeout = cfg.Timeout

	// Create a timeout context for connection establishment
	ctx := context.Background()
	connCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(connCtx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(connCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

// NewCockroachSingleConnection creates and returns a single connection to CockroachDB
func NewCockroachSingleConnection(i *do.Injector) (*pgx.Conn, error) {
	var cfg CRDBConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}

	// Build the connection string
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
	)

	// Create a timeout context for connection establishment
	ctx := context.Background()
	connCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	// Connect to the database
	conn, err := pgx.Connect(connCtx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := conn.Ping(connCtx); err != nil {
		_ = conn.Close(ctx)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, nil
}
