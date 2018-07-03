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

package model

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

// Item type
type Item struct {
	ID       string `json:"id" validate:"-"`
	Hash     string `json:"hash" validate:"-"`
	Name     string `json:"name" validate:"required,min=3"`
	Location string `json:"location" validate:"-"`
	FileType string `json:"file_type" validate:"-"`
	Status   string `json:"status" validate:"-"`
}

// ItemValidator type
type ItemValidator struct {
	Validator *validator.Validate
}

// Validate will perform validation for the item
func (iv *ItemValidator) Validate(i interface{}) error {
	err := iv.Validator.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logrus.Infof(err.Namespace())
			logrus.Infof(err.Field())
			logrus.Infof(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			logrus.Infof(err.StructField())     // by passing alt name to ReportError like below
			logrus.Infof(err.Tag())
			logrus.Infof(err.ActualTag())
			// logrus.Infof(err.Kind())
			// logrus.Infof(err.Type())
			// logrus.Infof(err.Value())
			logrus.Infof(err.Param())
		}
	}
	return err
}
