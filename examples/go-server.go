package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/prashanthpai/sunrpc"
)

func main() {
	server := rpc.NewServer()
	arith := new(Arith)
	server.Register(arith)

	programNumber := uint32(12345)
	programVersion := uint32(1)

	// TODO: Automate this by parsing the .x file ?
	_ = sunrpc.RegisterProcedure(sunrpc.ProcedureID{programNumber, programVersion, uint32(1)}, "Arith.Add")
	_ = sunrpc.RegisterProcedure(sunrpc.ProcedureID{programNumber, programVersion, uint32(2)}, "Arith.Multiply")

	sunrpc.DumpProcedureRegistry()

	// TODO: Get port from portmapper
	listener, err := net.Listen("tcp", "127.0.0.1:41707")
	if err != nil {
		log.Fatal("net.Listen() failed: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener.Accept() failed: ", err)
		}
		go server.ServeCodec(sunrpc.NewServerCodec(conn))
	}
}
