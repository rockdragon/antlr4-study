package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rockdragon/antlr4-study/output/arithmetic"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	codeFile := filepath.Join(dir, "output/arithmetic/bin/code.source")

	bytes, err := ioutil.ReadFile(codeFile)
	if err != nil {
		log.Fatal(err)
	}

	is := antlr.NewInputStream(string(bytes))

	lexer := parser.NewlangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewlangParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&parser.BaselangListener{}, p.Prog())

	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
}
