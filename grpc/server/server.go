package main

import (
	"context"
	"dailyPractice/grpc/server/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) Ask(ctx context.Context, req *proto.AskReq) (*proto.AskResp, error) {
	var res proto.AskResp
	if len(req.Question) > 0 {
		res.Answer = "yes"
	}
	return &res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterTestServer(srv, &server{})
	if err = srv.Serve(listen); err != nil {
		panic(err)
	}
}
