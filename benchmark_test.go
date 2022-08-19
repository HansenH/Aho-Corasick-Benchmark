package benchmark

import (
	"bufio"
	"log"
	"os"
	"test/ikmp"
	"test/iregexp"
	"testing"

	ahocorasick "github.com/HansenH/Aho-Corasick"
)

func BenchmarkAC(b *testing.B) {
	dict := []string{}
	dictFile, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}
	book, err := os.Open("Pride and Prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		book.Seek(0, 0)
		ac := ahocorasick.NewACAutomaton(dict)
		scanner = bufio.NewScanner(book)
		for scanner.Scan() {
			ac.FindAllIndex(scanner.Text())
		}
	}
}

func BenchmarkIKMP(b *testing.B) {
	dict := []string{}
	dictFile, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}
	book, err := os.Open("Pride and Prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		book.Seek(0, 0)
		ikmp := ikmp.NewIKMP(dict)
		scanner := bufio.NewScanner(book)
		for scanner.Scan() {
			ikmp.FindAllIndex(scanner.Text())
		}
	}
}

func BenchmarkIRegexp(b *testing.B) {
	dict := []string{}
	dictFile, err := os.Open("dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}
	book, err := os.Open("Pride and Prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		book.Seek(0, 0)
		ire := iregexp.NewIRegexp(dict)
		scanner := bufio.NewScanner(book)
		for scanner.Scan() {
			ire.FindAllIndex(scanner.Text())
		}
	}
}
