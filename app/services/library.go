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

// LibraryService encapulates high-level library logic.
type LibraryService struct {
	Session *r.Session
	Table   r.Term
}

// NewLibraryService instantiates a new library service.
func NewLibraryService(s *r.Session) *LibraryService {
	return &LibraryService{Session: s, Table: r.DB("lyceum").Table("library")}
}

// Create will create a new library resource.
func (s LibraryService) Create(library interface{}) (*models.Library, error) {
	res := new(models.Library)
	err := db.InsertRethinkDBDocument(library, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the library with the given id.
func (s LibraryService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a library resource with the given id.
func (s LibraryService) Get(id string) (*models.Library, error) {
	library := new(models.Library)
	err := db.GetRethinkDBDocument(id, library, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return library, nil
}

// List will retrieve the all library resources.
func (s LibraryService) List() ([]models.Library, error) {
	libraries := make([]models.Library, 0)
	err := db.GetRethinkDBAllDocuments(&libraries, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return libraries, nil
}

// Update will update the library resource with the given id.
func (s LibraryService) Update(id string, library interface{}) (*models.Library, error) {
	res := new(models.Library)
	err := db.UpdateRethinkDBDocument(id, library, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
