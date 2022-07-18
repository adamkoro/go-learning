package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

var PORT = 8080

func main() {
	fmt.Println("Basic grpc server")
	port := convToString(PORT)
	startServer(port)

}

func convToString(num int) string {
	return strconv.Itoa(num)
}

func startServer(PORT string) {
	srv, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(srv); err != nil {
		log.Fatal(err)
	}
}
