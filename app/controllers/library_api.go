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
	"net/http"
	"time"

	"github.com/jmckind/lyceum/app"
	"github.com/revel/revel"
)

// LibraryAPI is the controller for the library API resource.
type LibraryAPI struct {
	LyceumController
}

// Create will add a new library resource.
func (c LibraryAPI) Create() revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	library := make(map[string]interface{})
	library["Author"] = "unknown"
	library["DateCreated"] = utc
	library["DateModified"] = utc
	library["Name"] = "New Library"
	library["Status"] = "new"

	res, err := app.Services.LibraryService.Create(library)
	if err != nil {
		return c.RenderJSONError(err)
	}

	result := map[string]interface{}{
		"library": res,
	}
	return c.RenderJSON(result)
}

// Delete will remove the library resource with the given id.
func (c LibraryAPI) Delete(id string) revel.Result {
	err := app.Services.LibraryService.Delete(id)
	if err != nil {
		return c.RenderJSONError(err)
	}
	c.Response.Status = http.StatusNoContent
	return c.RenderText("")
}

// Get will retrieve the library resource with the given id.
func (c LibraryAPI) Get(id string) revel.Result {
	library, err := app.Services.LibraryService.Get(id)
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"library": library,
	}
	return c.RenderJSON(result)
}

// List will retrieve all library resources.
func (c LibraryAPI) List() revel.Result {
	librarys, err := app.Services.LibraryService.List()
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"libraries": librarys,
		"total":     len(librarys),
	}
	return c.RenderJSON(result)
}

// Update will update the library resource with the given id.
func (c LibraryAPI) Update(id string) revel.Result {
	library := make(map[string]interface{})
	library["DateModified"] = time.Now().UTC().Format(time.RFC3339)

	res, err := app.Services.LibraryService.Update(id, library)
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"library": res,
	}
	return c.RenderJSON(result)
}
