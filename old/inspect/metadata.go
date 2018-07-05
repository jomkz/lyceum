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

package inspect

import (
	"bytes"

	"github.com/jmckind/lyceum/store"
	"github.com/kapmahc/epub"
	"github.com/ledongthuc/pdf"
	"github.com/sirupsen/logrus"
	r "gopkg.in/gorethink/gorethink.v4"
)

// ProcessNewItems will process the metadata for all new items
func ProcessNewItems() {
	logrus.Infof("process new items")
	opts := map[string]interface{}{
		"listen_ip":      "",
		"listen_port":    4778,
		"db_url":         "localhost:28015",
		"db_con_initial": 10,
		"db_con_max":     10,
	}
	session, err := store.ConnectRethinkDB(opts)
	if err != nil {
		logrus.Fatalf("unable to connect to database: %v", err)
	}
	items, err := store.GetRethinkDBAllDocuments(r.DB("lyceum").Table("Library"), session)
	if err != nil {
		logrus.Fatalf("unable to get items")
	}
	for _, item := range items {
		if item.Status != "new" {
			continue
		}

		if item.FileType == "epub" {
			readEPUB(item.Location)
		} else if item.FileType == "pdf" {
			readPDF(item.Location)
		}
	}
}

func readEPUB(filename string) {
	logrus.Infof("read epub: %s", filename)
	bk, err := epub.Open(filename)
	if err != nil {
		logrus.Errorf("unable to open epub file: %v", err)
		return
	}
	defer bk.Close()

	logrus.Infof("files: %+v", bk.Files())
	logrus.Infof("book: %+v", bk)
}

func readPDF(filename string) {
	logrus.Infof("read pdf: %s", filename)
	r, err := pdf.Open(filename)
	if err != nil {
		logrus.Errorf("unable to open pdf file: %v", err)
		return
	}

	var buf bytes.Buffer
	plain, errp := r.GetPlainText()
	if errp != nil {
		logrus.Errorf("unable to read pdf contents: %v", err)
		return
	}
	buf.ReadFrom(plain)

	logrus.Infof("book: %+v", buf)
}
