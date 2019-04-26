
clean:
	rm -rf build

build: clean
	mkdir build
	GOOS=linux go build -o build/main cmd/main.go
	scp build/main 10.5.119.133:/tmp
