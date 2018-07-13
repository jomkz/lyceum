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

package controllers

import (
	"strings"
	"time"

	"github.com/jmckind/lyceum/app"
	"github.com/revel/revel"
)

// Library controller for library resources.
type Library struct {
	LyceumController
}

// Create will save a new library.
func (c Library) Create() revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	library := make(map[string]interface{})
	library["DateCreated"] = utc
	library["DateModified"] = utc
	library["Description"] = c.Params.Get("description")
	library["Name"] = c.Params.Get("name")
	library["Tags"] = strings.Split(c.Params.Get("tags"), ",")

	_, err := app.Services.LibraryService.Create(library)
	if err != nil {
		return c.RenderError(err)
	}
	return c.Redirect((*Library).List)
}

// Detail will show an existing library.
func (c Library) Detail(id string) revel.Result {
	library, err := app.Services.LibraryService.Get(id)
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["library"] = library
	return c.Render()
}

// List will retrieve all library resources.
func (c Library) List() revel.Result {
	libraries, err := app.Services.LibraryService.List()
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["libraries"] = libraries
	return c.Render()
}
