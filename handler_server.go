package main

import (
	"context"
	"fmt"

	//"demo_thrift/gen-go/demothrift/demo_thrift"
	"demothrift/gen-go/demothrift"
)

type DemoHandler struct {
	log map[int]*demothrift.Demo
}

func (d DemoHandler) HelloWould(ctx context.Context) (err error) {
	fmt.Println("HELLO WORLD!!")
	return nil
}

func NewManagerStudentHandler() *DemoHandler {
	return &DemoHandler{log: make(map[int]*demothrift.Demo)}
}



