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

	"github.com/jmckind/lyceum/app"
	"github.com/revel/revel"
)

// LyceumController adds application specific properties.
type LyceumController struct {
	*revel.Controller
}

// RenderJSONError will render a generic error message in JSON format.
func (c LyceumController) RenderJSONError(err error) revel.Result {
	res := map[string]string{
		"message": "An unexpected error has occurred.",
	}
	c.Response.Status = http.StatusInternalServerError
	c.Log.Errorf("rendering JSON error: %v", err)
	return c.RenderJSON(res)
}

// App controller for general application resources.
type App struct {
	LyceumController
}

// Index renders the main "home" page for the application.
func (c App) Index() revel.Result {
	orgs, err := app.Services.OrganizationService.List()
	if err != nil {
		return c.RenderJSONError(err)
	}
	c.ViewArgs["orgs"] = orgs
	return c.Render()
}

// Search will perform an application-wide search and display the results.
func (c App) Search(q string) revel.Result {
	c.ViewArgs["term"] = q
	return c.Render()
}
