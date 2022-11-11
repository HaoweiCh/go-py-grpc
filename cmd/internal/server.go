package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "haowei.ch/go-py-grpc/py"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func ExecScript(name string, arg ...string) error {

	cmd := exec.CommandContext(context.Background(), name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("script not start: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("script exit with error: %w", err)
	}

	return nil
}
