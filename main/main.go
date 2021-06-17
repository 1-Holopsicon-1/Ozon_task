package main

import (
	"google.golang.org/grpc"
	"../pkg/grpcserver"
	"../pkg/api"
	"log"
	"net"
)

func main()  {
	s := grpc.NewServer()
	srv := &grpcserver.GRPCSever{}
	RegisterUrlCutterServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}