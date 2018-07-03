// Copyright 2018 Lyceum Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"encoding/json"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/jmckind/lyceum/model"
	"github.com/sirupsen/logrus"
)

// deleteBoltItem will delete an item from a bolt database
func deleteBoltItem(id string) error {
	db, err := openBoltDatabase("lyceum.db")
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("library"))
		return b.Delete([]byte(id))
	})
	return err
}

// GetBoltItem will return a single item from a bolt database
func getBoltItem(id string) (item *model.Item, err error) {
	db, err := openBoltDatabase("lyceum.db")
	if err != nil {
		return
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("library"))
		val := b.Get([]byte(id))
		item = new(model.Item)
		erru := json.Unmarshal(val, item)
		if erru != nil {
			return erru
		}
		return nil
	})
	return
}

// getBoltItems will return an array of items from a bolt database
func getBoltItems() (items []model.Item, err error) {
	db, erro := openBoltDatabase("lyceum.db")
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("library"))
		c := b.Cursor()
		for key, val := c.First(); key != nil; key, val = c.Next() {
			logrus.Infof("key=%s, value=%s", key, val)
			item := new(model.Item)
			erru := json.Unmarshal(val, item)
			if erru != nil {
				return erru
			}
			items = append(items, *item)
		}
		return nil
	})
	return items, err
}

// saveBoltItem will save an item in a bolt database
func saveBoltItem(item *model.Item) (*model.Item, error) {
	db, err := openBoltDatabase("lyceum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		buf, errm := json.Marshal(item)
		if errm != nil {
			return errm
		}
		b := tx.Bucket([]byte("library"))
		return b.Put([]byte(item.ID), buf)
	})
	return item, err
}

// initializeDatabase will create the buckets needed in a bolt database
func initializeDatabase(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("library"))
		return err
	})
}

// openBoltDatabase will open the given bolt database file
func openBoltDatabase(filename string) (*bolt.DB, error) {
	opts := &bolt.Options{Timeout: 1 * time.Second}
	db, err := bolt.Open(filename, 0600, opts)
	if err != nil {
		logrus.Errorf("unable to open database: %v", err)
		return nil, err
	}

	err = initializeDatabase(db)
	if err != nil {
		logrus.Errorf("unable to initialize database: %v", err)
		return nil, err
	}
	return db, nil
}
