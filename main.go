package main

import (
	"fmt"
	"net"
	"os"

	"github.com/namsral/flag"

	pb "github.com/iochti/thing-group-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Db is a PostgresDL instance used by the service
var Db *MgoDL

func init() {
	Db = new(MgoDL)
	Db.Init()
}

func dieIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s. Try --help for help.\n", err)
	os.Exit(-1)
}

func main() {
	addr := flag.String("srv", ":5003", "TCP address to listen on (in host:port form)")
	certFile := flag.String("cert", "", "Path to PEM-encoded certificate")
	keyFile := flag.String("key", "", "Path to PEM-encoded secret key")
	flag.Parse()
	svc := &ThingGroupSvc{
		Db: Db,
	}
	var server *grpc.Server

	// Create server, with TLS if cert & key are specified
	if *keyFile == "" && *certFile == "" {
		server = grpc.NewServer()
	} else if *keyFile == "" {
		dieIf(fmt.Errorf("Cert specified with no key"))
	} else if *certFile == "" {
		dieIf(fmt.Errorf("Key specified with no cert"))
	} else {
		pair, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		dieIf(err)
		creds := grpc.Creds(pair)
		server = grpc.NewServer(creds)
	}
	lis, err := net.Listen("tcp", *addr)
	dieIf(err)
	pb.RegisterThingGroupSvcServer(server, svc)
	server.Serve(lis)
}
