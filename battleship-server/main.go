package main

import ()

func main() {
	server := Server{Host: DEFAULT_CONN_HOST, Port: DEFAULT_CONN_PORT}
	server.Run()
}
