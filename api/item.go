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

	"github.com/jmckind/lyceum/model"
	"github.com/jmckind/lyceum/store"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func createItem(c echo.Context) error {
	logrus.Infof("create item")
	newItem := new(model.Item)
	if err := c.Bind(newItem); err != nil {
		logrus.Errorf("unable to bind model data: %v", err)
		return err
	}
	if err := c.Validate(newItem); err != nil {
		return c.JSON(http.StatusBadRequest, validationFailedResponse(err))
	}

	logrus.Infof("new item: %v", newItem)
	item, err := store.CreateItem(newItem)
	if err != nil {
		msg := "unable to create item"
		logrus.Errorf("%s: %v", msg, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusCreated, item)
}

func deleteItem(c echo.Context) error {
	err := store.DeleteItem(c.Param("id"))
	if err != nil {
		msg := "unable to delete item"
		logrus.Errorf("%s: %v", msg, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.NoContent(http.StatusOK)
}

func getItem(c echo.Context) error {
	item, err := store.GetItem(c.Param("id"))
	if err != nil {
		msg := "unable to get item"
		logrus.Errorf("%s: %v", msg, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, item)
}

func listItems(c echo.Context) error {
	logrus.Infof("list items")
	items, err := store.GetItems()
	if err != nil {
		msg := "unable to get items"
		logrus.Errorf("%s: %v", msg, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	logrus.Infof("items: %v", items)
	return c.JSON(http.StatusOK, items)
}

func updateItem(c echo.Context) error {
	logrus.Infof("update item")
	updItem := new(model.Item)
	if err := c.Bind(updItem); err != nil {
		return err
	}
	if err := c.Validate(updItem); err != nil {
		return c.JSON(http.StatusBadRequest, validationFailedResponse(err))
	}

	logrus.Infof("updated item: %v", updItem)
	item, err := store.UpdateItem(c.Param("id"), updItem)
	if err != nil {
		msg := "unable to update item"
		logrus.Errorf("%s: %v", msg, err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, item)
}
