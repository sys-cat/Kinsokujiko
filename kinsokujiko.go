package kinsokujiko

import (
	"github.com/ikawaha/kagome/tokenizer"
)

type Master struct {
	Sentence string
}

type Surface struct {
    Surf string
    Pos string
} 

type Surfaces struct {
    []Surface
}

func tokenize(s Master.Sentence)(Surfaces) {
	t := tokenizer.New()
	tokens := t.Analyze(s, tokenizer.Normal)
    var surf Surfaces
    for _, token := range tokens {
        surf = append(surf, &Surface{token.Surface, token.Pos()})
    }
    return surf
}
