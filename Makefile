APP=shorty
TEST_FLAGS=-race -buildvcs

all: test build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP} .

test:
	go test ${TEST_FLAGS} ./...

format:
	go fmt .
	go mod tidy -v

clean:
	go clean
	\rm -fr ${APP}

up: 
	\podman build . -t ${APP}
	\podman run -p 8080:8080 ${APP}

audit:
	go mod verify
	go vet ./...
	go test ${TEST_FLAGS} -vet=off ./...
