package cockroach

const (
	createUserQuery = `INSERT INTO users (name, email)
					VALUES ($1, $2)
					RETURNING user_id;`

	createAccountQuery = `INSERT INTO accounts (user_id)
						VALUES ($1)
						RETURNING account_id;`
	createZeroAccountBalanceQuery = `INSERT INTO account_balances (account_id, user_id) 
						VALUES ($1, $2);`
	checkUserExistsQuery = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`
)
