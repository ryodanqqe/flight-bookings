package repository

import "database/sql"

func (r *UserPostgres) BeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *UserPostgres) CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}

func (r *UserPostgres) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
