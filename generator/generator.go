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
	"bytes"
	"fmt"
	"io"
	"os"
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
	var n int64

	if len(apath) == 0 {
		return nil
	}

	for _, fp := range apath {
		f, err := os.Open(fp)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return errors.Wrap(err, 0)
		}

		if fi.IsDir() {
			var sapath []string
			err := filepath.Walk(fp, func(spath string, sf os.FileInfo, err error) error {
				if sf == nil {
					return errors.Wrap(err, 0)
				}

				if spath != fp {
					sapath = append(sapath, spath)
				}
				return nil
			})
			if err != nil {
				return err
			}

			return g.read(sapath)
		}

		if size := fi.Size(); size < 1e9 {
			n = size
		}

		if ext := filepath.Ext(fp); ext == ".graphql" {
			data, err := readAll(f, n+bytes.MinRead)
			if err != nil {
				return errors.Wrap(err, 0)
			}
			g.caches[fp] = data
		}
	}

	if len(g.caches) == 0 {
		return errors.Errorf("not found graphql file.")
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

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
