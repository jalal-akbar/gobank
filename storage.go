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
	GetAccountByNumber(number int) (*Account, error)
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
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account
	(first_name, last_name, number,encrypted_password, balance, created_at) 
	VALUES
	($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Query(query,
		acc.Firstname,
		acc.Lastname,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(acc *Account) error {
	query := "UPDATE account SET balance=$2 WHERE id=$1"
	rows, err := s.db.Query(query, acc.Id, acc.Balance)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rows)

	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := "DELETE FROM account WHERE id = $1"
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccountByNumber(number int) (a *Account, err error) {
	query := "SELECT * FROM account WHERE number = $1"
	rows, err := s.db.Query(query, number)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoRows(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", number)
}

func (s *PostgresStore) GetAccountByID(id int) (a *Account, err error) {
	query := "SELECT * FROM account WHERE id = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoRows(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	accounts := []*Account{}

	rows, err := s.db.Query(`SELECT * FROM account`)
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	for rows.Next() {

		acc, err := scanIntoRows(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func scanIntoRows(rows *sql.Rows) (*Account, error) {
	acc := new(Account)
	if err := rows.Scan(&acc.Id, &acc.Firstname, &acc.Lastname, &acc.Number, &acc.EncryptedPassword, &acc.Balance, &acc.CreatedAt); err != nil {
		return nil, err
	}
	return acc, nil
}
