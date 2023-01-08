default: build

test:
	go test ./...

build:
	go build -o bin/yatas-notion

update:
	go get -u 
	go mod tidy

install: build
	mkdir -p ~/.yatas.d/plugins/github.com/Thibaut-Padok/yatas-notion/local/
	mv ./bin/yatas-notion ~/.yatas.d/plugins/github.com/Thibaut-Padok/yatas-notion/local/

release: test
	standard-version
	git push --follow-tags origin main 