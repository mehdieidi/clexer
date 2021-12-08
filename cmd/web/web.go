package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/MehdiEidi/clexer/lexer"
	"github.com/MehdiEidi/clexer/token"
)

type PageData struct {
	PageTitle string
	Tokens    []token.Token
}

func main() {
	fs := http.FileServer(http.Dir("../../ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	src, err := os.ReadFile("../../input/source.c")
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

	tmpl, err := template.New("index.html").ParseFiles("../../ui/template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Lexical Analysis - Result",
			Tokens:    tokens,
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":6050", nil)
}
