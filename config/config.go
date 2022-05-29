package config

type Config interface {
	Neo4j() Neo4j
}

type config struct {
	neo4j Neo4j
}

func NewDefaultConfig() Config {
	return &config{
		neo4j: NewDefaultNeo4jConfig(),
	}
}

func (c config) Neo4j() Neo4j {
	return c.neo4j
}
