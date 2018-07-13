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
	r "gopkg.in/gorethink/gorethink.v4"
)

// AuthService encapulates high-level library logic.
type AuthService struct {
	Session *r.Session
}

// NewAuthService instantiates a new library service.
func NewAuthService(s *r.Session) *AuthService {
	return &AuthService{Session: s}
}
