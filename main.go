package main

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
	"github.com/c12s/oort/store/cache"
	"github.com/c12s/oort/store/cache/gocache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.NewConfig()

	manager, err := neo4j.NewTransactionManager(
		cfg.Neo4j().Uri(),
		cfg.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Stop()

	aclStore := neo4j.NewAclStore(manager, neo4j.NewSimpleCypherFactory())
	log.Println("STAAAAART")
	c, err := gocache.NewGoCache(
		cfg.Redis().Address(),
		cfg.Redis().Eviction(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Stop()

	checkerHandler := checker.NewHandler(
		aclStore, c,
		cache.NewProtoAttributeSerializer(),
		cache.NewCustomCheckPermissionSerializer(),
	)
	syncerHandler := syncer.NewHandler(
		aclStore,
		syncerpb.NewSyncRespOutboxMessage,
	)

	checkerGrpcApi := checkergrpc.NewCheckerGrpcApi(checkerHandler)
	syncerGrpcApi := syncergrpc.NewSyncerGrpcApi(syncerHandler)

	natsConn, err := nats.NewConnection(cfg.Nats().Uri())
	if err != nil {
		log.Fatal(err)
	}
	defer natsConn.Close()
	subscriber, err := nats.NewSubscriber(natsConn)
	if err != nil {
		log.Fatal(err)
	}

	err = syncerasync.NewSyncerAsyncApi(
		subscriber,
		"sync",
		"oort",
		syncerpb.NewProtoSyncMessageSerializer(),
		syncerHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	syncerpb.RegisterSyncerServiceServer(s, syncerGrpcApi)
	checkerpb.RegisterCheckerServiceServer(s, checkerGrpcApi)
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGINT)
	<-quit

	s.GracefulStop()
}
