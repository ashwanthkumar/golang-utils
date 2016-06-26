test:
	go test -v -race github.com/ashwanthkumar/golang-utils/maps
	go test -v -race github.com/ashwanthkumar/golang-utils/sets
	go test -v -race github.com/ashwanthkumar/golang-utils/netutil
	go test -v -race github.com/ashwanthkumar/golang-utils/worker

test-only:
	go test -v -race github.com/ashwanthkumar/golang-utils/${name}

setup:
	# Test deps
	go get -u github.com/stretchr/testify/assert
