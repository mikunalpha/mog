package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mikunalpha/mog/api/shared/store"
)

// postsStore is a mongo implementation of PostsStore.
type postsStore struct {
	store *Store
	table *gorm.DB
}

// Count returns number of documnets of this collection.
func (s *postsStore) Count() (int, error) {
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

// PublishedCount returns number of published documnets of this collection.
func (s *postsStore) PublishedCount() (int, error) {
	if s.store.database == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var count int

	err = s.table.Where("status = ?", 9999).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get returns Posts by given opts.
func (s *postsStore) Get(opts *store.QueryOptions) ([]*store.Post, error) {
	var err error
	var posts []*store.Post

	query := s.table.New().Order("`id` desc")

	if opts != nil {
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)

			if opts.Offset != nil {
				query = query.Offset(*opts.Offset)
			}
		}
	}

	err = query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*store.Post{}
	}

	return posts, nil
}

// GetPublished returns published Posts by given opts.
func (s *postsStore) GetPublished(opts *store.QueryOptions) ([]*store.Post, error) {
	var err error
	var posts []*store.Post

	query := s.table.New().Where("`published` = ?", true).Order("`id` desc")

	if opts != nil {
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)

			if opts.Offset != nil {
				query = query.Offset(*opts.Offset)
			}
		}
	}

	err = query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*store.Post{}
	}

	return posts, nil
}

// Find returns a Post by given id.
func (s *postsStore) Find(id string) (*store.Post, error) {
	var err error
	var post store.Post

	return &post, err
}

// FindPublished returns a published Post by given id.
func (s *postsStore) FindPublished(id string) (*store.Post, error) {
	var err error
	var post store.Post

	return &post, err
}

// FindByAuthorId returns a Post by given authorId and id.
func (s *postsStore) FindByAuthorId(authorId, id string) (*store.Post, error) {
	var err error
	var post store.Post

	return &post, err
}

// CreateByAuthorId inserts a new post into database.
func (s *postsStore) CreateByAuthorId(authorId string, post *store.Post) error {
	var err error

	return err
}

// UpdateByAuthorId updates a post by its id and a store.Post model.
func (s *postsStore) UpdateByAuthorId(authorId, id string, post store.Post) error {
	var err error

	return err
}

// DeleteByAuthorId removes a post by its id.
func (s *postsStore) DeleteByAuthorId(authorId, id string) error {
	var err error

	return err
}
