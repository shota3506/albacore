package stanfordcorenlp

import (
	"context"
	"encoding/json"
)

// Tokenize turns text into tokens.
// This method calls Stanford CoreNLP API and converts the response into struct.
func (c *client) Tokenize(ctx context.Context, text string) (*Document, error) {
	resp, err := c.post(ctx, text, &properties{
		Annotators:   &annotators{annotatorTokenize, annotatorSsplit, annotatorPos},
		OutputFormat: "json",
	})
	if err != nil {
		return nil, err
	}

	var d Document
	if err := json.Unmarshal(resp, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
