package database

import (
	"encoding/json"
	"path/filepath"

	"github.com/maihb/sub2clash/common"
	"github.com/maihb/sub2clash/model"

	"go.etcd.io/bbolt"
)

var DB *bbolt.DB

func ConnectDB() error {
	path := filepath.Join("data", "sub2clash.db")

	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return common.NewDatabaseConnectError(err)
	}
	DB = db

	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ShortLinks"))
		if err != nil {
			return common.NewDatabaseConnectError(err)
		}
		return nil
	})
}

func FindShortLinkByHash(hash string) (*model.ShortLink, error) {
	var shortLink model.ShortLink
	err := DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("ShortLinks"))
		v := b.Get([]byte(hash))
		if v == nil {
			return common.NewRecordNotFoundError("ShortLink", hash)
		}
		return json.Unmarshal(v, &shortLink)
	})
	if err != nil {
		return nil, err
	}
	return &shortLink, nil
}

func SaveShortLink(shortLink *model.ShortLink) error {
	return DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("ShortLinks"))
		encoded, err := json.Marshal(shortLink)
		if err != nil {
			return err
		}
		return b.Put([]byte(shortLink.Hash), encoded)
	})
}

func CheckShortLinkHashExists(hash string) (bool, error) {
	exists := false
	err := DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("ShortLinks"))
		v := b.Get([]byte(hash))
		exists = v != nil
		return nil
	})
	if err != nil {
		return false, err
	}
	return exists, nil
}
