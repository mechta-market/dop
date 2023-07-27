package grpc

import (
	"fmt"
	"net"

	"github.com/mechta-market/dop/adapters/logger"
	"google.golang.org/grpc"
)

type Server struct {
	lg logger.Lite

	Server *grpc.Server
	eChan  chan error
}

func New(lg logger.Lite, opts ...grpc.ServerOption) *Server {
	return &Server{
		lg: lg,

		Server: grpc.NewServer(opts...),
		eChan:  make(chan error, 1),
	}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("fail to listen: %w", err)
	}

	s.lg.Infow("Start grpc server", "addr", lis.Addr().String())

	go func() {
		err = s.Server.Serve(lis)
		if err != nil {
			s.lg.Errorw("GRPC server closed", err)
			s.eChan <- err
		}
	}()

	return nil
}

func (s *Server) Wait() <-chan error {
	return s.eChan
}

func (s *Server) Shutdown() bool {
	defer close(s.eChan)

	s.Server.GracefulStop()

	return true
}
