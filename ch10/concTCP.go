package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn, myCount int) {
	
	fmt.Print(".")
	defer c.Close()

	for {
		
	netData, err := bufio.NewReader(c).ReadString('\n')  // FIX: Read data INSIDE the loop so it waits for new input each time
	if err != nil {
		fmt.Println(err)
		return
	}
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)

		counter := "Client number: " + strconv.Itoa(myCount) + "\n"
		c.Write([]byte(string(counter)))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(5)
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, count)
		count++
	}
}
