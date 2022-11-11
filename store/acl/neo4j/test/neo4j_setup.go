package test

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	Neo4jHttpPort = "7474"
	Neo4jBoltPort = "7687"
)

type neo4jContainer struct {
	testcontainers.Container
	uri    string
	dbName string
}

func setupNeo4jContainer(ctx context.Context) (*neo4jContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "neo4j:4.4.12",
		ExposedPorts: []string{Neo4jBoltPort},
		Env: map[string]string{
			"NEO4J_apoc_export_file_enabled":            "true",
			"NEO4J_apoc_import_file_enabled":            "true",
			"NEO4J_apoc_import_file_use__neo4j__config": "true",
			"NEO4JLABS_PLUGINS":                         "[\"apoc\"]",
			"NEO4J_dbms_connector_bolt_listen__address": ":" + Neo4jBoltPort,
			"NEO4J_dbms_connector_http_listen__address": ":" + Neo4jHttpPort,
			"NEO4J_dbms_security_auth__enabled":         "false",
		},
		WaitingFor: wait.ForLog("Applying default values for plugin apoc to neo4j.conf"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, Neo4jBoltPort)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("bolt://%s:%s", ip, mappedPort)

	return &neo4jContainer{Container: container, uri: uri, dbName: "neo4j"}, nil
}
