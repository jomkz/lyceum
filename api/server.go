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

  "github.com/sirupsen/logrus"
)

func index(w http.ResponseWriter, r *http.Request) {
  logrus.Infof("handle index")
  w.Write([]byte("lyceum"))
}

func Listen() {
  logrus.Infof("listening for connections...")
  http.HandleFunc("/", index)
  http.ListenAndServe(":4778", nil)
}
