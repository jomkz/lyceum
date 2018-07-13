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

// ItemAPI is the controller for the item API resource.
type ItemAPI struct {
	LyceumController
}

// Create will add a new item resource.
func (c ItemAPI) Create() revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	item := make(map[string]interface{})
	item["Author"] = "unknown"
	item["DateCreated"] = utc
	item["DateModified"] = utc
	item["Name"] = "New Item"
	item["Status"] = "new"

	res, err := app.Services.ItemService.Create(item)
	if err != nil {
		return c.RenderJSONError(err)
	}

	result := map[string]interface{}{
		"item": res,
	}
	return c.RenderJSON(result)
}

// Delete will remove the item resource with the given id.
func (c ItemAPI) Delete(id string) revel.Result {
	err := app.Services.ItemService.Delete(id)
	if err != nil {
		return c.RenderJSONError(err)
	}
	c.Response.Status = http.StatusNoContent
	return c.RenderText("")
}

// Get will retrieve the item resource with the given id.
func (c ItemAPI) Get(id string) revel.Result {
	item, err := app.Services.ItemService.Get(id)
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"item": item,
	}
	return c.RenderJSON(result)
}

// List will retrieve all item resources.
func (c ItemAPI) List() revel.Result {
	items, err := app.Services.ItemService.List()
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"items": items,
		"total": len(items),
	}
	return c.RenderJSON(result)
}

//Update will update the item resource with the given id.
func (c ItemAPI) Update(id string) revel.Result {
	item := make(map[string]interface{})
	item["DateModified"] = time.Now().UTC().Format(time.RFC3339)

	res, err := app.Services.ItemService.Update(id, item)
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"item": res,
	}
	return c.RenderJSON(result)
}
