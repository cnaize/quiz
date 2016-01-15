package words

import (
	"strings"
)

func HandleWord(args ...interface{}) interface{} {
	word, ok := args[0].(string)
	if !ok {
		panic("invalid word type")
	}
	wordList, ok := args[1].(WordList)
	if !ok {
		panic("invalid word list type")
	}

	return processWord(word, wordList, word)
}

func processWord(subWord string, wordList WordList, word string) string {
	var res string
	if len(subWord) == 0 {
		return word
	}

	subList := wordList.SubList(len(subWord))

	// go through "subList" from min words lenght to max
	for i := 0; i < len(subList); i++ {
		// get all words of "subList" with lenght "i+1" and started with first "subWord"'s char
		wordSlice := subList.Words(i, subWord[0])
		for k := 0; k < len(wordSlice); k++ {
			tmpWord := wordSlice[k]
			if tmpWord == word {
				continue
			}

			if strings.HasPrefix(subWord, tmpWord) {
				var subRes string
				// if "subWord" can be processed, recursively call the func
				if len(subWord) >= len(tmpWord) {
					subRes = processWord(subWord[len(tmpWord):], wordList, word)
				}

				if len(subRes) > len(res) {
					res = subRes
				}
			}
		}
	}

	return res
}
