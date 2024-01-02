run-stage:
	docker compose up -d

init-src:
	go mod init example/pg-gin-connection

