package server

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    proto "github.com/radiopapus/bitmex-grpc-proto"

    "bitmex-grpc-proxy/config"
)

type server struct {
    proto.UnimplementedEchoServiceServer
}

func (s *server) Ping(
    ctx context.Context,
    req *proto.PingRequest,
) (*proto.PingResponse, error) {
    return &proto.PingResponse{Reply: "Pong"}, nil
}

// Run grpc server
func Run(config config.Config) {

    lis, err := net.Listen("tcp", config.GrpcPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterEchoServiceServer(s, &server{})
    reflection.Register(s)

    log.Println("Server started on " + config.GrpcPort)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
