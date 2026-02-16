package main

import (
	"testeFreteRapido/internal/config"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockConfig struct {
	forceFreteError bool
}

func (m *mockConfig) GetDBConnURL() string { return "dummy-db" }
func (m *mockConfig) GetAuthToken() string {
	if m.forceFreteError {
		return ""
	}
	return "dummy-token"
}
func (m *mockConfig) GetPlataformCode() string    { return "dummy-code" }
func (m *mockConfig) GetRegisteredNumber() string { return "123" }
func (m *mockConfig) NewDatabase(dbConn config.DBConn) (*gorm.DB, error) {
	db, _ := SetupMockDB(nil)
	return db, nil
}

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

func Test_setupServer(t *testing.T) {
	tests := []struct {
		name        string
		cfg         config.AppProvider
		dbNil       bool
		expectError bool
	}{
		{"success", &mockConfig{}, false, false},
		{"db is nil", &mockConfig{}, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var db *gorm.DB
			if !tt.dbNil {
				db, _ = SetupMockDB(t)
			}

			engine, err := setupServer(tt.cfg, db)

			if tt.expectError {
				if err == nil {
					t.Fatalf("esperava erro mas recebeu nil")
				}
				if engine != nil {
					t.Fatalf("esperava engine nil em caso de erro")
				}
			} else {
				if err != nil {
					t.Fatalf("erro inesperado: %v", err)
				}
				if engine == nil {
					t.Fatal("esperava engine não-nil")
				}
				if len(engine.Routes()) == 0 {
					t.Error("esperava pelo menos uma rota registrada")
				}
			}
		})
	}
}
