
.PHONY: install-dev
install-dev:
	aqua i -l

.PHONY: lint
lint: install-dev
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: test
test: lint
	mkdir -p tmp
	CGO_ENABLED=1 go test -race -p=4 -parallel=8 -timeout=300s -cover -coverprofile=./tmp/coverage.txt -shuffle=on ./...

.PHONY: build
build: test
	go build -o out/example-golang-cobra .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install-for-bash
install-for-bash: build
	./out/example-golang-cobra completion bash > out/completion.sh
	sudo cp out/example-golang-cobra /usr/local/bin/example-golang-cobra
	sudo mkdir -p /usr/local/share/bash-completion/completions
	sudo cp out/completion.sh /usr/local/share/bash-completion/completions/example-golang-cobra
	@echo "補完を有効にしたので、一度シェルを再起動してください (ex) 'exec bash'"

