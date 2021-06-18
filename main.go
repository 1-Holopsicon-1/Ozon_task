package main

import (
	_ "google.golang.org/grpc"
	"../pkg/grpcserver"
	_ "github.com/1-Holopsicon-1/Ozon_task/pkg/api"
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