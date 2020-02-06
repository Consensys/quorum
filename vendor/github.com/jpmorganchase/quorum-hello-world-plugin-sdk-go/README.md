# Quorum Plugin SDK for Security Plugin Interface

**Note:** Github Actions workflow is only triggered via an event which is dispatched by `quorum-plugin-definitions` repository.

The SDK is auto generated from `quorum-plugin-definitions` proto files using `protoc` compiler 3.9.1 and `protoc-gen-go` 1.3.2.

- `proto` package contains stubs generated from `protoc` compiler.
- `mock_proto` package contains mocks for unit testing.

```
go get github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/proto
go get github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/mock_proto
```
