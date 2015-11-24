package main

import (
	//"fmt"
	"github.com/jonfk/battleship/protocol"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/boltdb/bolt"
)

const (
	DEFAULT_CONN_HOST = "0.0.0.0"
	DEFAULT_CONN_PORT = "8888"
	CONN_TYPE         = "tcp"
	DEFAULT_DB_FILE   = "~/.battleship/battleship.db"
)

type Server struct {
	Host           string
	Port           string
	connections    []net.Conn
	connectionsMut sync.Mutex
	BoltDBFile     string
	boltdb         *bolt.DB
}

func (server Server) Run() {
	l, err := net.Listen(CONN_TYPE, server.Host+":"+server.Port)
	if err != nil {
		log.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	if server.BoltDBFile == "" {
		server.BoltDBFile = DEFAULT_DB_FILE
	}
	server.boltdb, err = bolt.Open(server.BoltDBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer server.boltdb.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		log.Printf("Accepting new connection from %v\n", conn.RemoteAddr())
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Save connection
		server.connectionsMut.Lock()
		server.connections = append(server.connections, conn)
		server.connectionsMut.Unlock()
		// Handle connections in a new goroutine.
		go server.handleRequest(conn)
	}

}

func (server Server) handleRequest(conn net.Conn) {
	for {
		msg, err := protocol.ReadMsg(conn)
		if err != nil {
			if err == io.EOF {
				// Close the connection when you're done with it.
				server.removeConn(conn)
				conn.Close()
				return
			}
			log.Println(err)
			return
		}
		log.Printf("Message Received: %#v\n", msg)
		switch msg.(type) {
		case protocol.PingMsg:
		case protocol.OkMsg:
		case protocol.ErrorMsg:
		case protocol.GameMoveMsg:
		case protocol.ChatMessageMsg:
		case protocol.ConnectMsg:
		case protocol.RequestOpenGamesListMsg:
		case protocol.CreateGameMsg:
		case protocol.JoinGameMsg:
		case protocol.AcceptGameMsg:
		case protocol.RejectGameMsg:
		case protocol.GameSetPieceMsg:
		case protocol.RequestGameStateMsg:
		case protocol.AbandonGameMsg:
		case protocol.OpenGamesListMsg:
		case protocol.GamePreGameStatusMsg:
		case protocol.GameStateMsg:
		case protocol.GameWonMsg:
		case protocol.GameLostMsg:
		default:
		}
		//broadcast(conn, msg)
	}
}

func (server Server) removeConn(conn net.Conn) {
	server.connectionsMut.Lock()
	var i int
	for i = range server.connections {
		if server.connections[i] == conn {
			break
		}
	}
	server.connections = append(server.connections[:i], server.connections[i+1:]...)
	server.connectionsMut.Unlock()
}
