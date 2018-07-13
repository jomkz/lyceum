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

// OrganizationService encapulates high-level organization logic.
type OrganizationService struct {
	Session *r.Session
	Table   r.Term
}

// NewOrganizationService instantiates a new organization service.
func NewOrganizationService(s *r.Session) *OrganizationService {
	return &OrganizationService{Session: s, Table: r.DB("lyceum").Table("organization")}
}

// Create will create a new organization resource.
func (s OrganizationService) Create(organization interface{}) (*models.Organization, error) {
	res := new(models.Organization)
	err := db.InsertRethinkDBDocument(organization, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the organization with the given id.
func (s OrganizationService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a organization resource with the given id.
func (s OrganizationService) Get(id string) (*models.Organization, error) {
	organization := new(models.Organization)
	err := db.GetRethinkDBDocument(id, organization, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

// List will retrieve the all organization resources.
func (s OrganizationService) List() ([]models.Organization, error) {
	organizations := make([]models.Organization, 0)
	err := db.GetRethinkDBAllDocuments(&organizations, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

// Update will update the organization resource with the given id.
func (s OrganizationService) Update(id string, organization interface{}) (*models.Organization, error) {
	res := new(models.Organization)
	err := db.UpdateRethinkDBDocument(id, organization, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
