package stanfordcorenlp

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite

	c Client
}

func TestClient(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}

func (s *ClientTestSuite) SetupSuite() {
	url := os.Getenv("STANFORD_CORENLP_URL")
	s.c = NewClient(context.Background(), url)
}

func (s *ClientTestSuite) TestDoTokenize() {
	resp, err := s.c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize,
	)
	s.Require().NoError(err)

	var sen Sentence
	err = json.Unmarshal(resp, &sen)
	s.Require().NoError(err)

	s.Require().NotEmpty(len(sen.Tokens))
	s.NotZero(sen.Tokens[0].Word)
}

func (s *ClientTestSuite) TestDoPos() {
	resp, err := s.c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos,
	)
	s.Require().NoError(err)

	var doc Document
	err = json.Unmarshal(resp, &doc)
	s.Require().NoError(err)

	s.Require().NotEmpty(len(doc.Sentences))
	s.Require().NotEmpty(len(doc.Sentences[0].Tokens))
	s.NotZero(doc.Sentences[0].Tokens[0].Word)
}

func (s *ClientTestSuite) TestDoLemma() {
	resp, err := s.c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos|AnnotatorLemma,
	)
	s.Require().NoError(err)

	var doc Document
	err = json.Unmarshal(resp, &doc)
	s.Require().NoError(err)

	s.Require().NotEmpty(len(doc.Sentences))
	s.Require().NotEmpty(len(doc.Sentences[0].Tokens))
	s.NotZero(doc.Sentences[0].Tokens[0].Lemma)
}

func (s *ClientTestSuite) TestDoParse() {
	resp, err := s.c.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos|AnnotatorParse,
	)
	s.Require().NoError(err)

	var doc Document
	err = json.Unmarshal(resp, &doc)
	s.Require().NoError(err)

	s.Require().GreaterOrEqual(len(doc.Sentences), 1)

	sentence := doc.Sentences[0]
	s.NotZero(sentence.Parse)

	s.Require().NotEmpty(len(sentence.BasicDependencies))
	s.NotZero(sentence.BasicDependencies[0].Dep)
	s.NotZero(sentence.BasicDependencies[0].GovernorGloss)
	s.NotZero(sentence.BasicDependencies[0].DependentGloss)

	s.Require().NotEmpty(len(sentence.EnhancedDependencies))
	s.NotZero(sentence.EnhancedDependencies[0].Dep)
	s.NotZero(sentence.EnhancedDependencies[0].GovernorGloss)
	s.NotZero(sentence.EnhancedDependencies[0].DependentGloss)

	s.Require().NotEmpty(len(sentence.EnhancedPlusPlusDependencies))
	s.NotZero(sentence.EnhancedPlusPlusDependencies[0].Dep)
	s.NotZero(sentence.EnhancedPlusPlusDependencies[0].GovernorGloss)
	s.NotZero(sentence.EnhancedPlusPlusDependencies[0].DependentGloss)
}
