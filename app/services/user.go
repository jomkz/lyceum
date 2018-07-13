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

// UserService encapulates high-level user logic.
type UserService struct {
	Session *r.Session
	Table   r.Term
}

// NewUserService instantiates a new user service.
func NewUserService(s *r.Session) *UserService {
	return &UserService{Session: s, Table: r.DB("lyceum").Table("user")}
}

// Create will create a new user resource.
func (s UserService) Create(user interface{}) (*models.User, error) {
	res := new(models.User)
	err := db.InsertRethinkDBDocument(user, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the user with the given id.
func (s UserService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a user resource with the given id.
func (s UserService) Get(id string) (*models.User, error) {
	user := new(models.User)
	err := db.GetRethinkDBDocument(id, user, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// List will retrieve the all user resources.
func (s UserService) List() ([]models.User, error) {
	users := make([]models.User, 0)
	err := db.GetRethinkDBAllDocuments(users, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update will update the user resource with the given id.
func (s UserService) Update(id string, user interface{}) (*models.User, error) {
	res := new(models.User)
	err := db.UpdateRethinkDBDocument(id, user, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
