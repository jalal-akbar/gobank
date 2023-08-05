package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type Account struct {
	Id                int       `json:"id"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func NewAccount(firstname, lastname, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		Firstname:         firstname,
		Lastname:          lastname,
		EncryptedPassword: string(encpw),
		Number:            int64(rand.Intn(100000)),
		CreatedAt:         utcToGMT8(),
	}, nil
}

func (a *Account) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type UpdateAccountRequest struct {
	Balance int64 `json:"balance"`
}

type TrasferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

func UpdateAccount(id int, balance int64) *Account {
	return &Account{
		Id:      id,
		Balance: balance,
	}
}

func utcToGMT8() time.Time {
	utcTime := time.Now().UTC()

	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		panic(err)
	}
	bimaTime := utcTime.In(loc)
	return bimaTime
}
