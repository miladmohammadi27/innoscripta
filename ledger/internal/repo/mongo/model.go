package mongo

type transaction struct {
	AccountID       int32  `bson:"account_id" json:"account_id"`
	UserID          int32  `bson:"user_id" json:"user_id"`
	Amount          int32  `bson:"amount" json:"amount"`
	TransactionType string `bson:"transaction_type" json:"transaction_type"`
	Status          string `bson:"status" json:"status"`
	CreatedAt       string `bson:"created_at" json:"created_at"`
}
