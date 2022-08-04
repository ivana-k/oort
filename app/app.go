package app

import (
	"fmt"
	checkergrpc "github.com/c12s/oort/api/checker/grpc"
	syncerasync "github.com/c12s/oort/api/syncer/async"
	syncergrpc "github.com/c12s/oort/api/syncer/grpc"
	"github.com/c12s/oort/async/nats"
	"github.com/c12s/oort/config"
	"github.com/c12s/oort/domain/checker"
	"github.com/c12s/oort/domain/syncer"
	"github.com/c12s/oort/proto/checkerpb"
	"github.com/c12s/oort/proto/syncerpb"
	"github.com/c12s/oort/store/acl/neo4j"
	cache2 "github.com/c12s/oort/store/cache"
	"github.com/c12s/oort/store/cache/gocache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func Run(config config.Config) {
	manager, err := neo4j.NewTransactionManager(
		config.Neo4j().Uri(),
		config.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	aclStore := neo4j.NewAclStore(manager)
	cache, err := gocache.NewGoCache(config.Redis().Address(), config.Redis().Eviction())
	if err != nil {
		panic(err)
	}

	checkerHandler := checker.NewHandler(aclStore, cache, cache2.NewProtoAttributeSerializer(), cache2.NewCustomCheckPermissionSerializer())
	syncerHandler := syncer.NewHandler(aclStore, syncerpb.NewSyncRespOutboxMessage)

	checkerGrpcApi := checkergrpc.NewCheckerGrpcApi(checkerHandler)
	syncerGrpcApi := syncergrpc.NewSyncerGrpcApi(syncerHandler)

	natsConn, err := nats.NewConnection(config.Nats().Uri())
	if err != nil {
		panic(err)
	}
	subscriber, err := nats.NewSubscriber(natsConn)
	if err != nil {
		panic(err)
	}
	err = syncerasync.NewSyncerAsyncApi(subscriber, "sync", "oort", syncerpb.NewSyncMessageProtoSerializer(), syncerHandler)
	if err != nil {
		panic(err)
	}

	startGrpcServer(config.Server().Port(),
		checkerGrpcApi, syncerGrpcApi)
}

func startGrpcServer(port string, checkerGrpcApi checkergrpc.CheckerGrpcApi, syncerGrpcApi syncergrpc.SyncerGrpcApi) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
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
