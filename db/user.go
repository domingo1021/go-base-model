package db

import (
	"context"
	"database/sql"
	"log"
)

// User represents the user model in our application.
type User struct {
	BaseModel
	ID    int
	Name  string
	Email string
}

// CreateUser is a method on User to insert a new user into the database.
func (u *User) CreateUser(ctx context.Context, user User) error {
	// Implement the SQL insert logic using u.DB
	// Example: _, err := u.DB.ExecContext(ctx, "INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	// return err

	// Placeholder implementation
	return nil
}

// CreateUserAndLogOperation is an operation that creates a user and logs the creation.
type CreateUserAndLogOperation struct {
	User User
	Log  string
}

// ExecuteInTransaction executes the user creation and logging in a transaction.
func (op CreateUserAndLogOperation) ExecuteInTransaction(ctx context.Context, tx *sql.Tx) error {
	// Insert the user into the database
	userInsertQuery := "INSERT INTO users (name, email) VALUES ($1, $2)"
	log.Printf("Executing query: %s with Name: %s, Email: %s", userInsertQuery, op.User.Name, op.User.Email)
	_, err := tx.ExecContext(ctx, userInsertQuery, op.User.Name, op.User.Email)
	if err != nil {
		return err
	}

	// Log the user creation event
	logInsertQuery := "INSERT INTO logs (event) VALUES ($1)"
	log.Printf("Executing query: %s with Event: %s", logInsertQuery, op.Log)
	_, err = tx.ExecContext(ctx, logInsertQuery, op.Log)
	if err != nil {
		return err
	}

	return nil
}

// TransactionalCreateUserAndLog creates a user and logs the creation in a transaction.
func (u *User) TransactionalCreateUserAndLog(ctx context.Context, user User, logMessage string) error {
	createUserAndLogOp := CreateUserAndLogOperation{User: user, Log: logMessage}
	return u.ExecTx(ctx, nil, createUserAndLogOp)
}
