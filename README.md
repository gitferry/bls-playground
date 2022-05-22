# bls-playground

This repo is to play with the Go binding of [blst](https://github.com/supranational/blst).

Run the sample with:
```
go run bls_playground.go
```

### Usage

There are two primary modes of operation that can be chosen based on type definitions in the application.

For minimal-pubkey-size operations:
```
type PublicKey = blst.P1Affine
type Signature = blst.P2Affine
type AggregateSignature = blst.P2Aggregate
type AggregatePublicKey = blst.P1Aggregate
```

For minimal-signature-size operations:
```
type PublicKey = blst.P2Affine
type Signature = blst.P1Affine
type AggregateSignature = blst.P1Aggregate
type AggregatePublicKey = blst.P2Aggregate
```

Add either the above type definitions in the begining of `utils/helpers.go`. The default is min-signature.

### Cross-compiling
If you're cross-compiling, you have to set `CC` environment variable to the target C cross-compiler and `CGO_ENABLED` to 1. For example, to run the code with Apple M1:

```
env GOARCH=arm64 CC=clang CGO_ENABLED=1 go run bls_playground.go
```