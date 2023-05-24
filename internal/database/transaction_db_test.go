package database

import (
	"database/sql"
	"testing"

	"github.com.br/renanhartwig/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	accountDB     *AccountDB
	clientDB      *ClientDB
	transactionDB *TransactionDB
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance number, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from varchar(255), account_to varchar(255), amount number, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.clientDB = NewClientDB(db)
	s.transactionDB = NewTransactionDB(db)
	client1, _ := entity.NewClient("Renan 1", "r@r.com")
	client2, _ := entity.NewClient("Renan 2", "r@r.com")
	s.clientDB.Save(client1)
	s.clientDB.Save(client2)
	accountFrom := entity.NewAccount(client1)
	accountFrom.Balance = 1000
	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	s.accountDB.Save(accountFrom)
	s.accountDB.Save(accountTo)
	s.accountFrom = accountFrom
	s.accountTo = accountTo
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100.50)
	s.Nil(err)
	err = s.transactionDB.Save(transaction)
	s.Nil(err)

	s.Equal(transaction.AccountFrom.Balance, 899.50)
	s.Equal(transaction.AccountTo.Balance, 1100.50)
}
