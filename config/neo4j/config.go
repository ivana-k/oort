package neo4j

type Config interface {
	Uri() string
	Username() string
	Password() string
	DbName() string
}

type config struct {
	uri      string
	username string
	password string
	dbName   string
}

func NewDefaultNeo4jConfig() Config {
	return config{
		uri:      "bolt://localhost:7687",
		username: "neo4j",
		password: "t",
		dbName:   "neo4j",
	}
}

func (c config) Uri() string {
	return c.uri
}

func (c config) Username() string {
	return c.username
}

func (c config) Password() string {
	return c.password
}

func (c config) DbName() string {
	return c.dbName
}
