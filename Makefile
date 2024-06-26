run:
	go run ./cmd/ customers test.csv
switch-sqlite:
	cp .env.sqlite.example .env
switch-mysql:
	cp .env.mysql.example .env
switch-pgsql:
	cp .env.pgsql.example .env
switch-firebird:
	cp .env.firebird.example .env
lint:
	gocritic check ./...
	revive ./...
	golint ./...
	go vet ./...
	staticcheck ./...
	golangci-lint run
	goconst ./...
test:
	go test ./...
