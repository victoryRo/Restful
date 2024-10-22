DB_USER ?= victoJSON
DB_NAME ?= store
DB_TYPE ?= postgres
DB_PWD  ?= secretoJSONB

IP=127.0.0.1

URL_DB ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

.PHONY : rundb, execdb deletedb

rundb:
	@echo "Run the database container, volume, password ok"
	podman container run -d --name pgdb \
	-p 54321:5432 \
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

execdb:
	podman exec -it pgdb psql $(URL_DB)

deletedb:
	@echo "Delete container, image, volume ok"
	podman container stop pgdb || true && podman container rm pgdb || true && podman rmi postgres || true && podman volume prune -f 
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "--"
	@echo "------------------------------------------------ Deleted."

# Final line
