package main

import (
	"github.com/cnaize/quiz/words"
	"testing"
)

func TestQuiz(t *testing.T) {
	var wordList words.WordList
	wordList.AddWord("a")
	wordList.AddWord("b")
	wordList.AddWord("c")
	wordList.AddWord("df")
	wordList.AddWord("bdfc")
	wordList.AddWord("acdfe")

	res := words.HandleWord("df", wordList).(words.Result)
	if len(res) != 0 {
		t.Errorf("failed: df")
	}

	res = words.HandleWord("acdfe", wordList).(words.Result)
	if len(res) != 0 {
		t.Errorf("failed: acdfe")
	}

	res = words.HandleWord("bdfc", wordList).(words.Result)
	if len(res) != 1 || res[0] != "bdfc" {
		t.Errorf("failed: bdfc")
	}
}
