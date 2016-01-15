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

	var i int
	var found bool
	var res string
	for {
		stats := mypool.Status()
		if !found {
			// add payload depending on "numCPU" and running workers
			for j := 0; j < numCPU-stats.Running; j++ {
				if i < len(allWords) {
					mypool.Add(words.HandleWord, allWords[i], wordList)
					i++
				}
			}
		}

		job := mypool.WaitForJob()
		if job == nil {
			break
		}

		jres, ok := job.Result.(string)
		if !ok {
			panic("job: invalid result type")
		}

		if len(jres) > len(res) {
			res = jres
			found = true
		}

		if found && stats.Completed == stats.Submitted {
			break
		}
	}

	fmt.Printf("Result: %s, len - %d\n", res, len(res))
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
