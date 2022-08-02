export OORT_HOSTNAME=oort
export OORT_PORT=8000

export NEO4J_HOSTNAME=neo4j
export NEO4J_BOLT_PORT=7687
export NEO4J_HTTP_PORT=7474
export NEO4J_AUTH_ENABLED=false
export NEO4J_DBNAME=neo4j

export NEO4J_apoc_export_file_enabled=true
export NEO4J_apoc_import_file_enabled=true
export NEO4J_apoc_import_file_use__neo4j__config=true
export NEO4JLABS_PLUGINS="[\"apoc\"]"

export NATS_HOSTNAME=nats
export NATS_PORT=4222
export NATS_ENABLE_AUTH=yes
export NATS_USERNAME=user
export NATS_PASSWORD=pass

export POLL_INTERVAL_IN_MS=5000

docker-compose build
docker-compose up