e2e_test: docker_build_test
	docker compose up -d
	docker compose exec -T pokedex-service go test ./...
	docker compose down

docker_build_test:
	docker build . -t pokedex-app-test
