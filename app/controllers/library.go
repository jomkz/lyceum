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
	"bytes"
	"time"

	"github.com/jmckind/lyceum/app/db"
	"github.com/revel/revel"
	rethink "gopkg.in/gorethink/gorethink.v4"
)

type Library struct {
	LyceumController
}

func (c Library) Detail(id string) revel.Result {
	item, err := db.GetRethinkDBDocument(id, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to get document: %v", err)
		return c.RenderError(err)
	}
	c.ViewArgs["item"] = item
	return c.Render()
}

func (c Library) List() revel.Result {
	items, err := db.GetRethinkDBAllDocuments(c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to get all documents: %v", err)
		return c.RenderError(err)
	}
	c.ViewArgs["items"] = items
	return c.Render()
}

func (c Library) Read(id string) revel.Result {
	item, err := db.GetRethinkDBDocument(id, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to get document: %v", err)
		return c.RenderError(err)
	}
	return c.RenderBinary(
		bytes.NewReader(item.Content),
		item.Filename,
		revel.Inline,
		time.Now().UTC(),
	)
}

func (c Library) Search(q string) revel.Result {
	c.ViewArgs["term"] = q
	return c.Render()
}

func (c Library) Upload(data []byte) revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	item := make(map[string]interface{})
	item["Author"] = "unknown"
	item["Content"] = data
	item["ContentType"] = c.Params.Files["data"][0].Header.Get("Content-Type")
	item["DateCreated"] = utc
	item["DateModified"] = utc
	item["Filename"] = c.Params.Files["data"][0].Filename
	item["Name"] = c.Params.Files["data"][0].Filename
	item["Size"] = len(data)
	item["Status"] = "new"
	item["Tags"] = []string{"foo", "bar", "baz"}

	_, err := db.InsertRethinkDBDocument(item, c.getTable(), c.getSession())
	if err != nil {
		c.Log.Errorf("unable to insert document: %v", err)
		return c.RenderError(err)
	}
	return c.Redirect((*Library).List)
}

func (c Library) getTable() rethink.Term {
	return c.getDB().Table("library")
}
