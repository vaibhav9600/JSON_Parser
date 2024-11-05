package main

import (
	"fmt"
	"json-parser-and-query-tool/pkg/lexer"
	"json-parser-and-query-tool/pkg/parser"
)

func main() {
	const testJSONObject = `{
		"item1": ["aryitem1", "aryitem2", {"some": {"thing": "coolObj"}}],
		"item2": "simplestringvalue"
	}`

	c := lexer.New(testJSONObject)

	p := parser.New(c)

	jsonParsed, err := p.ParseProgram()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonParsed)
}
