package main

import (
	"bufio"
	"fmt"
	"github.com/jonfk/battleship/protocol"
	"io"
	"log"
	"net"
	"os"
)

const (
	DEFAULT_CONN_HOST = "0.0.0.0"
	DEFAULT_CONN_PORT = "8888"
	CONN_TYPE         = "tcp"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", DEFAULT_CONN_HOST+":"+DEFAULT_CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to server through tcp.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go printOutput(conn)
	writeInput(conn)
}

func printOutput(conn *net.TCPConn) {

	for {

		msg, err := protocol.ReadMsg(conn)
		// Receiving EOF means that the connection has been closed
		if err == io.EOF {
			// Close conn and exit
			conn.Close()
			fmt.Println("Connection Closed. Bye bye.")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}

func writeInput(conn *net.TCPConn) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(text)
		err = protocol.WriteMsg(conn, protocol.ConnectMsg{Username: "jonfk"})
		if err != nil {
			log.Println(err)
		}
	}
}
