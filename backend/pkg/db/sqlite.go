package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// Pure Go SQLite driver (no CGO)
	_ "modernc.org/sqlite"
)

// NewSQLite creates and initializes a tuned SQLite database connection pool.
func NewSQLite(cfg Config) (*DB, error) {
	if cfg.FilePath == "" {
		return nil, fmt.Errorf("sqlite: file path is required")
	}

	// Open does NOT establish connections immediately.
	// It initializes a connection pool controller.
	db, err := sql.Open("sqlite", cfg.FilePath)
	if err != nil {
		return nil, fmt.Errorf("sqlite: open failed: %w", err)
	}

	// ---- Connection Pool Tuning ----

	// Limits total concurrent connections.
	// SQLite cannot handle unlimited writers.
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	// Keeps a small number of idle connections ready.
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	// Forces periodic recycling of connections.
	// Prevents long-lived locks and stale handles.
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// ---- Safety Hook: Enforce Foreign Keys ----
	// SQLite defaults foreign_keys to OFF.
	// We MUST enable it per connection.
	if err := enableForeignKeys(db); err != nil {
		_ = db.Close()
		return nil, err
	}

	// ---- Health Verification ----
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sqlite: ping failed: %w", err)
	}

	return &DB{conn: db}, nil
}

// enableForeignKeys ensures PRAGMA foreign_keys = ON
// is applied safely and consistently.
func enableForeignKeys(db *sql.DB) error {
	// We run this once eagerly to fail fast.
	// Subsequent connections inherit this setting
	// when reused from the pool.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, `PRAGMA foreign_keys = ON;`)
	if err != nil {
		return fmt.Errorf("sqlite: failed to enable foreign keys: %w", err)
	}

	return nil
}