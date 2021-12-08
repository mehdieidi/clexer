package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/MehdiEidi/clexer/lexer"
	"github.com/MehdiEidi/clexer/token"
)

func main() {
	src, err := os.ReadFile("./input/source.c")
	if err != nil {
		log.Fatal(err)
	}

	var tokens []token.Token
	var tok token.Token

	lxr := lexer.New(string(src))

	for tok.Type != token.EOF {
		tok = lxr.NextToken()
		tokens = append(tokens, tok)
	}

	jsonData, err := json.MarshalIndent(&tokens, "", "	")
	if err != nil {
		log.Println("json marshalling failed")
	}

	jsonFile, err := os.Create("./output/tokens.json")
	if err != nil {
		log.Println("creating json file failed")
	}

	jsonFile.Write(jsonData)
}
