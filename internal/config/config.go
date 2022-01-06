package config

import "os"

const (
	// HttpPort - http server port
	HttpPort = "PORT"

	// DbHost - db host key
	DbHost = "DB_HOST"
	// DbPort - db port key
	DbPort = "DB_PORT"
	// DbUser - db user key
	DbUser = "DB_USER"
	// DbPass - db pass key
	DbPass = "DB_PASS"
	// DbName - db name key
	DbName = "DB_NAME"
	// DbSsl - db ssl mode key [disable/enable]
	DbSsl = "DB_SSL"
)

// Config - structure to store all organized config values
type Config struct {
	DbCredentials *DatabaseCredentials
	Http          *Http
}

// Http - structure for storing http server configs
type Http struct {
	Port string
}

// DatabaseCredentials - structure to store all database config values
type DatabaseCredentials struct {
	Host        string
	Port        string
	Name        string
	User        string
	Pass        string
	SslMode     string
	AutoMigrate bool
}

// InitDbCredentials - init database credentials from .env and return a new struct
func InitDbCredentials() *DatabaseCredentials {
	return &DatabaseCredentials{
		Host:        os.Getenv(DbHost),
		Name:        os.Getenv(DbName),
		Port:        os.Getenv(DbPort),
		Pass:        os.Getenv(DbPass),
		User:        os.Getenv(DbUser),
		SslMode:     os.Getenv(DbSsl),
		AutoMigrate: true,
	}
}

// InitHttpConfig - return pointer of http server config
func InitHttpConfig() *Http {
	return &Http{
		Port: os.Getenv(HttpPort),
	}
}

// Init - returns a pointer of Config structure with defined values
func Init() *Config {
	return &Config{
		DbCredentials: InitDbCredentials(),
		Http:          InitHttpConfig(),
	}
}
