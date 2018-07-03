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
	"github.com/jmckind/lyceum/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// Listen starts the web server
func Listen() {
	e := echo.New()
	addMiddleware(e)
	addRoutes(e)
	e.Start(":4778")
}

func addMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func addRoutes(e *echo.Echo) {
	e.Validator = &model.ItemValidator{Validator: validator.New()}
	e.GET("/items", listItems)
	e.POST("/items", createItem)
	e.GET("/items/:id", getItem)
	e.PUT("/items/:id", updateItem)
	e.DELETE("/items/:id", deleteItem)
}
