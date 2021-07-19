package stanfordcorenlp

import "strings"

type Token struct {
	Index                int    `json:"index"`
	Word                 string `json:"word"`
	OriginalText         string `json:"originalText,omitempty"`
	CharacterOffsetBegin int    `json:"characterOffsetBegin,omitempty"`
	CharacterOffsetEnd   int    `json:"characterOffsetEnd,omitempty"`
	Pos                  string `json:"pos,omitempty"`
	Before               string `json:"before,omitempty"`
	After                string `json:"after,omitempty"`
}

func (t *Token) String() string {
	return t.Word
}

type Sentence struct {
	Index  int      `json:"index"`
	Tokens []*Token `json:"tokens"`
}

func (s *Sentence) String() string {
	strs := make([]string, len(s.Tokens))
	for i, t := range s.Tokens {
		strs[i] = t.String()
	}
	return strings.Join(strs, " ")
}

type Document struct {
	Sentences []*Sentence `json:"sentences"`
}

func (d *Document) String() string {
	strs := make([]string, len(d.Sentences))
	for i, s := range d.Sentences {
		strs[i] = s.String()
	}
	return strings.Join(strs, " ")
}
