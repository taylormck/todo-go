run:
    go run .

build:
    go build -o bin/todo

run-release: build
    ./bin/todo
