package main

import (
	"net"
	"google.golang.org/grpc/reflection"
	"flag"
	"addorg/proto/server"
	"addorg/rpc"
	"google.golang.org/grpc"
	"os"
	"log"
)

func main(){
	// Defining flags
	listen := flag.String("address",":8080","Address of server")
	dev := flag.Bool("dev",true,"To setup reflection")
	flag.Parse()

	log := log.New(os.Stdout,"Build_server",log.LstdFlags)

	grpcServer := grpc.NewServer()

	rpcServer := rpc.NewOrgServer(log)

	server.RegisterNewOrgServer(grpcServer,rpcServer)

	if *dev{
		reflection.Register(grpcServer)
	}

	l, err := net.Listen("tcp",*listen)
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(l)
}