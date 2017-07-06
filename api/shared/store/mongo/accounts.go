package mongo

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/mikunalpha/mog/api/shared/store"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// accountsStore is a mongo implementation of AccountsStore.
type accountsStore struct {
	store      *Store
	collection *mgo.Collection
}

// Count returns number of documnets of this collection.
func (s *accountsStore) Count() (int, error) {
	if s.store.session == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error

	count, err := countRetry(DefaultRetryTimes, s.collection, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// AdminCount returns number of admin documnets of this collection.
func (s *accountsStore) AdminCount() (int, error) {
	if s.store.session == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error

	count, err := countRetry(DefaultRetryTimes, s.collection, bson.M{"role": 9999})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Find returns a Account by given id.
func (s *accountsStore) Find(id string) (*store.Account, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var account store.Account

	selector := bson.M{
		"_id": id,
	}

	err = findRetry(DefaultRetryTimes, s.collection, s.collection.Find(selector), &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// FindByCredentials returns a Account by given id.
func (s *accountsStore) FindByCredentials(username, password string) (*store.Account, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var account store.Account

	selector := bson.M{
		"username": username,
	}

	err = findRetry(DefaultRetryTimes, s.collection, s.collection.Find(selector), &account)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*account.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// UsernameIsExist checks whether username is exist.
func (s *accountsStore) UsernameIsExist(username string) (bool, error) {
	if s.store.session == nil {
		return false, fmt.Errorf("Not work until database is setup.")
	}

	var err error

	count, err := countRetry(DefaultRetryTimes, s.collection, bson.M{"username": username})
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
	if s.store.session == nil {
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

	err = insertRetry(DefaultRetryTimes, s.collection, account)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a account by its id and a store.Account model.
func (s *accountsStore) Update(id string, account store.Account) error {
	if s.store.session == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	account.Id = nil
	account.Username = nil
	account.Role = nil
	account.CreatedAt = nil

	err = updateRetry(DefaultRetryTimes, s.collection, bson.M{"_id": id}, bson.M{"$set": account})
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a account by its id.
func (s *accountsStore) Delete(id string) error {
	if s.store.session == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	err = removeRetry(DefaultRetryTimes, s.collection, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
