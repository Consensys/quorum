// generate stub/mock files for plugins, documentation and unit tests for plugin interfaces defined in .proto files
//
// need to install:
//  - protoc: 3.9.0+
//  - protoc-gen-go: 1.3.2+
//  - protoc-gen-doc: `go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc`
//  - mockgen: `go get -u github.com/golang/mock/mockgen`
//  - goimports: `go get -u golang.org/x/tools/cmd/goimports`
//
// go to terminal and run `go generate` from this directory

// generate stubs
//go:generate protoc -I ../../vendor/github.com/jpmorganchase/quorum-plugin-definitions -I ../../vendor --go_out=plugins=grpc:proto_common init.proto

// generate mocks for unit testing
//go:generate mockgen -package proto_common -destination proto_common/mock_init.go -source proto_common/init.pb.go

// fix fmt
//go:generate goimports -w ./

// generate documentation
//go:generate protoc -I ../../vendor/github.com/jpmorganchase/quorum-plugin-definitions -I ../../vendor --doc_out=docs.markdown.tmpl,init_interface.md:../../docs/PluggableArchitecture/Plugins/ init.proto
//go:generate protoc -I ../../vendor/github.com/jpmorganchase/quorum-plugin-definitions -I ../../vendor --doc_out=docs.markdown.tmpl,interface.md:../../docs/PluggableArchitecture/Plugins/helloworld/ helloworld.proto
//go:generate protoc -I ../../vendor/github.com/jpmorganchase/quorum-plugin-definitions -I ../../vendor --doc_out=docs.markdown.tmpl,interface.md:../../docs/PluggableArchitecture/Plugins/security/ security.proto
//go:generate protoc -I ../../vendor/github.com/jpmorganchase/quorum-plugin-definitions -I ../../vendor --doc_out=docs.markdown.tmpl,interface.md:../../docs/PluggableArchitecture/Plugins/account/ account.proto

package gen
