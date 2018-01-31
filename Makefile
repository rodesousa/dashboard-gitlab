
all: build run

build:
	@go build .

run:
	@./dashboard-gitlab

prod:
	@cd ./frontend && \
		yarn build

front-end:
	@cd ./frontend && \
		yarn start

test-merge-get:
	@curl http://localhost:8080/merge_request/all

test-repo-add:
	@curl http://localhost:8080/repo/add

test-repo-save:
	@curl http://localhost:8080/repo/save

test-merge-maj:
	@curl http://localhost:8080/merge_request/maj

