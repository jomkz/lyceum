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

package models

// Item model represents an item in the library
type Item struct {
	ArtifactID     string   `json:"artifact_id"`
	Author         string   `json:"author"`
	Content        []byte   `json:"content"`
	ContentType    string   `json:"content_type"`
	DateCreated    string   `json:"date_created"`
	DateModified   string   `json:"date_modified"`
	Description    string   `json:"description"`
	Filename       string   `json:"filename"`
	Hash           string   `json:"hash"`
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	OrganizationID string   `json:"organization_id"`
	Size           int      `json:"size"`
	Status         string   `json:"status"`
	Tags           []string `json:"tags"`
}
