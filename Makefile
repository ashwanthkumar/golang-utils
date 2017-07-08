test:
	go test -v -race github.com/ashwanthkumar/golang-utils/maps
	go test -v -race github.com/ashwanthkumar/golang-utils/sets
	go test -v -race github.com/ashwanthkumar/golang-utils/netutil
	go test -v -race github.com/ashwanthkumar/golang-utils/worker
	go test -v -race github.com/ashwanthkumar/golang-utils/sync

test-only:
	go test -v -race github.com/ashwanthkumar/golang-utils/${name}

setup:
	glide install
