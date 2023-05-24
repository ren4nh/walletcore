package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Renan Hartwig", "renan@r.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Renan Hartwig", client.Name)
	assert.Equal(t, "renan@r.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "r@r.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "renan@r.com")
	err := client.Update("Renan dos Santos Hartwig", "renan@r.com")
	assert.Nil(t, err)
	assert.Equal(t, "Renan dos Santos Hartwig", client.Name)
	assert.Equal(t, "renan@r.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "renan@r.com")
	err := client.Update("", "renan@r.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Renan Hartwig", "renan@r.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
