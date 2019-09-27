
clean:
	rm word-count

build: clean
	go build -o word-count cmd/main.go