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

type Surfaces []Surface

func Tokenize(s Master) Surfaces {
	t := tokenizer.New()
	tokens := t.Analyze(s.Sentence, tokenizer.Normal)
    var surf Surfaces
    for _, token := range tokens {
        surf = append(surf, Surface{token.Surface, token.Pos()})
    }
    return surf
}
