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

	wordList, err := loadWords(inPath)
	if err != nil {
		panic(fmt.Sprintf("can't load input file: %+v\n", err))
	}

	mypool := pool.New(runtime.NumCPU())
	mypool.Run()

	for _, w := range wordList.AllWords() {
		mypool.Add(words.HandleWord, w, wordList)
	}

	var res words.Result
	for {
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

		if len(res) == 0 {
			res = jres
		} else if len(jres[0]) == len(res[0]) {
			res = append(res, jres...)
		} else if len(jres[0]) > len(res[0]) {
			res = jres
		}
	}

	fmt.Printf("Result: %v\n", res)
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
