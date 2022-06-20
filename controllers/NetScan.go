package controllers

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	Colors "github.com/muhammadqazi/Mr.Hack3r/utils"
)

func NetScan(ip string) {

	subnetToScan := ip

	activeThreads := 0
	doneChannel := make(chan bool)

	for ip := 0; ip <= 255; ip++ {
		fullIP := subnetToScan + "." + strconv.Itoa(ip)
		go resolve(fullIP, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func resolve(ip string, doneChannel chan bool) {
	addresses, err := net.LookupAddr(ip)
	if err == nil {
		fmt.Println(Colors.Blue(), "[+] Host ", Colors.White(), ip, "\t:", Colors.Red(), strings.Join(addresses, ", "))
	}
	doneChannel <- true
}
