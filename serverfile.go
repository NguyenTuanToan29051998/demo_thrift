package main

import (
	//"demo_thrift/gen-go/demothrift/demo_thrift"
	"demothrift/gen-go/demothrift"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

func main() {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket("127.0.0.1:7778")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", transport)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()
	handler := NewManagerStudentHandler()
	processor := demothrift.NewDemoProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", "127.0.0.1:7778")
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}


