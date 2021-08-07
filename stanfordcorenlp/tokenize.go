package stanfordcorenlp

import (
	"context"
	"encoding/json"
)

// Tokenize turns text into tokens.
func (c *client) Tokenize(ctx context.Context, text string) (*Document, error) {
	resp, err := c.Do(ctx, text, AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos)
	if err != nil {
		return nil, err
	}

	var d Document
	if err := json.Unmarshal(resp, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
