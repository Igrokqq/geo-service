package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadEnvData(t *testing.T) {
	t.Setenv(HttpPort, HttpPort)
	t.Setenv(DbName, DbName)
	t.Setenv(DbHost, DbHost)
	t.Setenv(DbSsl, DbSsl)
	t.Setenv(DbUser, DbUser)
	t.Setenv(DbPass, DbPass)
	t.Setenv(DbPort, DbPort)
}

func TestInitDatabaseCredentials(t *testing.T) {
	loadEnvData(t)

	expected := &DatabaseCredentials{
		Name:        DbName,
		Host:        DbHost,
		Port:        DbPort,
		User:        DbUser,
		Pass:        DbPass,
		SslMode:     DbSsl,
		AutoMigrate: true,
	}

	actual := InitDbCredentials()

	assert.Equal(t, *expected, *actual)
}

func TestInitHttpConfig(t *testing.T) {
	loadEnvData(t)

	expected := &Http{
		Port: HttpPort,
	}

	actual := InitHttpConfig()

	assert.Equal(t, *expected, *actual)
}

func TestInit(t *testing.T) {
	loadEnvData(t)

	dbCredentials := &DatabaseCredentials{
		Name:        DbName,
		Host:        DbHost,
		Port:        DbPort,
		User:        DbUser,
		Pass:        DbPass,
		SslMode:     DbSsl,
		AutoMigrate: true,
	}

	httpConfig := &Http{
		Port: HttpPort,
	}

	expected := &Config{
		DbCredentials: dbCredentials,
		Http:          httpConfig,
	}

	actual := Init()

	assert.Equal(t, *expected, *actual)
}
