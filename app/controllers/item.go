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

	"github.com/jmckind/lyceum/app/db"
	"github.com/jmckind/lyceum/app/models"
	"github.com/revel/revel"
	rethink "gopkg.in/gorethink/gorethink.v4"
)

type ItemAPI struct {
	LyceumController
}

func (c ItemAPI) Create() revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	item := new(models.Item)
	item.Author = "unknown"
	item.DateCreated = utc
	item.DateModified = utc
	item.Name = "New Item"
	item.Status = "new"

	item, err := db.InsertRethinkDBDocument(item, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to insert document: %v", err)
		return c.RenderError(err)
	}
	result := map[string]interface{}{
		"item": item,
	}
	return c.RenderJSON(result)
}

func (c ItemAPI) Delete(id string) revel.Result {
	err := db.DeleteRethinkDBDocument(id, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to delete document: %v", err)
		return c.RenderError(err)
	}
	c.Response.Status = http.StatusNoContent
	return c.RenderText("")
}

func (c ItemAPI) Get(id string) revel.Result {
	item, err := db.GetRethinkDBDocument(id, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to get document: %v", err)
		return c.RenderError(err)
	}
	result := map[string]interface{}{
		"item": item,
	}
	return c.RenderJSON(result)
}

func (c ItemAPI) List() revel.Result {
	items, err := db.GetRethinkDBAllDocuments(c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to get all documents: %v", err)
		return c.RenderError(err)
	}
	result := map[string]interface{}{
		"items": items,
		"total": len(items),
	}
	return c.RenderJSON(result)
}

func (c ItemAPI) Update(id string) revel.Result {
	item := new(models.Item)
	item.DateModified = time.Now().UTC().Format(time.RFC3339)

	item, err := db.UpdateRethinkDBDocument(id, item, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to insert document: %v", err)
		return c.RenderError(err)
	}
	result := map[string]interface{}{
		"item": item,
	}
	return c.RenderJSON(result)
}

func (c ItemAPI) getTable() rethink.Term {
	return c.getDB().Table("library")
}
