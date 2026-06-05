package db

import (
	"context"
	"database/sql"
	"time"
)

// Config defines database-level configuration.
// This keeps tuning explicit and testable.
type Config struct {
	// FilePath is the path to the SQLite .db file
	FilePath string

	// MaxOpenConns limits concurrent SQLite connections
	// Critical for avoiding "database is locked" errors
	MaxOpenConns int

	// MaxIdleConns controls how many idle connections
	// are kept ready for reuse
	MaxIdleConns int

	// ConnMaxLifetime ensures connections are recycled
	// to prevent stale locks and file handles
	ConnMaxLifetime time.Duration
}

// HealthChecker defines a minimal contract
// used by health probes, readiness checks, etc.
type HealthChecker interface {
	HealthCheck(ctx context.Context) error
}

// DB wraps sql.DB to expose only what the app needs.
// This avoids passing raw *sql.DB everywhere.
type DB struct {
	conn *sql.DB
}

// Conn returns the underlying sql.DB when needed
func (d *DB) Conn() *sql.DB {
	return d.conn
}

// HealthCheck verifies the database is reachable
// and able to respond within the provided context.
func (d *DB) HealthCheck(ctx context.Context) error {
	return d.conn.PingContext(ctx)
}

// Close cleanly shuts down all connections
func (d *DB) Close() error {
	return d.conn.Close()
}