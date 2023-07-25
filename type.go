package main

import (
	"math/rand"
	"time"
)

type Account struct {
	Id        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: utcToGMT8(),
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
