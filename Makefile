create-migration: ## Create new migration file. It takes parameter `file` as filename. Usage: `make create-migration file=add_column_time`
	ls -x migrations/postgres/*.sql | tail -1 | awk -F"migrations/postgres/" '{print $$2}' | awk -F"_" '{print $$1}' | { read cur_v; expr $$cur_v + 1; } | { read new_v; printf "%06d" $$new_v; } | { read v; touch migrations/postgres/$$v"_$(file)".up.sql; touch migrations/postgres/$$v"_$(file)".down.sql; }
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

db-migrate:
	migrate -database ${DBUrl} -lock-timeout 30 -path migrations/postgres up