package stanfordcorenlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnnotatorType_String(t *testing.T) {
	for _, testcase := range []struct {
		annotators AnnotatorType
		expected   string
	}{
		{AnnotatorTokenize, "tokenize"},
		{AnnotatorQuote, "quote"},
		{AnnotatorTokenize | AnnotatorSsplit | AnnotatorPos, "tokenize,ssplit,pos"},
		{0, ""},
		{1 << 30, ""},
	} {
		assert.Equal(t, testcase.expected, testcase.annotators.String())
	}
}

func TestAnnotatorType_MarshalJSON(t *testing.T) {
	for _, testcase := range []struct {
		annotators AnnotatorType
		expected   []byte
	}{
		{AnnotatorTokenize, []byte(`"tokenize"`)},
		{AnnotatorQuote, []byte(`"quote"`)},
		{AnnotatorTokenize | AnnotatorSsplit | AnnotatorPos, []byte(`"tokenize,ssplit,pos"`)},
		{0, []byte(`""`)},
		{1 << 30, []byte(`""`)},
	} {
		b, err := testcase.annotators.MarshalJSON()
		require.NoError(t, err)
		assert.Equal(t, testcase.expected, b)
	}
}
