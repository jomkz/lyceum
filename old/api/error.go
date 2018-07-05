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
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func validationFailedResponse(err error) map[string]string {
	logrus.Errorf("validation failed: %v", err)
	result := make(map[string]string)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			s := fmt.Sprintf("%s %s %s", err.Field(), err.Tag(), err.Param())
			errors = append(errors, strings.TrimSpace(s))
		}
		result["message"] = strings.Join(errors[:], ".")
	}
	return result
}
