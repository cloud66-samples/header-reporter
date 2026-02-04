# Header Reporter

A simple Go web application that displays HTTP request information including all headers.

## Features

- Shows request method, URL, path, query string, protocol, host, and remote address
- Displays all HTTP headers sorted alphabetically
- Lists query parameters if present
- Plain text output for easy reading

## Running Locally

```bash
go run main.go
```

By default, the server listens on port 80. Set the `PORT` environment variable to change it:

```bash
PORT=8080 go run main.go
```

## Running with Docker

```bash
docker build -t header-reporter .
docker run -p 8080:80 header-reporter
```

Then visit `http://localhost:8080` to see your request details.

## Example Output

```
=== REQUEST INFO ===

Method:         GET
URL:            /test?foo=bar
Path:           /test
Raw Query:      foo=bar
Protocol:       HTTP/1.1
Host:           localhost:8080
Remote Address: 172.17.0.1:54321
Request URI:    /test?foo=bar

=== HEADERS ===

Accept: */*
Host: localhost:8080
User-Agent: curl/8.7.1

=== QUERY PARAMETERS ===

foo: bar
```
