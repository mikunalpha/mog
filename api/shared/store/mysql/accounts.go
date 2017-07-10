package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/mikunalpha/mog/api/shared/store"
)

// accountsStore is a mongo implementation of AccountsStore.
type accountsStore struct {
	store *Store
	table *gorm.DB
}

// Count returns number of documnets of this collection.
func (s *accountsStore) Count() (int, error) {
	return 0, nil
}

// AdminCount returns number of admin documnets of this collection.
func (s *accountsStore) AdminCount() (int, error) {
	return 0, nil
}

// Find returns a Account by given id.
func (s *accountsStore) Find(id string) (*store.Account, error) {
	var err error
	var account store.Account

	return &account, err
}

// FindByCredentials returns a Account by given id.
func (s *accountsStore) FindByCredentials(username, password string) (*store.Account, error) {
	var err error
	var account store.Account

	return &account, err
}

// UsernameIsExist checks whether username is exist.
func (s *accountsStore) UsernameIsExist(username string) (bool, error) {
	var err error

	return false, err
}

// Create inserts a new account into database.
func (s *accountsStore) Create(account *store.Account) error {
	var err error

	return err
}

// Update updates a account by its id and a store.Account model.
func (s *accountsStore) Update(id string, account store.Account) error {
	var err error

	return err
}

// Delete removes a account by its id.
func (s *accountsStore) Delete(id string) error {
	var err error

	return err
}
