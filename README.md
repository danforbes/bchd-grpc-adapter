# Chainlink `bchd` gRPC Adapter

A Chainlink adapter for [the `bchd` gRPC API](https://github.com/gcash/bchd/tree/master/bchrpc).

## Environment Variables

* `BCHD_GRPC_URL` - the URL to the `bchd` gRPC server (default: `bchd.greyh.at:8335`)
* `BCHD_CERT_PATH` - the path to the TLS client certificate for the `bchd` gRPC server (leave blank for servers with CA-signed certificates)
* `BCHD_SERVER_OVERRIDE` - for testing purposes only, can be used to override virtual host name in requests

## Usage

The `bchd` gRPC adapter exposes the following capabilities:

### Get Mempool Info

Get info about the mempool.

#### Params

* `proc` - `mempoolInfo`

### Get Mempool

Get info about all of the transactions currently in the memory pool.

#### Params

* `proc` - `mempool`
* `fullTrx` - if provided value converts to `false`, the response will only contain transaction hashes, otherwise, it will contain full transactions

### Get Blockchain Info

Get info about the blockchain including most recent block hash and height.

#### Params

* `proc` - `blockchainInfo`

## Errors

This adapter will throw errors in the following scenarios:

* It is not possible to create a TLS certificate using the path provided in the `BCHD_CERT_PATH` environment variable.
* It is not possible to parse the URL provided in the `BCHD_GRPC_URL` environment variable.
* It is not possible to connect to the `bchd` gRPC server.

## Test

```
go test
```

## Build

```
go build -o cl-bchd-grpc
```

## Lambda

Zip the binary

```
zip cl-bchd-grpc.zip cl-bchd-grpc
```

Then upload to AWS Lambda and use `cl-bchd-grpc` as the handler.
