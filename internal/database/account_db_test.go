package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com.br/renanhartwig/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	clientDB  *ClientDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance number, created_at date, updated_at date)")
	s.accountDB = NewAccountDB(db)
	s.clientDB = NewClientDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	client, _ := entity.NewClient("Renan", "renan@r.com")

	err := s.clientDB.Save(client)
	s.Nil(err)
	account := &entity.Account{
		ID:        "1",
		Balance:   100,
		Client:    client,
		CreatedAt: time.Now(),
	}

	err = s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Renan", "renan@r.com")
	s.clientDB.Save(client)

	account := entity.NewAccount(client)
	s.accountDB.Save(account)

	accountDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.NotNil(accountDB.CreatedAt)
	s.Equal(client.ID, accountDB.Client.ID)
	s.Equal(client.Name, accountDB.Client.Name)
	s.Equal(client.Email, accountDB.Client.Email)
	s.NotNil(accountDB.CreatedAt)
}
