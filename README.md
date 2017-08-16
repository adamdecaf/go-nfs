## go-nfs

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
