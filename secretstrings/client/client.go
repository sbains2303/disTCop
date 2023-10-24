package main

import (
	"flag" // for command line arguments
	"fmt"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	// Parse CLI to get IP and port of server to connect to
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	// Dial RPC server using provided server address
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	dat, err := os.ReadFile("../wordlist")
	if err != nil {
		panic(err)
	}

	request := stubs.Request{Message: string(dat)}
	// Initialise response structure
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)
}
