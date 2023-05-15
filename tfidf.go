package tinyidf

import (
	"fmt"
	"github.com/andyzhou/tinyidf/internal/common"
	"sort"
	"strings"
)

var (
	defaultStopWords = []string{
		"the", "of", "is", "and", "to", "in", "that", "we", "for", "an", "are",
		"by", "be", "as", "on", "with", "can", "if", "from", "which", "you", "it",
		"this", "then", "at", "have", "all", "not", "one", "has", "or", "that",
	}
)

// TF-IDF is the core struct of this file.
type TfIdf struct {
	tokenizer *Tokenizer
	idf       *IDFTable
	stopWords common.StringSet
}

// NewTfIdf create an instance of TF-IDF with the given tokenizer and idf file.
func NewTfIdf(t *Tokenizer, idf string) *TfIdf {
	tfIdf := &TfIdf{
		tokenizer: t,
		idf:       NewIDFTable(idf),
		stopWords: common.NewStringSet(),
	}
	for _, sw := range defaultStopWords {
		tfIdf.AddStopWord(sw)
	}
	return tfIdf
}

func (t *TfIdf) String() string {
	return fmt.Sprintf("TFIDF(idf='%s', tokenizer=%v)", t.idf.GetFile(), t.tokenizer)
}

// AddStopWord add a stop word to stopWords set.
func (t *TfIdf) AddStopWord(word string) bool {
	if word == "" {
		return false
	}
	return t.stopWords.Add(strings.ToLower(word))
}

// Extract extract topK keywords from the given sentence.
func (t *TfIdf) Extract(sentence string, topK int) []Keyword {
	totalWords := float64(0)
	wordFreqMap := make(map[string]float64)

	for _, word := range t.tokenizer.Cut(sentence, true) {
		r := []rune(strings.TrimSpace(word))
		if len(r) < 2 || t.stopWords.Has(strings.ToLower(word)) {
			continue
		}
		wordFreqMap[word]++
		totalWords++
	}

	for word := range wordFreqMap {
		freq, ok := t.idf.GetFreq(word)
		if !ok {
			freq = t.idf.GetMedian()
		}
		wordFreqMap[word] *= freq / totalWords
	}

	kws := make([]Keyword, 0, len(wordFreqMap))
	for word, freq := range wordFreqMap {
		kws = append(kws, Keyword{word, freq})
	}

	sort.Slice(kws, func(i, j int) bool {
		return kws[i].score > kws[j].score
	})
	return kws[:topK]
}

