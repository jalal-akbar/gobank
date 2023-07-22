package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

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

func utcToGMT8() time.Time {
	utcTime := time.Now().UTC()

	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		panic(err)
	}
	bimaTime := utcTime.In(loc)
	return bimaTime
}
