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

// User represents a single user with credentials
type User struct {
	DateCreated   string   `json:"date_created"`
	DateModified  string   `json:"date_modified"`
	DateLastLogin string   `json:"date_last_login"`
	ID            string   `json:"id"`
	Email         string   `json:"email"`
	Name          string   `json:"name"`
	Organizations []string `json:"organizations"`
	Password      string   `json:"password"`
	Roles         []string `json:"roles"`
	Username      string   `json:"username"`
}
