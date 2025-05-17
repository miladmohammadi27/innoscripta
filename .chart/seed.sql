-- Create sequences
CREATE SEQUENCE user_id_seq START 1;
CREATE SEQUENCE account_id_seq START 10001;

-- Users table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY DEFAULT nextval('user_id_seq'),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Accounts table
CREATE TABLE accounts (
    account_id INTEGER PRIMARY KEY DEFAULT nextval('account_id_seq'),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Account balances table
CREATE TABLE account_balances (
    account_id INTEGER PRIMARY KEY REFERENCES accounts(account_id),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    balance NUMERIC(20,2) NOT NULL DEFAULT 0,
    version INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Grant permissions to banking user
GRANT ALL ON DATABASE banking TO banking;
GRANT ALL ON SCHEMA public TO banking;
GRANT ALL ON ALL TABLES IN SCHEMA public TO banking;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO banking;

-- CREATE TABLE ledger_entries (
--     id SERIAL PRIMARY KEY,
--     account_id INTEGER NOT NULL REFERENCES accounts(id),
--     amount NUMERIC(20, 2) NOT NULL,
--     description TEXT,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );