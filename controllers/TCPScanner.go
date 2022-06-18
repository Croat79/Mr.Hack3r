package controllers

import (
	"fmt"
	"net"
	"strconv"
	"time"

	Colors "github.com/muhammadqazi/Mr.Hack3r/utils"
)

func TCPScanner(target string) {

	activeThreads := 0
	doneChannel := make(chan bool)

	for port := 0; port <= 65535; port++ {
		go testTCPConnection(target, port, doneChannel)
		activeThreads++
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port),
		time.Second*10)
	if err == nil {
		fmt.Println(Colors.Blue(), "[+] Port ", Colors.White(), port, "\t:", Colors.Red(), "Open")
	}
	doneChannel <- true
}
