package stanfordcorenlp

import (
	"context"
	"encoding/json"
)

// Parse performs constituency parsing and dependency parsing.
func (c *client) Parse(ctx context.Context, text string) (*Document, error) {
	resp, err := c.Do(ctx, text, AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos|AnnotatorParse)
	if err != nil {
		return nil, err
	}

	var d Document
	if err := json.Unmarshal(resp, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
