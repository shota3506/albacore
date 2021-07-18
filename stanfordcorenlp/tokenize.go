package stanfordcorenlp

import (
	"context"
	"encoding/json"
)

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

type Sentence struct {
	Index  int      `json:"index"`
	Tokens []*Token `json:"tokens"`
}

type Response struct {
	Sentences []*Sentence `json:"sentences"`
}

// Tokenize turns text into tokens.
// This method calls Stanford CoreNLP API and converts the response into struct.
func (c *client) Tokenize(ctx context.Context, text string) (*Response, error) {
	resp, err := c.post(ctx, text, &properties{
		Annotators:   "tokenize,ssplit,pos",
		OutputFormat: "json",
	})
	if err != nil {
		return nil, err
	}

	var r Response
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
