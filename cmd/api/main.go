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

package main

import (
	"github.com/jmckind/lyceum/api"
	"github.com/jmckind/lyceum/util"
	"github.com/jmckind/lyceum/version"
)

func main() {
	util.PrintVersion("lyceum-api", version.Version)
	opts := map[string]interface{}{
		"listen_ip":      "",
		"listen_port":    4778,
		"db_url":         "localhost:28015",
		"db_con_initial": 10,
		"db_con_max":     10,
	}
	server := api.NewLyceumServer(opts)
	server.Listen()
}
