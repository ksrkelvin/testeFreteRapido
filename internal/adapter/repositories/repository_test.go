package repositories_test

import (
	"testeFreteRapido/internal/adapter/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewRepository(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro criando sqlmock: %s", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("erro criando gorm com sqlmock: %s", err)
	}
	repo := repositories.NewRepository(gormDB)

	tests := []struct {
		name string
		db   *gorm.DB
		want *repositories.Repository
	}{
		{
			name: "basic repository creation",
			db:   gormDB,
			want: &repositories.Repository{
				Quote:   repo.Quote,
				Metrics: repo.Metrics,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repositories.NewRepository(tt.db)
			if got.Metrics == nil || got.Quote == nil {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
