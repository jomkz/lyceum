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

// ArtifactService encapulates high-level artifact logic.
type ArtifactService struct {
	Session *r.Session
	Table   r.Term
}

// NewArtifactService instantiates a new artifact service.
func NewArtifactService(s *r.Session) *ArtifactService {
	return &ArtifactService{Session: s, Table: r.DB("lyceum").Table("artifact")}
}

// Create will create a new artifact resource.
func (s ArtifactService) Create(artifact interface{}) (*models.Artifact, error) {
	res := new(models.Artifact)
	err := db.InsertRethinkDBDocument(artifact, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete will remove the artifact with the given id.
func (s ArtifactService) Delete(id string) error {
	return db.DeleteRethinkDBDocument(id, s.Table, s.Session)
}

// Get will return a artifact resource with the given id.
func (s ArtifactService) Get(id string) (*models.Artifact, error) {
	artifact := new(models.Artifact)
	err := db.GetRethinkDBDocument(id, artifact, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return artifact, nil
}

// List will retrieve the all artifact resources.
func (s ArtifactService) List() ([]models.Artifact, error) {
	artifacts := make([]models.Artifact, 0)
	err := db.GetRethinkDBAllDocuments(artifacts, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return artifacts, nil
}

// Update will update the artifact resource with the given
func (s ArtifactService) Update(id string, artifact interface{}) (*models.Artifact, error) {
	res := new(models.Artifact)
	err := db.UpdateRethinkDBDocument(id, artifact, res, s.Table, s.Session)
	if err != nil {
		return nil, err
	}
	return res, nil
}
