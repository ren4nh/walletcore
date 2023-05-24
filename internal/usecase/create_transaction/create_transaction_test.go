package createtransaction

import (
	"testing"

	"github.com.br/renanhartwig/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Renan Hartwig", "r@r")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Renan dos Santos", "r@r")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", account1.ID).Return(account1, nil)
	accountMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionMock := &TransactionGatewayMock{}
	transactionMock.On("Create", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(transactionMock, accountMock)

	output, err := uc.Execute(CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	transactionMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionMock.AssertNumberOfCalls(t, "Create", 1)
}
