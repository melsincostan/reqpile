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

In the following commands, replace `$port` by the port the tool should listen on.
All of this assumes a linux host, as it is fully untested on any other OS (but should work the same)

Running without keeping an artefact:
```bash
go run . $port
```

Running a version that was compiled using `go build` (this should create an executable file called `reqpile`):
```bash
./reqpile $port
```

## Running over TLS

In the following commands, replace `$port` by the port the tool should listen on, `$cert` with the path to the TLS certificate and `$key` with the TLS certificate key.

Running direcctly, without a binary:
```bash
go run . $port tls $cert $key  
```

Running a binary (generated using `go build`)
```bash
./reqpile $port tls $cert $key
```

### Generating a testing self-signed certificate

A self-signed certificate can be generated interactively using openssl:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes
```

This will create a certificate valid for 365 days, without a passphrase, after asking a few questions. The certificate file itself will be `cert.pem`, and the private key will be `key.pem`.
This method requires apps wanting to interact with reqpile to be configured to accept self-signed certificates.
For non self-signed certificate, one could acquire a domain and use Let's Encrypt to get certificates for free.
