run-dev:
	go run cmd/api/main.go -env-mode=development -config-path=environments/example.yaml
run-test:
	cd internal/pkg/storage/sql; go test