package cockroach

import (
	"context"
	"errors"
	"fmt"

	"backoffice/internal/helper/log"
	"backoffice/internal/manager/entity"
	"backoffice/internal/repo"

	"github.com/jackc/pgx/v5"
	"github.com/samber/do"
)

type userRepo struct {
	db *pgx.Conn
	lg log.Logger
}

func NewUserRepo(i *do.Injector) (repo.UserRepo, error) {
	return userRepo{
		db: do.MustInvoke[*pgx.Conn](i),
		lg: do.MustInvoke[log.Logger](i),
	}, nil
}

func (ur userRepo) CreateUser(ctx context.Context, user entity.User) (string, error) {
	var returnedID string
	err := ur.db.QueryRow(ctx, createUserQuery, user.Name, user.Email).Scan(&returnedID)
	if err != nil {
		return returnedID, errors.Join(errInsert, err)
	}
	return returnedID, nil
}

func (ur userRepo) CreateAccount(ctx context.Context, userID int) (int32, error) {
	var accountID int32

	// Check if user exists
	var exists bool
	err := ur.db.QueryRow(ctx, checkUserExistsQuery, userID).Scan(&exists)
	if err != nil {
		return accountID, fmt.Errorf("error checking user existence: %w", err)
	}

	if !exists {
		return accountID, fmt.Errorf("user with ID %d does not exist", userID)
	}

	// Start a transaction
	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return accountID, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create account
	err = tx.QueryRow(ctx, createAccountQuery, userID).Scan(&accountID)
	if err != nil {
		tx.Rollback(ctx)
		return accountID, errors.Join(errInsert, err)
	}

	// Create account balance
	_, err = tx.Exec(ctx, createZeroAccountBalanceQuery, accountID, userID)
	if err != nil {
		tx.Rollback(ctx)
		return accountID, fmt.Errorf("failed to create account balance: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return accountID, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return accountID, nil
}
