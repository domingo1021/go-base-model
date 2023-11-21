dockerup:
	docker compose --env-file ./config/docker.env up
dockerdown:
	docker compose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
run:
	go run main.go
.PHONY:
	dockerup