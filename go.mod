module github.com/ethereum/go-ethereum

go 1.19

// Quorum - Replace Go modules that use modifications done by us
replace github.com/coreos/etcd => github.com/Consensys/etcd v3.3.13-quorum197+incompatible

// End Quorum

require (
	github.com/Azure/azure-storage-blob-go v0.7.0
	github.com/BurntSushi/toml v0.3.1
	github.com/ConsenSys/quorum-qlight-token-manager-plugin-sdk-go v0.0.0-20220427130631-ecd75caa6e73
	github.com/VictoriaMetrics/fastcache v1.5.7
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/config v1.1.1
	github.com/aws/aws-sdk-go-v2/credentials v1.1.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.1.1
	github.com/btcsuite/btcd/btcec/v2 v2.1.2
	github.com/cespare/cp v0.1.0
	github.com/cloudflare/cloudflare-go v0.14.0
	github.com/consensys/gnark-crypto v0.12.1
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/docker/docker v25.0.3+incompatible
	github.com/dop251/goja v0.0.0-20200721192441-a695b0cdd498
	github.com/eapache/channels v1.1.0
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fatih/color v1.7.0
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff
	github.com/go-stack/stack v1.8.1
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3
	github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/graph-gophers/graphql-go v1.3.0
	github.com/hashicorp/go-hclog v1.1.0
	github.com/hashicorp/go-plugin v1.2.2
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/holiman/bloomfilter/v2 v2.0.3
	github.com/holiman/uint256 v1.2.1
	github.com/huin/goupnp v1.0.1-0.20210310174557-0ca763054c88
	github.com/influxdata/influxdb v1.8.3
	github.com/jackpal/go-nat-pmp v1.0.2-0.20160603034137-1fa385a6f458
	github.com/jedisct1/go-minisign v0.0.0-20190909160543-45766022959e
	github.com/jpmorganchase/quorum-account-plugin-sdk-go v0.0.0-20200714175524-662195b38a5e
	github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go v0.0.0-20200210211148-57f99f69eeb3
	github.com/jpmorganchase/quorum-security-plugin-sdk-go v0.0.0-20200714173835-22a319bb78ce
	github.com/julienschmidt/httprouter v1.3.0
	github.com/karalabe/usb v0.0.2
	github.com/mattn/go-colorable v0.1.4
	github.com/mattn/go-isatty v0.0.14
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v0.0.0-20170112150404-1b00554d8222
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/prometheus/tsdb v0.7.1
	github.com/rjeczalik/notify v0.9.2
	github.com/rs/cors v1.7.0
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	github.com/stretchr/testify v1.8.2
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	github.com/tv42/httpunix v0.0.0-20191220191345-2ba4b9c3382c
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
	golang.org/x/crypto v0.17.0
	golang.org/x/sync v0.3.0
	golang.org/x/sys v0.15.0
	golang.org/x/term v0.15.0
	golang.org/x/text v0.14.0
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.31.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6
	gopkg.in/oleiade/lane.v1 v1.0.0
	gopkg.in/urfave/cli.v1 v1.20.0
)

require (
	github.com/Azure/azure-pipeline-go v0.2.1 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.21 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.1.1 // indirect
	github.com/aws/smithy-go v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.7.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/hashicorp/yamux v0.0.0-20180604194846-3520598351bb // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-ieproxy v0.0.0-20190610004146-91bb50d98149 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.1.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
