// RPC method is function/procedure that can be invoked on a remote server or service as if it were a local function
package main

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"time"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

// Super-Secret `reversing a string' method we can't allow clients to see.
func ReverseString(s string, i int) string {
	// Simulate a delay by sleeping for a random duration up to 'i' seconds
	time.Sleep(time.Duration(rand.Intn(i)) * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type SecretStringOperations struct{}

func (s *SecretStringOperations) Reverse(req stubs.Request, res *stubs.Response) (err error) {
	// Reverse the string in the request and delay the response by up to 10 seconds
	res.Message = ReverseString(req.Message, 10)
	return
}

func (s *SecretStringOperations) FastReverse(req stubs.Request, res *stubs.Response) (err error) {
	// Reverse the string in the request and respond faster with a delay of up to 2 seconds
	res.Message = ReverseString(req.Message, 2)
	return
}

func main() {
	// Parse command-line arguments to get the port on which the server should listen
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	// Send random generated number with current time
	rand.Seed(time.Now().UnixNano())
	// Register 'SecretStringOperations' struct as an RPC servide
	rpc.Register(&SecretStringOperations{})
	// Create TCP listener on specified port
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	defer listener.Close()
	// Accept incoming RPC requests and sever them. The server keeps running and listening for requests
	rpc.Accept(listener)
}
