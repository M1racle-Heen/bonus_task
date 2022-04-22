package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	var reply string
	client, err := rpc.DialHTTP("tcp", "localhost:9000")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	text, _ := reader.ReadString('\n')

	client.Call("API.Welcome", text, &reply)
	fmt.Println(reply)
}
