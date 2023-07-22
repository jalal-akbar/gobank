package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(id int) error
	GetAccountByID(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStrore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account
	(first_name, last_name, number, balance, created_at) 
	VALUES
	($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(query,
		acc.Firstname,
		acc.Lastname,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (a *Account, err error) {
	return nil, err
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	accounts := []*Account{}

	rows, err := s.db.Query(`SELECT * FROM account`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		acc := new(Account)
		err := rows.Scan(&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Number, &acc.Balance, &acc.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return accounts, nil
}
