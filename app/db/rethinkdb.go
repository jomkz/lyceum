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

package db

import (
	"fmt"

	"github.com/jmckind/lyceum/app/models"
	"github.com/revel/revel"
	r "gopkg.in/gorethink/gorethink.v4"
)

// ConnectRethinkDB connects to a rethinkdb database
func ConnectRethinkDB(opts map[string]interface{}) (*r.Session, error) {
	revel.AppLog.Debugf("connecting to database...")
	s, err := r.Connect(r.ConnectOpts{
		Address:    opts["db_url"].(string),
		InitialCap: opts["db_con_initial"].(int),
		MaxOpen:    opts["db_con_max"].(int),
	})
	if err != nil {
		return nil, err
	}
	revel.AppLog.Debugf("connected to database")
	return s, initializeRethinkDB(s)
}

// DeleteRethinkDBDocument will delete the document with the given ID from rethinkdb
func DeleteRethinkDBDocument(key string, table r.Term, session *r.Session) error {
	res, err := table.Get(key).Delete().Run(session)
	if err != nil {
		revel.AppLog.Errorf("unable to run delete: %v", err)
		return err
	}
	defer res.Close()
	return nil
}

// InsertRethinkDBDocument will insert the given document in rethinkdb
func InsertRethinkDBDocument(doc interface{}, table r.Term, session *r.Session) (*models.Item, error) {
	res, err := table.Insert(doc).RunWrite(session)
	if err != nil {
		revel.AppLog.Errorf("unable to run insert: %v", err)
		return nil, err
	}
	if res.Inserted != 1 {
		return nil, fmt.Errorf("Inserted unexpected document count: %d", res.Inserted)
	}
	key := res.GeneratedKeys[0]
	item, err := GetRethinkDBDocument(key, table, session)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetRethinkDBDocument will get the document with the given ID from rethinkdb
func GetRethinkDBDocument(key string, table r.Term, session *r.Session) (*models.Item, error) {
	res, err := table.Get(key).Run(session)
	if err != nil {
		revel.AppLog.Errorf("unable to run get: %v", err)
		return nil, err
	}
	defer res.Close()

	item := new(models.Item)
	err = res.One(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetRethinkDBAllDocuments will get all documents from the given rethinkdb table
func GetRethinkDBAllDocuments(table r.Term, session *r.Session) ([]models.Item, error) {
	res, err := table.Run(session)
	//res, err := ic.table.Filter(r.Row.Field("status").Ne("deleted")).Run(ic.session)
	if err != nil {
		revel.AppLog.Errorf("unable to run query: %v", err)
		return nil, err
	}
	defer res.Close()

	items := make([]models.Item, 0)
	err = res.All(&items)
	if err != nil {
		revel.AppLog.Errorf("unable to get all rows: %v", err)
		return nil, err
	}

	return items, nil
}

// UpdateRethinkDBDocument will update the given document in rethinkdb
func UpdateRethinkDBDocument(id string, doc interface{}, table r.Term, session *r.Session) (*models.Item, error) {
	res, err := table.Get(id).Update(doc).RunWrite(session)
	if err != nil {
		revel.AppLog.Errorf("unable to run update: %v", err)
		return nil, err
	}
	if res.Replaced != 1 {
		return nil, fmt.Errorf("unexpected document replaced count: %d", res.Replaced)
	}
	item, err := GetRethinkDBDocument(id, table, session)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func initializeRethinkDB(session *r.Session) error {
	revel.AppLog.Debugf("initializing database...")

	createDatabase("lyceum", session)
	createTable("lyceum", "library", session)

	revel.AppLog.Debugf("initialized database")
	return nil
}

func databaseExists(name string, session *r.Session) bool {
	res, err := r.DBList().Run(session)
	if err != nil {
		revel.AppLog.Errorf("unable to list databases: %v", err)
		return false
	}
	defer res.Close()

	var rows []interface{}
	err = res.All(&rows)
	if err != nil {
		revel.AppLog.Errorf("unable to process database list: %v", err)
		return false
	}

	for _, db := range rows {
		if db == name {
			return true
		}
	}
	return false
}

func createDatabase(name string, session *r.Session) error {
	if databaseExists(name, session) {
		return nil
	}

	res, err := r.DBCreate(name).Run(session)
	if err != nil {
		return err
	}
	defer res.Close()

	var row interface{}
	err = res.One(&row)
	if err == r.ErrEmptyResult {
		revel.AppLog.Debugf("row not found")
		return err
	}
	if err != nil {
		revel.AppLog.Debugf("unable to read row: %v")
		return err
	}
	return nil
}

func createTable(db string, name string, session *r.Session) error {
	if tableExists(db, name, session) {
		return nil
	}

	res, err := r.DB(db).TableCreate(name).Run(session)
	if err != nil {
		return err
	}
	defer res.Close()

	var row interface{}
	err = res.One(&row)
	if err == r.ErrEmptyResult {
		revel.AppLog.Debugf("row not found")
		return err
	}
	if err != nil {
		revel.AppLog.Debugf("unable to read row: %v")
		return err
	}
	return nil
}

func tableExists(db string, name string, session *r.Session) bool {
	res, err := r.DB(db).TableList().Run(session)
	if err != nil {
		revel.AppLog.Errorf("unable to list tables: %v", err)
		return false
	}
	defer res.Close()

	var rows []interface{}
	err = res.All(&rows)
	if err != nil {
		revel.AppLog.Errorf("unable to process table list: %v", err)
		return false
	}

	for _, table := range rows {
		if table == name {
			return true
		}
	}
	return false
}
