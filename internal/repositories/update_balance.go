package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"pixel-pay/internal/types"
)

const (
	updateBalanceQuery = "SELECT * FROM pixel_pay.transfer($1, $2, $3)"
)

func UpdateBalance(transaction types.TransactionRequest, tx pgx.Tx) error {
	var success bool
	var message string

	err := tx.QueryRow(context.Background(), updateBalanceQuery, transaction.Payer, transaction.Payee, transaction.Value).Scan(&success, &message)

	log.Println("UpdateBalance: ", success, message)
	if err != nil || !success {
		return fmt.Errorf("error updating balance")
	}

	return nil

}
