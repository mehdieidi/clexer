package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	http.HandleFunc("/", home)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/panel", panel)
	http.HandleFunc("/download", download)
	http.HandleFunc("/view", view)

	http.ListenAndServe(":6050", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	// max size: 10MB
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("failed to retrieve the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	if filepath.Ext(handler.Filename) != ".c" {
		log.Fatal("file must be a C source code")
	}

	// save the file
	dst, err := os.Create("../../input/source.c")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/panel", http.StatusTemporaryRedirect)
}

func panel(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("panel.html").ParseFiles("../../ui/template/panel.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func download(w http.ResponseWriter, r *http.Request) {
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

	jsonData, err := json.MarshalIndent(&tokens, "", "	")
	if err != nil {
		log.Println("json marshalling failed")
	}

	jsonFile, err := os.Create("../../ui/static/json/tokens.json")
	if err != nil {
		log.Println("creating json file failed")
	}

	jsonFile.Write(jsonData)

	http.Redirect(w, r, "/static/json/tokens.json", 303)
}

func view(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template.New("view.html").ParseFiles("../../ui/template/view.html")
	if err != nil {
		log.Fatal(err)
	}
	data := PageData{
		PageTitle: "Lexical Analysis - Result",
		Tokens:    tokens,
	}

	tmpl.Execute(w, data)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home.html").ParseFiles("../../ui/template/home.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
