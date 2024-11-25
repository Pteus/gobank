package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgesStore struct {
	db *sql.DB
}

func NewPostgesStore() (*PostgesStore, error) {
	connStr := "user=postgres dbname=gobank password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgesStore{db: db}, nil
}

func (p *PostgesStore) Init() error {
	return p.createAccountTable()
}

func (p *PostgesStore) createAccountTable() error {
	sqlStatement := `create table if not exists account(
				id serial primary key,
				first_name varchar(50),
				last_name varchar(50),
				number serial,
				balance numeric,
				created_at timestamp
			)`

	_, err := p.db.Exec(sqlStatement)
	return err
}

func (p *PostgesStore) CreateAccount(a *Account) error {
	sqlStatement := `INSERT INTO account (first_name, last_name, number, balance, created_at)
			VALUES ($1, $2, $3, $4, $5)`
	_, err := p.db.Exec(sqlStatement, a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgesStore) DeleteAccount(id int) error {
	return nil
}

func (p *PostgesStore) UpdateAccount(a *Account) error {
	return nil
}

func (p *PostgesStore) GetAccounts() ([]*Account, error) {
	rows, err := p.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (p *PostgesStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
