package main

import (
	"context"
	"demothrift/gen-go/demothrift"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

var defaultCtx = context.Background()
func main() {
	var transport thrift.TTransport
	var err error
	transport ,err =thrift.NewTSocket("127.0.0.1:7778")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		log.Fatal(err)
	}
	fmt.Printf("%T\n", transport)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	if err := transport.Open(); err != nil {
		log.Fatal(err,"g")
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := demothrift.NewDemoClient(thrift.NewTStandardClient(iprot, oprot))
	client.HelloWould(defaultCtx)
}