package cockroach

const (
	selectAccountBalanceQuery = `SELECT  account_id, balance, user_id, version 
	FROM account_balances WHERE account_id = $1 AND user_id = $2;`
	updateAccountBalanceQuery = `UPDATE account_balances
	SET balance = $1, version = version + 1 WHERE account_id = $2 AND user_id = $3 AND version = $4
	RETURNING balance, updated_at;`
)
