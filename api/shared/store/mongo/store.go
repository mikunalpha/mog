package mongo

import (
	"github.com/mikunalpha/mog/api/shared/store"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DefaultRetryTimes uint32 = 2

// NewStore returns a new store with *mgo.DialInfo.
func NewStore(di *mgo.DialInfo) (store.Store, error) {
	var session *mgo.Session
	var err error

	for _ = range di.Addrs {
		session, err = mgo.DialWithInfo(di)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	s := Store{
		session:  session,
		database: session.DB(""),
	}

	return s.Copy(), nil
}

// Store is a mongo implementation of store.
type Store struct {
	session  *mgo.Session
	database *mgo.Database

	accountsStore *accountsStore
	postsStore    *postsStore
}

// Copy returns a new store with copied mgo session.
func (s *Store) Copy() store.Store {
	session := s.session.Copy()

	cs := Store{
		session:  session,
		database: session.DB(""),
	}

	cs.accountsStore = &accountsStore{
		store:      &cs,
		collection: cs.database.C((&store.Account{}).CollectionName()),
	}

	cs.postsStore = &postsStore{
		store:      &cs,
		collection: cs.database.C((&store.Post{}).CollectionName()),
	}

	return &cs
}

// Close closes the mgo session of this store.
func (s *Store) Close() {
	s.session.Close()
}

// Accounts returns a account store.
func (s *Store) Accounts() store.AccountsStore {
	return s.accountsStore
}

// Posts returns a account store.
func (s *Store) Posts() store.PostsStore {
	return s.postsStore
}

// getRetry retries to find t times unit error is not 'EOF' problem.
func getRetry(t uint32, c *mgo.Collection, q *mgo.Query, doc interface{}) error {
	var err error

	for {
		err = q.All(doc)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// findRetry retries to find t times unit error is not 'EOF' problem.
func findRetry(t uint32, c *mgo.Collection, q *mgo.Query, doc interface{}) error {
	var err error

	for {
		err = q.One(doc)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// insertRetry retries to insert t times unit error is not 'EOF' problem.
func insertRetry(t uint32, c *mgo.Collection, docs ...interface{}) error {
	var err error

	for {
		err = c.Insert(docs...)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// updateRetry retries to update or insert t times unit error is not 'EOF' problem.
func updateRetry(t uint32, c *mgo.Collection, selector, changes interface{}) error {
	var err error

	for {
		err = c.Update(selector, bson.M{"$set": changes})
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// updateAllRetry retries to update or insert t times unit error is not 'EOF' problem.
func updateAllRetry(t uint32, c *mgo.Collection, selector, changes interface{}) error {
	var err error

	for {
		_, err = c.UpdateAll(selector, bson.M{"$set": changes})
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// upsertRetry retries to update or insert t times unit error is not 'EOF' problem.
func upsertRetry(t uint32, c *mgo.Collection, selector, update interface{}) error {
	var err error

	for {
		_, err = c.Upsert(selector, update)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// removeRetry retries to remove t times unit error is not 'EOF' problem.
func removeRetry(t uint32, c *mgo.Collection, selector interface{}) error {
	var err error

	for {
		err = c.Remove(selector)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// removeAllRetry retries to remove t times unit error is not 'EOF' problem.
func removeAllRetry(t uint32, c *mgo.Collection, selector interface{}) error {
	var err error

	for {
		_, err = c.RemoveAll(selector)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}

// countRetry retries to count t times unit error is not 'EOF' problem.
func countRetry(t uint32, c *mgo.Collection, selector bson.M) (int, error) {
	var count int
	var err error

	for {
		count, err = c.Find(selector).Count()
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return 0, err
		}
		break
	}

	return count, err
}

// pipeGetRetry retries to find t times unit error is not 'EOF' problem.
func pipeGetRetry(t uint32, c *mgo.Collection, stages []bson.M, doc interface{}) error {
	var err error

	for {
		err = c.Find(stages).All(doc)
		if err != nil && err.Error() == "EOF" {
			if t > 0 {
				c.Database.Session.Refresh()
				t--
				continue
			}
			return err
		}
		break
	}

	return err
}
