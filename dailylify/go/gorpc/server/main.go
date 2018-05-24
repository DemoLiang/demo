package main

import(
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/DemoLiang/demo/dailylify/go/gorpc/protos"
)

const(
	port= ":50051"
)


type server struct{
}

func (s *server)SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloReply,error){
	return &pb.HelloReply{
		Message:"hello"+in.Name,
	},nil
}

func main(){
	lis,err := net.Listen("tcp",port)
	if err!= nil{
		log.Fatalf("failed to listen:%v",err)
	}
	s:= grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	s.Serve(lis)
}
