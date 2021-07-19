package stanfordcorenlp

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientDo(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	c := NewClient(context.Background(), url)

	resp, err := c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		&Properties{
			Annotators:   &Annotators{AnnotatorTokenize},
			OutputFormat: "json",
		},
	)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestProperties(t *testing.T) {
	for _, testcase := range []struct {
		properties *Properties
		expected   string
	}{
		{
			&Properties{Annotators: &Annotators{AnnotatorTokenize}},
			`{"annotators":"tokenize"}`,
		},
		{
			&Properties{Annotators: &Annotators{AnnotatorTokenize}, OutputFormat: "json"},
			`{"annotators":"tokenize","outputFormat":"json"}`,
		},
		{
			&Properties{Annotators: &Annotators{AnnotatorTokenize, AnnotatorSsplit}, OutputFormat: "json"},
			`{"annotators":"tokenize,ssplit","outputFormat":"json"}`,
		},
	} {
		p, err := json.Marshal(testcase.properties)
		require.NoError(t, err)
		assert.Equal(t, testcase.expected, string(p))
	}
}
