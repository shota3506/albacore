package stanfordcorenlp

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientDo(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	c := NewClient(context.Background(), url)

	resp, err := c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize,
	)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
