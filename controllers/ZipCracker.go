package controllers

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/alexmullins/zip"
)

func ZipCracker(zipfile string, dictonary string, concurrency int) {

	word := make(chan string, 0)
	found := make(chan string, 0)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go zipCrackWorker(word, found, &wg, zipfile)
	}

	dictFile, err := os.Open(dictonary)
	if err != nil {
		log.Fatalln(err)
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	go func() {
		for scanner.Scan() {
			pass := scanner.Text()
			word <- pass
		}
		close(word)
	}()

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case f := <-found:
		println("\n[+] Found password")
		println("[+] Password = ", f, "\n")
		return
	case <-done:
		println("[+] Password not found")
		return
	}

}

func zipCrackWorker(word <-chan string, found chan<- string, wg *sync.WaitGroup, zipfile string) {

	zipr, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}
	defer zipr.Close()
	defer wg.Done()
	for w := range word {
		for _, z := range zipr.File {
			z.SetPassword(w)
			_, err := z.Open()

			if err == nil {
				found <- w
			}
		}
	}
}
