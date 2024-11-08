.PHONY: lint
lint:
	# | jq > ./golangci-lint/report.json
	golangci-lint run --fix -c .golangci.yml > golangci-lint/report-unformatted.json
	goimports -local mmskazak -w .

.PHONY: lind
lind:
	docker run --rm \
		-v $(pwd):/app \
		-v $(pwd)/golangci-lint/.cache/golangci-lint/v1.57.2:/root/.cache \
		-w /app \
		golangci/golangci-lint:v1.57.2 \
			golangci-lint run \
				-c .golangci.yml \
			> ./golangci-lint/report-unformatted.json

lint-clean:
	sudo rm -rf ./golangci-lint

test:
	go test ./...

# Параметры контейнера и образа
CONTAINER_NAME=goph_keeper
IMAGE=postgres:16.3
POSTGRES_USER=gkuser
POSTGRES_PASSWORD=gkpass
POSTGRES_DB=goph_keeper
VOLUME_NAME=goph_keeper

# Команда для запуска контейнера PostgreSQL
db:
	docker run -d \
        --name $(CONTAINER_NAME) \
        -e POSTGRES_USER=$(POSTGRES_USER) \
        -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
        -e POSTGRES_DB=$(POSTGRES_DB) \
        -p 5432:5432 \
        -v $(VOLUME_NAME):/var/lib/postgresql/data \
        $(IMAGE)
