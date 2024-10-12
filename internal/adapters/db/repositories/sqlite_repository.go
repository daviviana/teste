package repositories

import (
	"log"
	"teste/internal/adapters/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository[T any] struct {
	*RepositoryGORM[T]
}

func NewSQLiteRepository[T any](databasePath string) *SQLiteRepository[T] {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados SQLite:", err)
	}

	db.AutoMigrate(&models.User{})

	return &SQLiteRepository[T]{
		RepositoryGORM: NewRepositoryGORM[T](db),
	}
}
