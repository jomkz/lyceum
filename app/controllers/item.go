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
	"crypto/md5"
	"fmt"
	"time"

	"github.com/jmckind/lyceum/app"
	"github.com/revel/revel"
)

// Item controller for item resources.
type Item struct {
	LyceumController
}

// Detail will show an existing item.
func (c Item) Detail(id string) revel.Result {
	item, err := app.Services.ItemService.Get(id)
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["item"] = item
	return c.Render()
}

// List will retrieve all item resources.
func (c Item) List() revel.Result {
	items, err := app.Services.ItemService.List()
	if err != nil {
		return c.RenderError(err)
	}
	c.ViewArgs["items"] = items
	return c.Render()
}

// Read will download the item with the given id.
func (c Item) Read(id string) revel.Result {
	item, err := app.Services.ItemService.Get(id)
	if err != nil {
		return c.RenderError(err)
	}

	artifact, err := app.Services.ArtifactService.Get(item.ArtifactID)
	if err != nil {
		return c.RenderError(err)
	}

	return c.RenderBinary(
		bytes.NewReader(artifact.Content),
		item.Filename,
		revel.Inline,
		time.Now().UTC(),
	)
}

// Upload will create a new item with its associated artifact.
func (c Item) Upload(data []byte) revel.Result {
	utc := time.Now().UTC().Format(time.RFC3339)
	newart := make(map[string]interface{})
	newart["DateCreated"] = utc
	newart["Content"] = data
	newart["Hash"] = fmt.Sprintf("%x", md5.Sum(data))
	newart["Size"] = len(data)

	artifact, err := app.Services.ArtifactService.Create(newart)
	if err != nil {
		c.Log.Errorf("unable to create artifact: %v", err)
		return c.RenderError(err)
	}

	item := make(map[string]interface{})
	item["Author"] = "unknown"
	item["ArtifactID"] = artifact.ID
	item["ContentType"] = c.Params.Files["data"][0].Header.Get("Content-Type")
	item["DateCreated"] = utc
	item["DateModified"] = utc
	item["Filename"] = c.Params.Files["data"][0].Filename
	item["Hash"] = artifact.Hash
	item["Name"] = c.Params.Files["data"][0].Filename
	item["Size"] = len(data)
	item["Status"] = "new"
	item["Tags"] = []string{"foo", "bar", "baz"}

	_, err = app.Services.ItemService.Create(item)
	if err != nil {
		c.Log.Errorf("unable to create item: %v", err)
		return c.RenderError(err)
	}
	return c.Redirect((*Item).List)
}
