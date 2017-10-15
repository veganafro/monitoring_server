# Summary
This program monitors the health of an HTTP endpoint by sending a `GET` request and logging the status code in a file containing test samples.
A request is sent every 30 seconds. The status code returned in the server's response is logged. The time (in seconds since the UNIX epoch) is
also logged. As of now, only http (not https) endpoints can be monitored. Redirects are not followed. Instead, the most recent response body
is used.

# Usage
Run the program from the command line with the following:

`go run $GOPATH/src/monitoring_server/monitoring_server.go <ENDPOINT_URL> <PATH_TO_LOG_FILE>`
