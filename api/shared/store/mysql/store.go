package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mikunalpha/mog/api/shared/store"
)

// Store is a mysql implementation of store.
type Store struct {
	database *gorm.DB

	accountsStore *accountsStore
	postsStore    *postsStore
}

// Copy returns a new store with copied gorm database.
func (s *Store) Copy() store.Store {
	database := s.database.New()

	cs := Store{
		database: database,
	}

	cs.accountsStore = &accountsStore{
		store: &cs,
		table: database.Model(&store.Account{}),
	}

	cs.postsStore = &postsStore{
		store: &cs,
		table: database.Model(&store.Post{}),
	}

	return &cs
}

// Close closes the gorm database of this store.
func (s *Store) Close() {
	s.database.Close()
}

// Accounts returns a account store.
func (s *Store) Accounts() store.AccountsStore {
	return s.accountsStore
}

// Posts returns a account store.
func (s *Store) Posts() store.PostsStore {
	return s.postsStore
}
