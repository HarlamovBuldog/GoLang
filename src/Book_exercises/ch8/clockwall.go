// Clockwall.go is TCP-client for handling data
// from connections to multiple servers.
// Exercise 8.1 realization from book page 265
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type TimeTable struct {
	curTime []byte
}

func (tt *TimeTable) Write(data []byte) (n int, err error) {
	tt.curTime = data
	return len(data), nil
}

func (tt TimeTable) String() string {
	return string(tt.curTime)
}

func main() {
	var areaNames []string
	var localhosts []string
	var validArg = regexp.MustCompile(`[a-zA-z]+=localhost:(6553[0-5]|655[0-2][0-9]\d|65[0-4](\d){2}|6[0-4](\d){3}|[1-5](\d){4}|[1-9](\d){0,3})`)
	for _, argToParse := range os.Args[1:] {
		if validArg.MatchString(argToParse) == true {
			// go func here
			indexOfEqualSign := strings.Index(argToParse, "=")

			iterAreaName := argToParse[:indexOfEqualSign]
			areaNames = append(areaNames, iterAreaName)

			iterLocalhost := argToParse[indexOfEqualSign+1:]
			localhosts = append(localhosts, iterLocalhost)

			//fmt.Println("Area: %s\nHost: %s\n", iterAreaName, iterLocalhost)
		} else {
			fmt.Printf("%s: invalid call\n", argToParse)
		}
	}

	for {
		for _, areaName := range areaNames {
			fmt.Print(areaName + "\t")
		}
		fmt.Println("\n----------------------------------------------")

		for _, localhost := range localhosts {
			conn, err := net.Dial("tcp", localhost)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			timeTable := TimeTable{}
			mustCopy(&timeTable, conn)
			fmt.Printf("%s\t", timeTable)
		}
		fmt.Println("")
		time.Sleep(1 * time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
