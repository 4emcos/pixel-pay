package types

type TransactionRequest struct {
	Value float32 `json:"value" binding:"required,min=1"`
	Payer int32   `json:"payer" binding:"required,min=1"`
	Payee int32   `json:"payee" binding:"required,min=1"`
}
