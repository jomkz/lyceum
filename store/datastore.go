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
	"github.com/jmckind/lyceum/model"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// CreateItem will create a new item
func CreateItem(item *model.Item) (*model.Item, error) {
	logrus.Infof("item: %v", item)
	item.ID = uuid.NewV4().String()
	item.Status = "new"
	return saveBoltItem(item)
}

// DeleteItem will delete an item
func DeleteItem(id string) error {
	logrus.Infof("id: %v", id)
	return deleteBoltItem(id)
}

// GetItems will return the list of items
func GetItems() ([]model.Item, error) {
	return getBoltItems()
}

// GetItem will return a single item with the given ID
func GetItem(id string) (*model.Item, error) {
	return getBoltItem(id)
}

// UpdateItem will delete an item
func UpdateItem(id string, item *model.Item) (*model.Item, error) {
	logrus.Infof("item: %v", item)
	oldItem, err := getBoltItem(id)
	if err != nil {
		return nil, err
	}

	if len(item.Hash) > 0 {
		oldItem.Hash = item.Hash
	}
	if len(item.Name) > 0 {
		oldItem.Name = item.Name
	}
	if len(item.Location) > 0 {
		oldItem.Location = item.Location
	}
	if len(item.Status) > 0 {
		oldItem.Status = item.Status
	}

	return saveBoltItem(oldItem)
}
