package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	messenger "grpc-workshop/messenger"
)

type Server struct {
	listener net.Listener
	core     *grpc.Server
}

func New(listener net.Listener) *Server {
	core := grpc.NewServer()

	messenger.RegisterMessengerServiceServer(core, &Server{})

	return &Server{listener, core}
}

func (*Server) Send(_ context.Context, msg *messenger.Content) (*messenger.Content, error) {
	log.Println(fmt.Sprintf("[SERVER]: received %v at %v", msg.Body, msg.Date))
	return &messenger.Content{
		Body: "hello from server!",
		Date: uint64(time.Now().Unix()),
	}, nil
}

func (s *Server) Start() error {
	err := s.core.Serve(s.listener)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	s.core.Stop()
}
