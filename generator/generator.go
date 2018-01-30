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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-errors/errors"
)

type Generator struct {
	caches    map[string][]byte
	validator *validation
}

func New() *Generator {
	return &Generator{
		caches:    make(map[string][]byte),
		validator: &validation{},
	}
}

func (g *Generator) Read(args []string) error {
	if err := g.verify(args); err != nil {
		return err
	}

	if err := g.read(args); err != nil {
		return err
	}

	// syntax
	if err := g.verifySyntax(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) verify(apath []string) error {
	if len(apath) == 0 {
		return errors.Errorf("input args nil.")
	}

	for _, fp := range apath {
		if err := g.validator.verify(fp); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) read(apath []string) error {
	for _, fp := range apath {
		if data, err := ioutil.ReadFile(fp); err != nil {
			return errors.Wrap(err, 0)
		} else {
			g.caches[fp] = data
		}
	}

	return nil
}

func (g *Generator) verifySyntax() error {
	for fp, data := range g.caches {
		if err := g.validator.verifySyntax(data); err != nil {
			return errors.Errorf("%s: %s", fp, err)
		}
	}

	return nil
}

func (g *Generator) Output(outPath string) error {
	for fp, data := range g.caches {
		op := filepath.Base(fp)
		nf := strings.TrimSuffix(op, filepath.Ext(op))
		np := fmt.Sprintf("%s.go", filepath.Join(outPath, nf))
		fmt.Printf("%s: %s\n", np, string(data))
	}

	return nil
}
