build: internal/*
	go build -o build/hentry-client internal/main.go

run:
	go run internal/main.go


compile:
	echo "Compiling for every OS and Platform"
	go build -o build/hentry-client internal/main.go
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm internal/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 internal/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 internal/main.go