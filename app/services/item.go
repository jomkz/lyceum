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

package services

import (
	"github.com/jmckind/lyceum/app/db"
	"github.com/jmckind/lyceum/app/models"
	r "gopkg.in/gorethink/gorethink.v4"
)

// ItemService encapulates high-level item logic.
type ItemService struct {
	Session *r.Session
	Table   r.Term
}

// NewItemService instantiates a new item service.
func NewItemService(s *r.Session) *ItemService {
	return &ItemService{Session: s, Table: r.DB("lyceum").Table("item")}
}

// Create will create a new item resource.
func (s ItemService) Create(item interface{}) (*models.Item, error) {
	res := new(models.Item)
	err := db.InsertRethinkDBDocument(item, &res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the item with the given id.
func (s ItemService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a item resource with the given id.
func (s ItemService) Get(id string) (*models.Item, error) {
	item := new(models.Item)
	err := db.GetRethinkDBDocument(id, &item, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// List will retrieve the all item resources.
func (s ItemService) List() ([]models.Item, error) {
	items := make([]models.Item, 0)
	err := db.GetRethinkDBAllDocuments(&items, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Update will update the item resource with the given id.
func (s ItemService) Update(id string, item interface{}) (*models.Item, error) {
	res := new(models.Item)
	err := db.UpdateRethinkDBDocument(id, item, &res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
