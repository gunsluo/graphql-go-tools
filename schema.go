// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gqltools

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ImportSchemaDslFromPath input dsl file path, return schema string.
func ImportSchemaDslFromPath(path string) (string, error) {
	dsl, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dsl), nil
}

// ImportSchemaFromDir find .graphql file from dir, return schema string.
func ImportSchemaDslFromDir(dir string) (string, error) {
	var (
		buffer = &bytes.Buffer{}
	)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ext := filepath.Ext(f.Name()); ext != ".graphql" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		buffer.Write(content)
		return nil
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func MakeExecutableSchema() {
	// TODO:
}
