package mongo

import (
	"fmt"
	"regexp"
	"time"

	"github.com/mikunalpha/mog/api/shared/store"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// postsStore is a mongo implementation of PostsStore.
type postsStore struct {
	store      *Store
	collection *mgo.Collection
}

// Count returns number of documnets of this collection.
func (s *postsStore) Count() (int, error) {
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

// PublishedCount returns number of published documnets of this collection.
func (s *postsStore) PublishedCount() (int, error) {
	if s.store.session == nil {
		return 0, fmt.Errorf("Not work until database is setup.")
	}

	var err error

	count, err := countRetry(DefaultRetryTimes, s.collection, bson.M{"status": 9999})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get returns Posts by given opts.
func (s *postsStore) Get(opts *store.QueryOptions) ([]*store.Post, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	var query *mgo.Query

	query = s.collection.Find(nil).Sort("-_id")
	// Find(nil).
	// Select(bson.M{
	// 	"content": 0})

	if opts != nil {
		if opts.Offset != nil {
			query = query.Skip(*opts.Offset)
		}
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)
		}
	}

	err = getRetry(DefaultRetryTimes, s.collection, query, &posts)
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
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	var query *mgo.Query

	query = s.collection.Find(bson.M{"published": true}).Sort("-_id")
	// Find(bson.M{
	// 	"published": true}).
	// Select(bson.M{
	// 	"content": 0})

	if opts != nil {
		if opts.Offset != nil {
			query = query.Skip(*opts.Offset)
		}
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)
		}
	}

	err = getRetry(DefaultRetryTimes, s.collection, query, &posts)
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*store.Post{}
	}

	return posts, nil
}

// GetByFuzzyTitleOrFuzzyContent returns Posts by given keyword and opts.
func (s *postsStore) GetByFuzzyTitleOrFuzzyContent(keyword string, opts *store.QueryOptions) ([]*store.Post, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	keyword = regexp.QuoteMeta(keyword)

	var query *mgo.Query

	query = s.collection.Find(bson.M{
		"$or": []bson.M{
			{
				"title": bson.RegEx{Pattern: keyword, Options: "i"},
			},
			{
				"content": bson.RegEx{Pattern: keyword, Options: "i"},
			},
		}}).
		Sort("-_id")

	if opts != nil {
		if opts.Offset != nil {
			query = query.Skip(*opts.Offset)
		}
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)
		}
	}

	err = getRetry(DefaultRetryTimes, s.collection, query, &posts)
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*store.Post{}
	}

	return posts, nil
}

// GetPublishedByFuzzyTitleOrFuzzyContent returns Posts by given keyowrd and opts.
func (s *postsStore) GetPublishedByFuzzyTitleOrFuzzyContent(keyword string, opts *store.QueryOptions) ([]*store.Post, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var posts []*store.Post

	keyword = regexp.QuoteMeta(keyword)

	var query *mgo.Query

	query = s.collection.Find(bson.M{
		"published": true,
		"$or": []bson.M{
			{
				"title": bson.RegEx{Pattern: keyword, Options: "i"},
			},
			{
				"content": bson.RegEx{Pattern: keyword, Options: "i"},
			},
		}}).
		Sort("-_id")

	if opts != nil {
		if opts.Offset != nil {
			query = query.Skip(*opts.Offset)
		}
		if opts.Limit != nil {
			query = query.Limit(*opts.Limit)
		}
	}

	err = getRetry(DefaultRetryTimes, s.collection, query, &posts)
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
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	selector := bson.M{
		"_id": id,
	}

	err = findRetry(DefaultRetryTimes, s.collection, s.collection.Find(selector), &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// FindPublished returns a published Post by given id.
func (s *postsStore) FindPublished(id string) (*store.Post, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	selector := bson.M{
		"_id":       id,
		"published": true,
	}

	err = findRetry(DefaultRetryTimes, s.collection, s.collection.Find(selector), &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// FindByAuthorId returns a Post by given authorId and id.
func (s *postsStore) FindByAuthorId(authorId, id string) (*store.Post, error) {
	if s.store.session == nil {
		return nil, fmt.Errorf("Not work until database is setup.")
	}

	var err error
	var post store.Post

	selector := bson.M{
		"author_id": authorId,
		"_id":       id,
	}

	err = findRetry(DefaultRetryTimes, s.collection, s.collection.Find(selector), &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// CreateByAuthorId inserts a new post into database.
func (s *postsStore) CreateByAuthorId(authorId string, post *store.Post) error {
	if s.store.session == nil {
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

	err = insertRetry(DefaultRetryTimes, s.collection, post)
	if err != nil {
		return err
	}

	return nil
}

// UpdateByAuthorId updates a post by its id and a store.Post model.
func (s *postsStore) UpdateByAuthorId(authorId, id string, post store.Post) error {
	if s.store.session == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	post.Id = nil
	post.CreatedAt = nil
	post.AuthorId = nil

	now := time.Now().Unix()
	post.UpdatedAt = &now

	err = updateRetry(DefaultRetryTimes, s.collection, bson.M{"author_id": authorId, "_id": id}, post)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by its id.
func (s *postsStore) Delete(id string) error {
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

// DeleteByAuthorId removes a post by given authorId and its id.
func (s *postsStore) DeleteByAuthorId(authorId, id string) error {
	if s.store.session == nil {
		return fmt.Errorf("Not work until database is setup.")
	}

	var err error

	err = removeRetry(DefaultRetryTimes, s.collection, bson.M{"author_id": authorId, "_id": id})
	if err != nil {
		return err
	}

	return nil
}
