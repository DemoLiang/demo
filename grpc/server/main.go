package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/DemoLiang/demo/grpc/pb"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server)SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloReply,error){
	return &pb.HelloReply{Message:"hello "+ in.Name},nil
}

func main(){
	lis,err:=net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("failed to listen:%v",err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)
	if err :=s.Serve(lis);err != nil{
		log.Fatalf("failed to serve:%v",err)
	}
}
