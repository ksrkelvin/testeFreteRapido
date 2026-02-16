package config_test

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mockDriverFail(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
	return nil, errors.New("simulação de falha")
}

func mockDriverSuccess(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	dialector := postgres.New(postgres.Config{
		Conn: dbMock,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
