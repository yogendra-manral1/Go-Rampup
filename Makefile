hello:
	echo "Hello"

build:
	echo "Installing goimports"
	go install golang.org/x/tools/cmd/goimports@latest
	echo "Installation complete"
	go build -o bin/main main.go

run:
	$(info $(M) Running DB migrations…)
	./goose up
	go run main.go

ifeq (migrate, $(firstword $(MAKECMDGOALS)))
  migrateargs := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
  $(eval $(migrateargs):;@true)
endif

build-goose:
	$(info $(M) Building goose…)
	go build -o goose goose.go

migrate:
	$(info $(M) Executing migrate command…)
	./goose $(migrateargs)

all: hello build run
