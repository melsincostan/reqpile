# Reqpile

Tiny, friendly HTTP debugging tool

The goal of this tool is to check that requests are being sent correctly when an actual API server is overkill or not present. This simply prints the request contents (headers + body bytes) to the console, surrounded by a bit of formatting.
The request display itself is made to look like the actual text sent over HTTP. It doesn't print much more than what you would see by listening with netcat: additions are separators between requests, a request counter that counts requests received since the start of the program and a source IP and port.
It replies to all requests with http status 204 and no body.

## Building

This doesnt use anything outside of the go standard library. Just run:

```bash
go build
```

To remove the artifact, run:

```bash
go clean
```

## Running

In the following commands, replace `$port` by the port you want the tool to listen on.
All of this assumes a linux host, as it is fully untested on any other OS (but should work the same)

Running without keeping an artefact:
```bash
go run . $port
```

Running a version that was compiled using `go build` (this should create an executable file called `reqpile`):
```bash
./reqpile $port
```
