package cockroach

import "time"

type CRDBConfig struct {
	Port              uint16        `env:"CRDB_PORT" envDefault:"26257"`
	Username          string        `env:"CRDB_USERNAME" envDefault:"banking"`
	Password          string        `env:"CRDB_PASSWORD" envDefault:""`
	Database          string        `env:"CRDB_DB_NAME" envDefault:"banking"`
	SSLMode           string        `env:"CRDB_SSL_MODE" envDefault:"disable"`
	SSLCert           string        `env:"CRDB_SSL_CERT" envDefault:""`
	Host              string        `env:"CRDB_HOST" envDefault:"localhost"`
	EnableMetrics     bool          `env:"CRDB_ENABLE_METRICS" envDefault:"false"`
	EnableTraces      bool          `env:"CRDB_ENABLE_TRACER" envDefault:"false"`
	Timeout           time.Duration `env:"CRDB_TIMEOUT" envDefault:"10s"`
	MaxConns          int32         `env:"CRDB_MAX_CONNS" envDefault:"10"`
	MinConns          int32         `env:"CRDB_MIN_CONNS" envDefault:"1"`
	MaxConnLifetime   time.Duration `env:"CRDB_MAX_CONN_LIFETIME" envDefault:"1m"`
	MaxConnIdleTime   time.Duration `env:"CRDB_MAX_CONN_IDLE_TIME" envDefault:"1m"`
	HealthCheckPeriod time.Duration `env:"CRDB_HEALTH_CHECK_PERIOD" envDefault:"1m"`
}
