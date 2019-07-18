all : dolphin cli

dolphin : server/main.go
	go build -o dolphin $^

cli : test_cli/main.go
	go build -o cli $^

clean:
	rm dolphin cli
