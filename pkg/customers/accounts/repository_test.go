// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package accounts

import (
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/customers/pkg/admin"
	"github.com/moov-io/customers/pkg/client"
	"github.com/moov-io/customers/pkg/secrets"
	"github.com/moov-io/identity/pkg/database"
)

func setupTestAccountRepository(t *testing.T) *sqlAccountRepository {
	db, close, err := database.NewAndMigrate(database.InMemorySqliteConfig, nil, nil)
	t.Cleanup(close)
	if err != nil {
		t.Fatal(err)
	}
	return NewRepo(db)
}

func TestRepository(t *testing.T) {
	customerID, userID := base.ID(), base.ID()
	repo := setupTestAccountRepository(t)

	// initial read, find no accounts
	accounts, err := repo.getCustomerAccounts(customerID)
	if len(accounts) != 0 || err != nil {
		t.Fatalf("got accounts=%#v error=%v", accounts, err)
	}

	// create account
	acct, err := repo.createCustomerAccount(customerID, userID, &createAccountRequest{
		AccountNumber: "123",
		RoutingNumber: "987654320",
		Type:          client.CHECKING,
	})
	if err != nil {
		t.Fatal(err)
	}

	// read after creating
	accounts, err = repo.getCustomerAccounts(customerID)
	if len(accounts) != 1 || err != nil {
		t.Fatalf("got accounts=%#v error=%v", accounts, err)
	}
	if accounts[0].AccountID != acct.AccountID {
		t.Errorf("accounts[0].AccountID=%s acct.AccountID=%s", accounts[0].AccountID, acct.AccountID)
	}

	// delete, expect no accounts
	if err := repo.deactivateCustomerAccount(acct.AccountID); err != nil {
		t.Fatal(err)
	}
	accounts, err = repo.getCustomerAccounts(customerID)
	if len(accounts) != 0 || err != nil {
		t.Fatalf("got accounts=%#v error=%v", accounts, err)
	}
}

func TestRepository__getEncryptedAccountNumber(t *testing.T) {
	customerID, userID := base.ID(), base.ID()
	repo := setupTestAccountRepository(t)

	keeper := secrets.TestStringKeeper(t)

	// create account
	req := &createAccountRequest{
		AccountNumber: "123",
		RoutingNumber: "987654320",
		Type:          client.CHECKING,
	}
	if err := req.disfigure(keeper); err != nil {
		t.Fatal(err)
	}
	acct, err := repo.createCustomerAccount(customerID, userID, req)
	if err != nil {
		t.Fatal(err)
	}

	// read encrypted account number
	encrypted, err := repo.getEncryptedAccountNumber(customerID, acct.AccountID)
	if err != nil {
		t.Fatal(err)
	}
	if encrypted == "" {
		t.Error("missing encrypted account number")
	}
}

func TestRepository__updateAccountStatus(t *testing.T) {
	customerID, userID := base.ID(), base.ID()
	repo := setupTestAccountRepository(t)

	keeper := secrets.TestStringKeeper(t)

	// create account
	req := &createAccountRequest{
		AccountNumber: "123",
		RoutingNumber: "987654320",
		Type:          client.CHECKING,
	}
	if err := req.disfigure(keeper); err != nil {
		t.Fatal(err)
	}
	acct, err := repo.createCustomerAccount(customerID, userID, req)
	if err != nil {
		t.Fatal(err)
	}

	// update status
	if err := repo.updateAccountStatus(acct.AccountID, admin.VALIDATED); err != nil {
		t.Fatal(err)
	}

	// check status after update
	acct, err = repo.getCustomerAccount(customerID, acct.AccountID)
	if err != nil {
		t.Fatal(err)
	}
	if acct.Status != client.VALIDATED {
		t.Errorf("unexpected status: %s", acct.Status)
	}
}

func TestRepositoryUnique(t *testing.T) {
	keeper := secrets.TestStringKeeper(t)

	check := func(t *testing.T, repo *sqlAccountRepository) {
		customerID, userID := base.ID(), base.ID()
		req := &createAccountRequest{
			AccountNumber: "156421",
			RoutingNumber: "123456780",
			Type:          client.SAVINGS,
		}
		if err := req.disfigure(keeper); err != nil {
			t.Fatal(err)
		}

		// first write should pass
		if _, err := repo.createCustomerAccount(customerID, userID, req); err != nil {
			t.Fatal(err)
		}
		// second write should fail
		if _, err := repo.createCustomerAccount(customerID, userID, req); err != nil {
			if !database.UniqueViolation(err) {
				t.Fatalf("unexpected error: %v", err)
			}
		}
	}

	// SQLite tests
	db, close, err := database.NewAndMigrate(database.InMemorySqliteConfig, nil, nil)
	t.Cleanup(close)
	if err != nil {
		t.Error(err)
	}
	check(t, NewRepo(db))

	// MySQL tests
	mysqlDB := database.CreateTestMySQLDB(t)
	defer mysqlDB.Close()
	check(t, NewRepo(mysqlDB.DB))
}