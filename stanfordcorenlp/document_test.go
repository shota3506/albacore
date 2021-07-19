package stanfordcorenlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	token := &Token{
		Index: 0,
		Word:  "word",
	}

	assert.Equal(t, "word", token.String())
}

func TestSentece(t *testing.T) {
	sentence := &Sentence{
		Index: 0,
		Tokens: []*Token{
			{Index: 0, Word: "The"},
			{Index: 1, Word: "quick"},
			{Index: 2, Word: "brown"},
			{Index: 3, Word: "fox"},
			{Index: 4, Word: "jumped"},
			{Index: 5, Word: "over"},
			{Index: 6, Word: "the"},
			{Index: 7, Word: "lazy"},
			{Index: 8, Word: "dog"},
			{Index: 9, Word: "."},
		},
	}

	assert.Equal(t, "The quick brown fox jumped over the lazy dog .", sentence.String())
}

func TestDocument(t *testing.T) {
	document := &Document{
		Sentences: []*Sentence{
			{
				Index: 0,
				Tokens: []*Token{
					{Index: 0, Word: "Hello"},
					{Index: 1, Word: "world"},
					{Index: 2, Word: "."},
				},
			},
			{
				Index: 1,
				Tokens: []*Token{
					{Index: 0, Word: "Hello"},
					{Index: 1, Word: "world"},
					{Index: 1, Word: "again"},
					{Index: 2, Word: "."},
				},
			},
		},
	}
	assert.Equal(t, "Hello world . Hello world again .", document.String())
}
