package main

import (
	"github.com/ikawaha/kagome/tokenizer"
	"github.com/k0kubun/pp"
)

func main() {
	t := tokenizer.New()
	tokens := t.Analyze("すもももももももものうち", tokenizer.Normal)
	for _, token := range tokens {
		pp.Printf("%s\n", token.Pos())
		continue
	}
}
