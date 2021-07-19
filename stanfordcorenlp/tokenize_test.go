package stanfordcorenlp

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientTokenize(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	c := NewClient(context.Background(), url)

	resp, err := c.Tokenize(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
	)
	require.NoError(t, err)
	require.NotEmpty(t, len(resp.Sentences))
	require.NotEmpty(t, len(resp.Sentences[0].Tokens))
	assert.NotZero(t, resp.Sentences[0].Tokens[0].Word)
}
