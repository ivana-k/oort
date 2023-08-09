package startup

import (
	"fmt"
	"github.com/c12s/magnetar/pkg/messaging/nats"
	"github.com/c12s/oort/internal/configs"
	"github.com/c12s/oort/internal/handlers"
	"github.com/c12s/oort/internal/repos/rhabac/neo4j"
	"github.com/c12s/oort/internal/services"
	"github.com/c12s/oort/pkg/api"
)

func StartApp(config configs.Config) error {
	manager, err := neo4j.NewTransactionManager(
		config.Neo4j().Uri(),
		config.Neo4j().DbName())
	if err != nil {
		return err
	}
	defer manager.Stop()

	neo4jRhabacStore := neo4j.NewRHABACStore(manager, neo4j.NewSimpleCypherFactory())

	evaluationService, err := services.NewEvaluationService(neo4jRhabacStore)
	if err != nil {
		return err
	}
	administrationService, err := services.NewAdministrationService(neo4jRhabacStore)
	if err != nil {
		return err
	}

	evaluatorGrpcServer, err := handlers.NewOortEvaluatorGrpcServer(*evaluationService)
	if err != nil {
		return err
	}
	administratorGrpcServer, err := handlers.NewOortAdministratorGrpcServer(*administrationService)
	if err != nil {
		return err
	}

	natsConn, err := newNatsConn(config.Nats().Uri())
	if err != nil {
		return err
	}
	defer natsConn.Close()
	adminReqSubscriber, err := nats.NewSubscriber(natsConn, api.AdministrationReqSubject, "oort")
	if err != nil {
		return err
	}
	adminRespPublisher, err := nats.NewPublisher(natsConn)
	if err != nil {
		return err
	}

	err = handlers.NewAsyncAdministratorHandler(
		adminReqSubscriber,
		adminRespPublisher,
		*administrationService,
	)
	if err != nil {
		return err
	}

	serverStoppedCh, err := startServer(fmt.Sprintf(":%s", config.Server().Port()), administratorGrpcServer, evaluatorGrpcServer)

	<-serverStoppedCh

	return nil
}
