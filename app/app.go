package app

import (
	"fmt"
	checkergrpc "github.com/c12s/oort/api/checker/grpc"
	syncergrpc "github.com/c12s/oort/api/syncer/grpc"
	"github.com/c12s/oort/config"
	"github.com/c12s/oort/domain/handler"
	"github.com/c12s/oort/proto/checkerpb"
	"github.com/c12s/oort/proto/syncerpb"
	"github.com/c12s/oort/store/acl/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func Run(config config.Config) {
	manager, err := neo4j.NewTransactionManager(
		config.Neo4j().Uri(),
		config.Neo4j().Username(),
		config.Neo4j().Password(),
		config.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	aclStore := neo4j.NewAclStore(manager)

	checkerHandler := handler.NewCheckerHandler(aclStore)
	syncerHandler := handler.NewSyncerHandler(aclStore)

	checkerGrpcApi := checkergrpc.NewCheckerGrpcApi(checkerHandler)
	syncerGrpcApi := syncergrpc.NewSyncerGrpcApi(syncerHandler)

	startGrpcServer(config.Server().Host(), config.Server().Port(),
		checkerGrpcApi, syncerGrpcApi)
}

func startGrpcServer(host, port string, checkerGrpcApi checkergrpc.CheckerGrpcApi, syncerGrpcApi syncergrpc.SyncerGrpcApi) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	syncerpb.RegisterSyncerServiceServer(s, syncerGrpcApi)
	checkerpb.RegisterCheckerServiceServer(s, checkerGrpcApi)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
