package words

import (
	"strings"
)

type Result []string

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

func processWord(subWord string, wordList WordList, word string) Result {
	var res Result
	if len(subWord) == 0 {
		res = []string{word}
		return res
	}

	subList := wordList.SubList(len(subWord))
	var found bool
	// go through "subWord"'s characters
	for i, c := range subWord {
		// go through "subList" from min words lenght to max
		for j := 0; j < len(subList); j++ {
			// get all words of "subList" with lenght "j+1" and started with "c"
			wordSlice := subList.Words(j, byte(c))
			for k := 0; k < len(wordSlice); k++ {
				tmpWord := wordSlice[k]
				if len(tmpWord) > len(subWord) || tmpWord == word {
					continue
				}

				if strings.HasPrefix(subWord, tmpWord) {
					var subRes Result
					// if "subWord" can be processed recursively call the func
					if len(subWord) >= i+len(tmpWord) {
						subRes = processWord(subWord[i+len(wordSlice[k]):], wordList, word)
					}

					rWordLen := 0
					srWordLen := 0
					if len(res) > 0 {
						rWordLen = len(res[0])
					}
					if len(subRes) > 0 {
						srWordLen = len(subRes[0])
					}

					if len(res) == 0 || srWordLen > rWordLen {
						res = subRes
					} else if srWordLen == rWordLen {
						for _, w := range res {
							for _, sw := range subRes {
								if sw != w {
									res = append(res, sw)
								}
							}
						}
					}

					found = true
				}
			}
		}

		if !found {
			break
		}
	}

	return res
}
