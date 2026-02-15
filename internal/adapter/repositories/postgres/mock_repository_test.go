package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro criando sqlmock: %s", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: dbMock,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("erro abrindo gorm DB: %s", err)
	}

	return gormDB, mock
}
