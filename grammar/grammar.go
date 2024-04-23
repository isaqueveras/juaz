//nolint:all
package grammar

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Juaz struct {
	Pos     lexer.Position
	Entries []*Entry `(@@*)*`
}

type Entry struct {
	Pos lexer.Position

	Package string  `"pkg" @Ident`
	Option  *Option `|@@`
	Type    *Type   `|@@`
	Node    *Node   `|@@`
	Impl    *Impl   `|@@`
}

type Option struct {
	Pos lexer.Position

	Name  string `"option" (@Ident":"`
	Value string `@String)`
}

type Value struct {
	Pos lexer.Position

	String    *string  `  @String`
	Number    *float64 `| @Float`
	Int       *int64   `| @Int`
	Bool      *bool    `| (@"true" | "false")`
	Reference *string  `| @Ident @( "." Ident )*`
	Map       *Map     `| @@`
	Array     *Array   `| @@`
}

type Array struct {
	Pos lexer.Position

	Elements []*Value `"[" ( @@ ( ","? @@ )* )? "]"`
}

type Map struct {
	Pos lexer.Position

	Entries []*MapEntry `"{" ( @@ ( ( "," )? @@ )* )? "}"`
}

type MapEntry struct {
	Pos lexer.Position

	Key   *Value `@@`
	Value *Value `":"? @@`
}

type Impl struct {
	Pos lexer.Position

	Name     string       `"impl" @Ident`
	Input    *ItemType    `"(" @@ ")"`
	Repeated bool         `@ "~"?`
	Output   *ItemType    `@@?`
	Entry    []*ImplEntry `"{" @@* "}"`
}

type ImplEntry struct {
	Pos lexer.Position

	Name  string `(@Ident":"`
	Value string `@String)`
}

type FieldImplEntry struct {
	Pos lexer.Position

	Name  string `(@Ident":"`
	Value string `@String)`
}

type Type struct {
	Pos lexer.Position

	Name   *string      `"type" @Ident`
	Values []*TypeEntry `"{" ( @@ ( ";" )* )* "}"`
}

type TypeEntry struct {
	Pos lexer.Position

	Value *TypeValue `  @@`
}

type TypeValue struct {
	Pos lexer.Position

	Key string `@Ident`
}

type Node struct {
	Pos lexer.Position

	Name    string        `"node" @Ident`
	Entries []*ModelEntry `"{" @@* "}"`
}

type ModelEntry struct {
	Pos lexer.Position

	Field *Field ` ( @@ )`
}

type Field struct {
	Pos lexer.Position

	Name     string    `(@Ident":"`
	Repeated bool      `(@ "repeated")?`
	Type     *ItemType `@@)`
}

type Scalar int

const (
	None Scalar = iota
	Float32
	Float64
	Int
	Int32
	Int64
	Uint32
	Uint64
	Bool
	String
	Byte
	Time
)

var scalarToString = map[Scalar]string{
	None:    "none",
	Float32: "float32",
	Float64: "float64",
	Int:     "int",
	Int32:   "int32",
	Int64:   "int64",
	Uint32:  "uint32",
	Uint64:  "uint64",
	Bool:    "bool",
	String:  "string",
	Byte:    "[]byte",
	Time:    "time.Time",
}

func (s Scalar) GoString() string {
	return scalarToString[s]
}

var stringToScalar = map[string]Scalar{
	"f32":    Float32,
	"f64":    Float64,
	"int":    Int,
	"i32":    Int32,
	"i64":    Int64,
	"uint32": Uint32,
	"uint64": Uint64,
	"bool":   Bool,
	"string": String,
	"byte":   Byte,
	"time":   Time,
}

func (s *Scalar) Parse(lex *lexer.PeekingLexer) error {
	v, ok := stringToScalar[lex.Peek().Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()
	*s = v
	return nil
}

type ItemType struct {
	Pos lexer.Position

	Scalar    Scalar   `  @@`
	Map       *MapType `| @@`
	Reference string   `| @(Ident ( "." Ident )*)`
}

type MapType struct {
	Pos lexer.Position

	Key   *ItemType `"map" "[" @@`
	Value *ItemType `"]" @@`
}
