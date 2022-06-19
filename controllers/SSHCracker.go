package controllers

import (
	"bufio"
	"fmt"
	"os"
)

func DictonaryBuilder(passDict string) (slicePass []string) {

	passDictFile, _ := os.Open(passDict)
	defer passDictFile.Close()
	scannerP := bufio.NewScanner(passDictFile)
	scannerP.Split(bufio.ScanLines)
	for scannerP.Scan() {
		slicePass = append(slicePass, scannerP.Text())
	}

	return slicePass
}

func SSHCracker(host, port, user, pass string) {

	fmt.Println(DictonaryBuilder(pass))
}
