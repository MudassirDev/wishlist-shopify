build:
	go build -o out

run:
	./out

build-run:
	go build -o out && ./out

dev:
	air
