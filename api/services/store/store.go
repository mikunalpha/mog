package store

import (
	"fmt"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/mikunalpha/mog/api/shared/store"
	"github.com/mikunalpha/mog/api/shared/store/mongo"
	"github.com/mikunalpha/mog/api/shared/utils"
)

// S is the global instnace of store.Store.
var S store.Store

// Init intializes S.
func Init() error {
	var err error

	addrs := strings.Split(utils.EnvString("MONGO_ADDRESSES", "127.0.0.1:27017"), ",")

	for _ = range addrs {
		S, err = mongo.NewStore(&mgo.DialInfo{
			Addrs:    addrs,
			Username: utils.EnvString("MONGO_USERNAME", "root"),
			Password: utils.EnvString("MONGO_PASSWORD", "pass"),
			Source:   "admin",
			Database: utils.EnvString("MONGO_DATABASENAME", "mog"),
			Timeout:  5 * time.Second,
		})
		if err == nil {
			break
		}
	}

	if err != nil {
		return fmt.Errorf("Cannot bootup service Mongo. %s", err)
	}

	return nil
}
