## go-nfs

[![GoDoc](https://godoc.org/github.com/adamdecaf/go-nfs?status.svg)](https://godoc.org/github.com/adamdecaf/go-nfs)
[![Build Status](https://travis-ci.com/adamdecaf/go-nfs.svg?branch=master)](https://travis-ci.com/adamdecaf/go-nfs)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdecaf/go-nfs)](https://goreportcard.com/report/github.com/adamdecaf/go-nfs)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/go-nfs/master/LICENSE)

A go client library for [NearlyFreeSpeech](https://nearlyfreespeech.net)'s api.

[Api Docs](https://members.nearlyfreespeech.net/wiki/API/Introduction)

## Getting / Usage

```bash
go get github.com/adamdecaf/go-nfs
```

## Building / Testing

You need the following environment variables set:

- `NFS_ACCOUNT_ID`
- `NFS_API_KEY`
- `NFS_LOGIN`


```bash
$ go get github.com/adamdecaf/go-nfs
$ cd go-nfs
$ go build .
$ go test -v .
```

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
