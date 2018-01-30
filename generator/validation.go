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
package generator

import (
	"os"

	"github.com/go-errors/errors"
)

type validation struct {
}

func (v *validation) verify(fp string) error {
	if exist, err := pathExists(fp); err != nil || !exist {
		return errors.Wrap(err, 0)
	}

	return nil
}

func (v *validation) verifySyntax(data []byte) error {
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, errors.Errorf("%s not found.", path)
	}
	return false, err
}
