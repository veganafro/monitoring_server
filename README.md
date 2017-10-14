# Summary
This program monitors the health of an HTTP endpoint by sending a `GET` request and logging the status code in a file containing test samples.
A request is sent every 30 seconds. The status code returned in the server's response is logged. The time (in seconds since the UNIX epoch) is
also logged.

# Usage
Run the program from the command line in the following way:
`go run $GOPATH/src/monitoring_server/monitoring_server.go <ENDPOINT_URL> <PATH_TO_LOG_FILE>`
