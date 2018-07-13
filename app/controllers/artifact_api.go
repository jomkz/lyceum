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
	"github.com/jmckind/lyceum/app"
	"github.com/revel/revel"
)

// ArtifactAPI is the controller for the artifact API resource.
type ArtifactAPI struct {
	LyceumController
}

// List will retrieve all artifact resources.
func (c ArtifactAPI) List() revel.Result {
	artifacts, err := app.Services.ArtifactService.List()
	if err != nil {
		return c.RenderJSONError(err)
	}
	result := map[string]interface{}{
		"artifacts": artifacts,
		"total":     len(artifacts),
	}
	return c.RenderJSON(result)
}
