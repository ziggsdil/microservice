CONFIG_PATH=./config/local.yaml
SERVER_PATH=./cmd/sso


run:
	go run $(SERVER_PATH) --config=$(CONFIG_PATH)

migrate:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations