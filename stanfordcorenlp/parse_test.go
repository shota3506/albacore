package stanfordcorenlp

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientParse(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	c := NewClient(context.Background(), url)

	resp, err := c.Parse(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
	)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.Sentences), 1)

	sentence := resp.Sentences[0]
	assert.NotZero(t, sentence.Parse)

	require.GreaterOrEqual(t, len(sentence.BasicDependencies), 1)
	assert.NotZero(t, sentence.BasicDependencies[0].Dep)
	assert.NotZero(t, sentence.BasicDependencies[0].GovernorGloss)
	assert.NotZero(t, sentence.BasicDependencies[0].DependentGloss)

	require.GreaterOrEqual(t, len(sentence.EnhancedDependencies), 1)
	assert.NotZero(t, sentence.EnhancedDependencies[0].Dep)
	assert.NotZero(t, sentence.EnhancedDependencies[0].GovernorGloss)
	assert.NotZero(t, sentence.EnhancedDependencies[0].DependentGloss)

	require.GreaterOrEqual(t, len(sentence.EnhancedPlusPlusDependencies), 1)
	assert.NotZero(t, sentence.EnhancedPlusPlusDependencies[0].Dep)
	assert.NotZero(t, sentence.EnhancedPlusPlusDependencies[0].GovernorGloss)
	assert.NotZero(t, sentence.EnhancedPlusPlusDependencies[0].DependentGloss)
}