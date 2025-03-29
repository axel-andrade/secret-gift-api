package postgres_repositories

import (
	postgres_database "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres"
	"gorm.io/gorm"
)

type BasePostgresRepository struct {
	Db *gorm.DB
	Tx *gorm.DB
}

func BuildBasePostgresRepository() *BasePostgresRepository {
	db := postgres_database.GetDB()
	return &BasePostgresRepository{Db: db, Tx: nil}
}

func (r *BasePostgresRepository) getQueryOrTx() *gorm.DB {
	if r.Tx != nil {
		return r.Tx
	}

	return r.Db
}

func (r *BasePostgresRepository) StartTransaction() error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := r.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	r.Tx = tx

	return nil
}

func (r *BasePostgresRepository) CommitTransaction() error {
	err := r.Tx.Commit().Error
	r.Tx = nil
	if err != nil {
		return err
	}
	return nil
}

func (r *BasePostgresRepository) CancelTransaction() error {
	err := r.Tx.Rollback().Error
	r.Tx = nil
	if err != nil {
		return err
	}
	return nil
}
