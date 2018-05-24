package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	pb "github.com/DemoLiang/demo/dailylify/go/gorpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func invoke(c pb.GreeterClient, name string) {
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	_ = r
}

func syncTest(c pb.GreeterClient, name string) {
	i := 10000
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		invoke(c, name)
	}
	fmt.Printf("took:%vms\n", (time.Now().UnixNano()-t)/1000000)
}

func asyncTest(c [20]pb.GreeterClient, name string) {
	var wg sync.WaitGroup
	wg.Add(10000)

	i := 10000
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		go func() { invoke(c[i%20], name); wg.Done() }()
	}
	wg.Wait()
	fmt.Println("took", (time.Now().UnixNano()-t)/1000000, "ms")
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	var c [20]pb.GreeterClient

	name := defaultName
	sync := true
	if len(os.Args) > 1 {
		sync, err = strconv.ParseBool(os.Args[1])
	}

	i := 0
	for ; i < 20; i++ {
		c[i] = pb.NewGreeterClient(conn)
		invoke(c[i], name)
	}
	if sync {
		syncTest(c[0], name)
	} else {
		asyncTest(c, name)
	}
}
