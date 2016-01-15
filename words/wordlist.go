package words

type WordList []map[byte][]string

// add word to list
func (wl *WordList) AddWord(word string) {
	listLen := len(*wl)
	wordLen := len(word)
	if listLen < wordLen {
		for i := 0; i < wordLen-listLen; i++ {
			*wl = append(*wl, make(map[byte][]string))
		}
	}

	wordSlice := (*wl)[wordLen-1][word[0]]
	found := false
	for _, w := range wordSlice {
		if word == w {
			found = true
			break
		}
	}
	if !found {
		wordSlice = append(wordSlice, word)
		(*wl)[wordLen-1][word[0]] = wordSlice
	}
}

// get all words (longest first)
func (wl WordList) AllWords() []string {
	var res []string
	for i := len(wl) - 1; i >= 0; i-- {
		for _, s := range wl[i] {
			res = append(res, s...)
		}
	}

	return res
}

// get sub list of words with "lenght < 'wordLen'"
func (wl WordList) SubList(wordLen int) WordList {
	return wl[:wordLen]
}

// get words with specified "lenght" and started with "char"
func (wl WordList) Words(lenght int, char byte) []string {
	return wl[lenght][char]
}
