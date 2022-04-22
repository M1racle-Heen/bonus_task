package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

func (a *API) Welcome(title string, reply *string) error {
	*reply = "Welcome " + title
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Your registering API is incorrect: ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("You have listener error: ", err)
	}
	log.Printf("Raised the server on port: %d", 9000)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("You have serve error: ", err)
	}
}
