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

package db

// Collection represents a named group of documents in the datastore.
type Collection struct {
	Datastore string
	Name      string
}

// Document represents a "record" in the datastore.
type Document struct {
	Collection Collection
	Content    interface{}
}

// Datastore interface abstracts the underlying storage implementation.
type Datastore interface {
	CreateDoc(*Document) (*Document, error)
	DeleteDoc(*Document) error
	GetAllDocs(*Collection) ([]Document, error)
	GetDoc(string, *Collection) (*Document, error)
	UpdateDoc(*Document) (*Document, error)
}
