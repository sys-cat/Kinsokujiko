package kinsokujiko

import (
	"errors"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/sys-cat/kinsokujiko/targets"
)

// Master is Master data for analyze
type Master struct {
	Sentence string
}

// Surface is surface, pos pair
type Surface struct {
	Surf string
	Pos  string
}

// Surfaces is Slice any Surface
type Surfaces []Surface

// Run is Masking Sentence what use Tokenize method.
func Run(s Master, t targets.Targets) (string, error) {
	return "", errors.New("anything is bad")
}

// Tokenize is analyze sentence method
func Tokenize(s Master) Surfaces {
	t := tokenizer.New()
	tokens := t.Analyze(s.Sentence, tokenizer.Normal)
	var surf Surfaces
	for _, token := range tokens {
		if token.Pos() != "" {
			surf = append(surf, Surface{token.Surface, token.Pos()})
		}
	}
	return surf
}

// AddDictionary is Create User Dictionary
func AddDictionary(dic Dictionary) (bool, error) {
	if err := insertItem(dic); err != nil {
		return false, err
	}
	return true, errors.New("")
}

func insertItem(dic Dictionary) error {
	// anything to Create User Dictionary
	return errors.New("anything is bad")
}
