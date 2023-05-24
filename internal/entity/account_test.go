package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "r@r")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "r@r")
	account := NewAccount(client)
	account.Credit(10.0)
	assert.Equal(t, account.Balance, 10.0)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "r@r")
	account := NewAccount(client)
	account.Credit(10.0)
	account.Debit(5.0)
	assert.Equal(t, account.Balance, 5.0)
}
