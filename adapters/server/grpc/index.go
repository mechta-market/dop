package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type St struct {
	Server *grpc.Server
	eChan  chan error
}

func New(opts ...grpc.ServerOption) *St {
	return &St{
		Server: grpc.NewServer(opts...),
		eChan:  make(chan error, 1),
	}
}

func (s *St) Start(port string) (string, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return "", fmt.Errorf("fail to listen: %w", err)
	}

	go func() {
		err = s.Server.Serve(lis)
		if err != nil {
			s.eChan <- err
		}
	}()

	return lis.Addr().String(), nil
}

func (s *St) Wait() <-chan error {
	return s.eChan
}

func (s *St) Shutdown() error {
	defer close(s.eChan)

	s.Server.GracefulStop()

	return nil
}
