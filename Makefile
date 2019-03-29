build:
	mkdir -p ./functions
	go build -o ./function ./main.go
	cp ./function ./functions/ping
