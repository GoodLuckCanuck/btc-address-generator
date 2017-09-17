btc-address-generator
=====================

###### Simple tool to generate bitcoin addresses (with corresponding private keys) to STDOUT.

##### [Install Go](http://golang.org/doc/install)

##### Install dependencies:
```bash
$ export GOPATH=~/btc-address-generator
$ go get github.com/btcsuite/btcutil
```

##### Run:
```bash
$ export GOPATH=~/btc-address-generator
$ go run btc-address-generator.go
```

Note: if using an older version of golang, you may run into issues with the btcsuite/btcutil code, in that LastIndexByte is not available. It might be safe for the purposes of this program to comment out affected lines and replace it with hardcoded values of 1, for example.
