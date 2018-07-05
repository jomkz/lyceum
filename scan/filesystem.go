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

package scan

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jmckind/lyceum/model"
	"github.com/jmckind/lyceum/store"
	"github.com/sirupsen/logrus"
	r "gopkg.in/gorethink/gorethink.v4"
)

// SearchDirectory will scan a directory for items
func SearchDirectory(abspath string) {
	items, err := findFiles(abspath)
	if err != nil {
		logrus.Fatalf("Unable to search directory: %v", err)
		return
	}

	err = saveItems(items)
	if err != nil {
		logrus.Fatalf("Unable to save items: %v", err)
		return
	}
}

func findFiles(dir string) ([]string, error) {
	logrus.Infof("scanning path: %s", dir)
	fileList := make([]string, 0)

	errw := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".pdf") || strings.HasSuffix(path, ".epub") {
			fileList = append(fileList, path)
		}
		return err
	})
	if errw != nil {
		panic(errw)
	}
	return fileList, nil
}

func saveItems(items []string) error {
	logrus.Infof("saving %d files", len(items))
	for _, file := range items {
		newItem := new(model.Item)

		paths := strings.Split(file, "/")
		filename := paths[len(paths)-1]
		parts := strings.Split(filename, ".")

		newItem.Name = parts[0]
		newItem.FileType = parts[len(parts)-1]
		newItem.Location = string(file)

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

		item, err := store.InsertRethinkDBDocument(newItem, r.DB("lyceum").Table("library"), session)
		if err != nil {
			return err
		}
		logrus.Debugf("saved item: %v", item)
	}
	return nil
}
