.PHONY: down
down: ## down
	docker-compose down --remove-orphans
build: ##build
	docker-compose build --no-cache
migrate: ## Migrate develop database
	mysqldef -u user -p user -h localhost -P 3306 dena-hack < ./_tools/mysql/schema.sql
