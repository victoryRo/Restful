DB_USER ?= victory
DB_NAME ?= urldb
DB_TYPE ?= postgres
DB_PWD  ?= secretoSafe

IP=127.0.0.1

URL_DB ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

.PHONY : runpostgres createdb pgexec pgdown teardown_recreate

runpostgres:
	@echo "Run the database container, volume, password ok"
	podman container run -d --name postgresdb \
	-p 54320:5432 \
	-v $(PWD):/postgres/data:z \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_DB=$(DB_NAME) \
	-e POSTGRES_PASSWORD=$(DB_PWD) \
	postgres:latest
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "------------------------------------------------ Container RUN..."

pgexec:
	podman exec -it postgresdb psql $(URL_DB)

createdb:
	podman exec -it postgresdb psql $(URL_DB) -c "\i /postgres/data/db/schema.sql"
	
pgdown:
	@echo "Delete container, image, volume ok"
	podman container stop postgresdb || true && podman container rm postgresdb || true && podman rmi postgres || true && podman volume prune -f 
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "------------------------------------------------ Deleted."

teardown_recreate: pgdown runpostgres  


# Final line
