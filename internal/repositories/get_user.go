package repositories

import (
	"context"
	"fmt"
	"pixel-pay/database"
	"pixel-pay/internal/domain"
)

const (
	getUserByIdQuery = "SELECT * from pixel_pay.users WHERE id = $1"
)

func GetUserById(id int32, db database.Pgx) (domain.User, error) {
	var user domain.User

	err := db.QueryRow(context.Background(), getUserByIdQuery, id).Scan(&user.Id, &user.Name, &user.Document, &user.DocumentType, &user.Email, &user.Balance, nil)
	if err != nil {
		return domain.User{}, fmt.Errorf("err geting user")
	}

	return user, nil

}
