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

package api

import (
	"net/http"
	"time"

	"github.com/jmckind/lyceum/model"
	"github.com/jmckind/lyceum/store"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	r "gopkg.in/gorethink/gorethink.v4"
)

type ItemController struct {
	db      r.Term
	session *r.Session
	table   r.Term
}

func NewItemController(s *r.Session) *ItemController {
	return &ItemController{
		db:      r.DB("lyceum"),
		session: s,
		table:   r.DB("lyceum").Table("library"),
	}
}

// Create will create a new item resource
func (ic *ItemController) Create(c echo.Context) error {
	logrus.Debugf("create item")
	item := new(model.Item)
	if err := c.Bind(item); err != nil {
		logrus.Errorf("unable to bind data: %v", err)
		return err
	}
	if err := c.Validate(item); err != nil {
		return c.JSON(http.StatusBadRequest, validationFailedResponse(err))
	}

	utc := time.Now().UTC().Format(time.RFC3339)
	item.DateCreated = utc
	item.DateModified = utc
	item.Status = "new"

	created, err := store.InsertRethinkDBDocument(item, ic.table, ic.session)
	if err != nil {
		logrus.Errorf("unable to insert document: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}
	result := map[string]interface{}{
		"item": created,
	}
	logrus.Debugf("created item: %v", created)
	return c.JSON(http.StatusCreated, result)
}

// Delete will delete a single resource
func (ic *ItemController) Delete(c echo.Context) error {
	logrus.Debugf("delete item")
	err := store.DeleteRethinkDBDocument(c.Param("id"), ic.table, ic.session)
	if err != nil {
		logrus.Errorf("unable to delete document: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}
	logrus.Debugf("deleted item with id: %s", c.Param("id"))
	return c.NoContent(http.StatusOK)
}

// Get will return a single resource
func (ic *ItemController) Get(c echo.Context) error {
	id := c.Param("id")
	logrus.Debugf("get item with id: %s", id)
	item, err := store.GetRethinkDBDocument(id, ic.table, ic.session)
	if err == r.ErrEmptyResult {
		return c.JSON(http.StatusNotFound, "item not found")
	}
	if err != nil {
		logrus.Errorf("unable to get document: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}
	result := map[string]interface{}{
		"item": item,
	}
	logrus.Debugf("got item: %v", item)
	return c.JSON(http.StatusOK, result)
}

// List will return the list of item resources
func (ic *ItemController) List(c echo.Context) error {
	logrus.Debugf("list items")
	items, err := store.GetRethinkDBAllDocuments(ic.table, ic.session)
	if err != nil {
		logrus.Errorf("unable to get all documents: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}
	result := map[string]interface{}{
		"items": items,
		"total": len(items),
	}
	logrus.Debugf("items: %v", items)
	return c.JSON(http.StatusOK, result)
}

// Update will update an existing item resource
func (ic *ItemController) Update(c echo.Context) error {
	id := c.Param("id")
	logrus.Debugf("update item with id: %s", id)
	item, err := store.GetRethinkDBDocument(id, ic.table, ic.session)
	if err == r.ErrEmptyResult {
		return c.JSON(http.StatusNotFound, "item not found")
	}
	if err != nil {
		logrus.Errorf("unable to get document: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}

	update := new(model.Item)
	if err := c.Bind(update); err != nil {
		logrus.Errorf("unable to bind data: %v", err)
		return err
	}
	if err := c.Validate(update); err != nil {
		return c.JSON(http.StatusBadRequest, validationFailedResponse(err))
	}

	utc := time.Now().UTC().Format(time.RFC3339)
	update.DateModified = utc

	logrus.Infof("old: %v", item)
	logrus.Infof("new: %v", update)

	updated, err := store.UpdateRethinkDBDocument(id, update, ic.table, ic.session)
	if err != nil {
		logrus.Errorf("unable to update document: %v", err)
		return c.JSON(http.StatusInternalServerError, "unexpected error")
	}
	result := map[string]interface{}{
		"item": updated,
	}
	logrus.Debugf("updated item: %v", updated)
	return c.JSON(http.StatusOK, result)
}
