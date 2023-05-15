package gojieba

import (
	"github.com/andyzhou/tinyidf"
	"testing"
)

func TestDictInit(t *testing.T) {
	tinyidf.NewDictionary("dict/dict.txt")
}

func TestIDFTable(t *testing.T) {
	tinyidf.NewIDFTable("dict/idf.txt")
}
