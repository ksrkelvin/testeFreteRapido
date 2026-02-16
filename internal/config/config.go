package config

import "gorm.io/gorm"

type Config struct {
	DBConnURL        string
	AuthToken        string
	PlataformCode    string
	RegisteredNumber string
	MigrateDB        func(db *gorm.DB)
}

func (c *Config) GetDBConnURL() string        { return c.DBConnURL }
func (c *Config) GetAuthToken() string        { return c.AuthToken }
func (c *Config) GetPlataformCode() string    { return c.PlataformCode }
func (c *Config) GetRegisteredNumber() string { return c.RegisteredNumber }

type AppProvider interface {
	GetDBConnURL() string
	GetAuthToken() string
	GetPlataformCode() string
	GetRegisteredNumber() string

	NewDatabase(DBConn) (*gorm.DB, error)
}

type DBConn func(dsn string, config *gorm.Config) (*gorm.DB, error)
