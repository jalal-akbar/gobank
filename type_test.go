package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	assert := assert.New(t)
	acc, err := NewAccount("a", "b", "jalal1417")
	assert.Nil(err)

	fmt.Printf("%+v\n", acc)

}
