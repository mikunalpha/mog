package mysql

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

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

	err = s.table.Where("`status` = ?", 9999).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get returns Posts by given opts.
func (s *postsStore) Get(opts *store.QueryOptions) ([]*store.Post, error) {
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	query := s.table.New().Order("`id` desc")

	if opts != nil {
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)

			if opts.Offset != nil {
				query = query.Offset(*opts.Offset - 1)
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
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	query := s.table.New().Where("`published` = ?", true).Order("`id` desc")

	if opts != nil {
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)

			if opts.Offset != nil {
				query = query.Offset(*opts.Offset - 1)
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
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	err = s.table.New().Where("`id` = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, err
}

// FindPublished returns a published Post by given id.
func (s *postsStore) FindPublished(id string) (*store.Post, error) {
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	err = s.table.New().Where("`id` = ? AND `published` = ?", id, true).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, err
}

// FindByAuthorId returns a Post by given authorId and id.
func (s *postsStore) FindByAuthorId(authorId, id string) (*store.Post, error) {
	if s.store.database == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	err = s.table.New().Where("`id` = ? AND `author_id` = ?", id, authorId).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, err
}

// CreateByAuthorId inserts a new post into database.
func (s *postsStore) CreateByAuthorId(authorId string, post *store.Post) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	newId := bson.NewObjectId().Hex()
	post.Id = &newId

	if post.Title == nil {
		title := "Unknown Post Title"
		post.Title = &title
	} else if *post.Title == "" {
		*post.Title = "Unknown Post Title"
	}

	if post.Published == nil {
		post.Published = new(bool)
	}

	now := time.Now().Unix()
	post.CreatedAt = &now

	post.AuthorId = &authorId

	err = s.table.Create(post).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateByAuthorId updates a post by its id and a store.Post model.
func (s *postsStore) UpdateByAuthorId(authorId, id string, post store.Post) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	post.Id = nil
	post.CreatedAt = nil
	post.AuthorId = nil

	now := time.Now().Unix()
	post.UpdatedAt = &now

	err = s.table.Where("`id` = ? AND `author_id` = ?", id, authorId).Updates(post).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by its id.
func (s *postsStore) Delete(id string) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	err = s.table.Where("`id` = ?", id).Delete(store.Post{}).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteByAuthorId removes a post by given authorId and its id.
func (s *postsStore) DeleteByAuthorId(authorId, id string) error {
	if s.store.database == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	err = s.table.Where("`id` = ? AND `author_id` = ?", id, authorId).Delete(store.Post{}).Error
	if err != nil {
		return err
	}

	return nil
}
