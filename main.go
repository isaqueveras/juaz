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
		fmt.Println("juaz: inform the juazeiro model file")
		return
	}

	fileIn, err := os.Open(*pathFile)
	if err != nil {
		panic(err)
	}

	if filepath.Ext(fileIn.Name()) != ".juaz" {
		fmt.Println("juaz: the juaz file must have the extension .juaz")
		return
	}

	parser := participle.MustBuild[grammar.Juaz](participle.UseLookahead(2))
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
