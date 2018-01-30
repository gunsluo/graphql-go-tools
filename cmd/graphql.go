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
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gunsluo/graphql-go-tools/generator"
)

const (
	version = "1.0.0"
)

func main() {
	outPath := flag.String("out", ".", "output path")
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")

	flag.Parse()
	if *h {
		flag.Usage()
		os.Exit(0)
	}

	if *v {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	g := generator.New()
	if err := g.Read(flag.Args()); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	if err := g.Output(*outPath); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
