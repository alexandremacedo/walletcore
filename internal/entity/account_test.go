package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Test", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, account.Client.ID, client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreateCreditAccount(t *testing.T) {
	client, _ := NewClient("Test", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, account.Balance, float64(100))
}

func TestCreateDebitAccount(t *testing.T) {
	client, _ := NewClient("Test", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, account.Balance, float64(50))
}
