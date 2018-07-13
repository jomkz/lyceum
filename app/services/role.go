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

// RoleService encapulates high-level role logic.
type RoleService struct {
	Session *r.Session
	Table   r.Term
}

// NewRoleService instantiates a new role service.
func NewRoleService(s *r.Session) *RoleService {
	return &RoleService{Session: s, Table: r.DB("lyceum").Table("role")}
}

// Create will create a new role resource.
func (s RoleService) Create(role interface{}) (*models.Role, error) {
	res := new(models.Role)
	err := db.InsertRethinkDBDocument(role, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the role with the given id.
func (s RoleService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a role resource with the given id.
func (s RoleService) Get(id string) (*models.Role, error) {
	role := new(models.Role)
	err := db.GetRethinkDBDocument(id, role, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// List will retrieve the all role resources.
func (s RoleService) List() ([]models.Role, error) {
	roles := make([]models.Role, 0)
	err := db.GetRethinkDBAllDocuments(roles, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// Update will update the role resource with the given id.
func (s RoleService) Update(id string, role interface{}) (*models.Role, error) {
	res := new(models.Role)
	err := db.UpdateRethinkDBDocument(id, role, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
