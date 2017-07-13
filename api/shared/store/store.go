package store

type QueryOptions struct {
	Offset *int
	Limit  *int
}

// Store is a data store interface.
type Store interface {
	Copy() Store
	Close()

	Accounts() AccountsStore
	Posts() PostsStore
}

// type StatusStore interface{
// 	Get()
// }

// AccountsStore is a accounts data store interface.
type AccountsStore interface {
	Count() (int, error)
	AdminCount() (int, error)
	Find(id string) (*Account, error)
	FindByCredentials(username, password string) (*Account, error)
	UsernameIsExist(username string) (bool, error)
	Create(account *Account) error
	Update(id string, updates Account) error
	Delete(id string) error
}

// PostsStore is a posts data store interface.
type PostsStore interface {
	Count() (int, error)
	PublishedCount() (int, error)
	Get(opts *QueryOptions) ([]*Post, error)
	GetPublished(opts *QueryOptions) ([]*Post, error)
	GetByFuzzyTitleOrFuzzyContent(keyword string, opts *QueryOptions) ([]*Post, error)
	GetPublishedByFuzzyTitleOrFuzzyContent(keyword string, opts *QueryOptions) ([]*Post, error)
	Find(id string) (*Post, error)
	FindPublished(id string) (*Post, error)
	FindByAuthorId(authorId, id string) (*Post, error)
	CreateByAuthorId(authorId string, account *Post) error
	UpdateByAuthorId(authorId, id string, updates Post) error
	Delete(id string) error
	DeleteByAuthorId(authorId, id string) error
}
