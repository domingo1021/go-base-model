// db.go
package db

import (
	"context"
	"database/sql"
	"fmt"
)

type BaseModel struct {
	*Queries
	DB *sql.DB
}

// TransactionHandler defines behavior for starting and handling transactions.
type TransactionHandler interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// TransactionalOperation defines the behavior of an operation that should be executed in a transaction.
type TransactionalOperation interface {
	ExecuteInTransaction(ctx context.Context, tx *sql.Tx) error
}

// BeginTx starts a database transaction using the BaseModel's database connection.
func (m *BaseModel) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return m.DB.BeginTx(ctx, opts)
}

type ExecTxFunc func(ctx context.Context, tx *sql.Tx) error

// ExecTx executes a TransactionalOperation within a database transaction.
func (m *BaseModel) ExecTx(ctx context.Context, opts *sql.TxOptions, op TransactionalOperation) error {
	tx, err := m.DB.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	if err := op.ExecuteInTransaction(ctx, tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
