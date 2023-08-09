package startup

import (
	"github.com/c12s/oort/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func startServer(address string, administrator api.OortAdministratorServer, evaluator api.OortEvaluatorServer) (chan bool, error) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	api.RegisterOortAdministratorServer(s, administrator)
	api.RegisterOortEvaluatorServer(s, evaluator)
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	var serverStoppedCh chan bool
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGINT)
		<-quit

		s.GracefulStop()
		serverStoppedCh <- true
	}()

	return serverStoppedCh, nil
}
