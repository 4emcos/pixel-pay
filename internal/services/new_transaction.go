package services

import (
	"context"
	"fmt"
	"log"
	"pixel-pay/database"
	"pixel-pay/internal/repositories"
	"pixel-pay/internal/rest"
	"pixel-pay/internal/types"
)

func NewTransaction(db database.Pgx, transaction types.TransactionRequest) error {
	user, err := repositories.GetUserById(transaction.Payer, db)
	if err != nil {
		return fmt.Errorf("error getting user by id")
	}

	if user.DocumentType == "cnpj" {
		return fmt.Errorf("transaction not allowed for user with document type cnpj")
	}

	if user.Balance < transaction.Value {
		return fmt.Errorf("insufficient balance for transaction")
	}

	log.Println("User: ", user)

	tx, err := db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error beginning transaction")
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	err = repositories.UpdateBalance(transaction, tx)
	if err != nil {
		return fmt.Errorf("error updating balance")
	}

	err = rest.Authorizer()
	if err != nil {
		return fmt.Errorf("authorization failed")
	}

	err = rest.Notification()
	if err != nil {
		return fmt.Errorf("notification failed")
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error committing transaction")
	}

	return nil
}
