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
	"fmt"

	"github.com/jmckind/lyceum/model"
	"github.com/jmckind/lyceum/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	r "gopkg.in/gorethink/gorethink.v4"
)

// LyceumServer represents the properties for a lyceum server
type LyceumServer struct {
	options     map[string]interface{}
	session     *r.Session
	controllers []interface{}
}

// LyceumContext is a custom context for echo
type LyceumContext struct {
	echo.Context
	Server *LyceumServer
}

// NewLyceumServer will create a new server instance
func NewLyceumServer(opts map[string]interface{}) *LyceumServer {
	logrus.SetLevel(logrus.DebugLevel)
	session, err := store.ConnectRethinkDB(opts)
	if err != nil {
		logrus.Fatalf("unable to connect to database: %v", err)
	}
	return &LyceumServer{options: opts, session: session}
}

// Listen starts the web server
func (ls *LyceumServer) Listen() {
	e := echo.New()
	ls.addMiddleware(e)
	ls.addRoutes(e)
	e.Start(
		fmt.Sprintf("%s:%d", ls.options["listen_ip"], ls.options["listen_port"]),
	)
}

func (ls *LyceumServer) addMiddleware(e *echo.Echo) {
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &LyceumContext{c, ls}
			return h(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func (ls *LyceumServer) addRoutes(e *echo.Echo) {
	e.Validator = &model.ItemValidator{Validator: validator.New()}

	ic := NewItemController(ls.session)
	e.GET("/items", ic.List)
	e.POST("/items", ic.Create)
	e.GET("/items/:id", ic.Get)
	e.PUT("/items/:id", ic.Update)
	e.DELETE("/items/:id", ic.Delete)
	ls.controllers = append(ls.controllers, ic)
}
