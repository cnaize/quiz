package words

type WordList []map[byte][]string

func (w *WordList) AddWord(word string) {
	listLen := len(*w)
	wordLen := len(word)
	if listLen < wordLen {
		for i := 0; i < wordLen-listLen; i++ {
			*w = append(*w, make(map[byte][]string))
		}
	}

	wordSlice := (*w)[wordLen-1][word[0]]
	wordSlice = append(wordSlice, word)
	(*w)[wordLen-1][word[0]] = wordSlice
}

func (w WordList) AllWords() []string {
	var res []string
	for _, m := range w {
		for _, s := range m {
			res = append(res, s...)
		}
	}

	return res
}

func (w WordList) SubList(wordLen int) WordList {
	return w[:wordLen]
}

func (w WordList) Words(lenght int, char byte) []string {
	return w[lenght][char]
}
