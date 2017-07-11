package mysql

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

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
	if s.store.database == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var count int

	err = s.table.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// AdminCount returns number of admin documnets of this collection.
func (s *accountsStore) AdminCount() (int, error) {
	if s.store.database == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var count int

	err = s.table.Where("`role` = ?", 9999).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Find returns a Account by given id.
func (s *accountsStore) Find(id string) (*store.Account, error) {
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var account store.Account

	err = s.table.New().Where("`id` = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// FindByCredentials returns a Account by given id.
func (s *accountsStore) FindByCredentials(username, password string) (*store.Account, error) {
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var account store.Account

	err = s.table.New().Where("`username` = ?", username).First(&account).Error
	if err != nil {
		return nil, err
	}
	if account.Id == nil {
		return nil, fmt.Errorf("Cannot find account by username %s.", username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*account.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// UsernameIsExist checks whether username is exist.
func (s *accountsStore) UsernameIsExist(username string) (bool, error) {
	if s.store.database == nil {
		return false, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var count int

	err = s.table.Where("`username` = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	}

	return false, nil
}

// Create inserts a new account into database.
func (s *accountsStore) Create(account *store.Account) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	newId := bson.NewObjectId().Hex()
	account.Id = &newId

	if account.Password == nil {
		account.Password = new(string)
	}
	h, err := bcrypt.GenerateFromPassword([]byte(*account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPassword := string(h)

	account.Password = &hashedPassword

	now := time.Now().Unix()
	account.CreatedAt = &now

	err = s.table.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

// Update updates a account by its id and a store.Account model.
func (s *accountsStore) Update(id string, account store.Account) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	account.Id = nil
	account.Username = nil
	account.Role = nil
	account.CreatedAt = nil

	err = s.table.Where("`id` = ?", id).Updates(account).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a account by its id.
func (s *accountsStore) Delete(id string) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	err = s.table.Where("`id` = ?", id).Delete(store.Account{}).Error
	if err != nil {
		return err
	}

	return nil
}
