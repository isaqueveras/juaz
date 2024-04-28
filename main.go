//  Created by Isaque Veras on 03/15/24.
//  Copyright © 2024 Isaque Veras. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/participle/v2"

	"github.com/isaqueveras/juaz/golang"
	"github.com/isaqueveras/juaz/grammar"
)

func main() {
	pathFile := flag.String("file", "", "path to juaz file")
	flag.Parse()

	if pathFile == nil || *pathFile == "" {
		fmt.Fprintln(os.Stderr, "juaz: inform the juazeiro model file\n\nUsage: -file=<path/file.juaz>")
		os.Exit(2)
	}

	fileIn, err := os.Open(*pathFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "juaz: unable to open Juaz model file")
		os.Exit(2)
	}

	if filepath.Ext(fileIn.Name()) != ".juaz" {
		fmt.Fprintln(os.Stderr, "juaz: the juaz file must have the extension .juaz\n\nUsage: -file=<path/file.juaz>")
		os.Exit(2)
	}

	parser := participle.MustBuild[grammar.Juaz](
		participle.UseLookahead(2),
		participle.Unquote("String"),
	)

	juaz, _ := parser.Parse("", fileIn)
	fileOut, err := os.Create(juaz.Pos.Filename + "_client.go")
	if err != nil {
		panic(err)
	}

	fileOut.WriteString(
		golang.Write(juaz),
	)

	b := bufio.NewWriter(fileOut)
	if err := b.Flush(); err != nil {
		panic(err)
	}
}
