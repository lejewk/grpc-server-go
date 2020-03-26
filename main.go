package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-server-go/config"
    "log"
    "net"
)

type server struct {
    pb.UnimplementedConfigStoreServer
}

func (s *server) Get(ctx context.Context, in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
    log.Printf("Received profile: %v", in.GetProfile())
    return &pb.ConfigResponse{JsonConfig: `"{"main":"http://google.com"}"`}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterConfigStoreServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}