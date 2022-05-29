package config

type Neo4j interface {
	Uri() string
	Username() string
	Password() string
	DbName() string
}

type neo4j struct {
	uri      string
	username string
	password string
	dbName   string
}

//TODO: ovo je za prvu pomoc, izmeni posle
func NewDefaultNeo4jConfig() Neo4j {
	return neo4j{
		uri:      "bolt://localhost:7687",
		username: "neo4j",
		password: "t",
		dbName:   "neo4j",
	}
}

func (c neo4j) Uri() string {
	return c.uri
}

func (c neo4j) Username() string {
	return c.username
}

func (c neo4j) Password() string {
	return c.password
}

func (c neo4j) DbName() string {
	return c.dbName
}
