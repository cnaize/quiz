package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cnaize/quiz/words"
	"github.com/stefantalpalaru/pool"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Running..")

	var inPath string
	flag.StringVar(&inPath, "in", "word.list", "path to file with words")
	flag.Parse()

	// load words
	wordList, err := loadWords(inPath)
	if err != nil {
		panic(fmt.Sprintf("can't load input file: %+v\n", err))
	}

	numCPU := runtime.NumCPU()

	// create pool
	mypool := pool.New(numCPU)
	mypool.Run()

	allWords := wordList.AllWords()

	// add payload
	for _, w := range allWords {
		mypool.Add(words.HandleWord, w, wordList)
	}

	var i int
	var found bool
	var res words.Result
	for {
		if !found {
			// add payload depending on "numCPU" and running workers
			for j := 0; j < numCPU-mypool.Status().Running; j++ {
				mypool.Add(words.HandleWord, allWords[i], wordList)
				i++
			}
		}

		job := mypool.WaitForJob()
		if job == nil {
			break
		}

		jres, ok := job.Result.(words.Result)
		if !ok {
			panic("job: invalid result type")
		}

		if len(jres) == 0 {
			continue
		}

		if len(res) == 0 || len(jres[0]) > len(res[0]) {
			res = jres
			found = true
		} else if len(jres[0]) == len(res[0]) {
			res = append(res, jres...)
		}

		if mypool.Status().Running == 0 {
			break
		}
	}

	lenght := 0
	if len(res) > 0 {
		lenght = len(res[0])
	}

	fmt.Printf("Results: %v, len - %d\n", res, lenght)
}

func loadWords(path string) (words.WordList, error) {
	var wordList words.WordList

	inFile, err := os.Open(path)
	if err != nil {
		return wordList, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		wordList.AddWord(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return wordList, err
	}

	return wordList, nil
}
