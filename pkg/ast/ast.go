package ast

type RootNodeType int

type Type int

// These are available root nodes type, In JSON it will either
// be an array of objects or objects
const (
	ObjectRoot RootNodeType = iota
	ArrayRoot
)

// Identifier represents a JSON object property key
type Identifier struct {
	Type      Type
	Value     string
	Delimiter string
}

// LiteralValueType is a type alias for int. Represents the type of the value in a Literal node
type LiteralValueType int

// Available ast value types
const (
	ObjectType Type = iota
	ArrayType
	ArrayItemType
	LiteralType
	PropertyType
	IdentifierType
)

const (
	StringLiteralValueType LiteralValueType = iota
	NumberLiteralValueType
	NullLiteralValueType
	BooleanLiteralValueType
)

type state int

// Available states for each type used in parsing
const (
	// Object states
	ObjectStart state = iota
	ObjectOpen
	ObjectProperty
	ObjectComma

	// Property states
	PropertyStart
	PropertyKey
	PropertyColon
	PropertyValue

	// Array states
	ArrayStart
	ArrayOpen
	ArrayValue
	ArrayComma

	// String states
	StringStart
	StringQuoteOrChar
	Escape

	// Number states
	NumberStart
	NumberMinus
	NumberZero
	NumberDigit
	NumberPoint
	NumberDigitFraction
	NumberExp
	NumberExpDigitOrSign
)

// ValueContent will eventually have some methods that all Values must implement. For now
// it represents any JSON value (object | array | boolean | string | number | null)
type Value interface{}

type StructuralItem struct {
	Value string
}

// Property holds a Type ("Property") as well as a `Key` and `Value`. The Key is an Identifier
// and the value is any Value.
type Property struct {
	Type  string // "Property"
	Key   Identifier
	Value Value
}

// Object represents a JSON object. It holds a slice of Property as its children,
// a Type ("Object"), and start & end code points for displaying.
type Object struct {
	Type     string // "Object"
	Children []Property
	Start    int
	End      int
}

// RootNode is what starts every parsed AST. There is a `Type` field so that
// you can ask which root node type starts the tree.
type RootNode struct {
	RootValue *Value
	Type      RootNodeType
}

// Array represents a JSON array It holds a slice of Value as its children,
// a Type ("Array"), and start & end code points for displaying.
type Array struct {
	Type     string // "Array"
	Children []Value
	Start    int
	End      int
}

// Literal represents a JSON literal value. It holds a Type ("Literal") and the actual value.
type Literal struct {
	Type  string // "Literal"
	Value Value
}
