package cockroach

type accountBalance struct {
	AccountID int32   `db:"account_id"`
	Balance   float32 `db:"balance"`
	UserID    int32   `db:"user_id"`
	Version   int32   `db:"version"`
}
