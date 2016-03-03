test:
	go test -v github.com/ashwanthkumar/golang-utils/maps
	go test -v github.com/ashwanthkumar/golang-utils/sets
	go test -v github.com/ashwanthkumar/golang-utils/netutil

setup:
	# Test deps
	go get -u github.com/stretchr/testify/assert
