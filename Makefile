build:
	go build -o server/otus ./server

run: build
	server/otus

