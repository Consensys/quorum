module github.com/ethereum/go-ethereum

go 1.15

// Quorum - Replace Go modules that use modifications done by us
replace github.com/coreos/etcd => github.com/Consensys/etcd v3.3.13-quorum197+incompatible

// End Quorum

require (
	github.com/Azure/azure-storage-blob-go v0.7.0
	github.com/BurntSushi/toml v0.3.1
	github.com/VictoriaMetrics/fastcache v1.5.7
	github.com/aws/aws-sdk-go v1.25.48
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/cespare/cp v0.1.0
	github.com/cloudflare/cloudflare-go v0.10.2-0.20190916151808-a80f83b9add9
	github.com/consensys/gurvy v0.3.8
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/docker/docker v20.10.12+incompatible
	github.com/dop251/goja v0.0.0-20200721192441-a695b0cdd498
	github.com/eapache/channels v1.1.0
	github.com/eapache/queue v1.1.0 // indirect
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fatih/color v1.7.0
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff
	github.com/go-stack/stack v1.8.1
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.5.1
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3
	github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa
	github.com/google/uuid v1.1.5
	github.com/gorilla/websocket v1.5.0
	github.com/graph-gophers/graphql-go v1.3.0
	github.com/hashicorp/go-hclog v1.1.0
	github.com/hashicorp/go-plugin v1.2.2
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/holiman/bloomfilter/v2 v2.0.3
	github.com/holiman/uint256 v1.1.1
	github.com/huin/goupnp v1.0.1-0.20200620063722-49508fba0031
	github.com/influxdata/influxdb v1.8.3
	github.com/jackpal/go-nat-pmp v1.0.2-0.20160603034137-1fa385a6f458
	github.com/jedisct1/go-minisign v0.0.0-20190909160543-45766022959e
	github.com/jpmorganchase/quorum-account-plugin-sdk-go v0.0.0-20200714175524-662195b38a5e
	github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go v0.0.0-20200210211148-57f99f69eeb3
	github.com/jpmorganchase/quorum-security-plugin-sdk-go v0.0.0-20200714173835-22a319bb78ce
	github.com/julienschmidt/httprouter v1.2.0
	github.com/karalabe/usb v0.0.0-20190919080040-51dc0efba356
	github.com/mattn/go-colorable v0.1.4
	github.com/mattn/go-isatty v0.0.10
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.2-0.20190409134802-7e037d187b0c
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v0.0.0-20170112150404-1b00554d8222
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/prometheus/tsdb v0.7.1
	github.com/rjeczalik/notify v0.9.2
	github.com/rs/cors v1.7.0
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
	github.com/tv42/httpunix v0.0.0-20191220191345-2ba4b9c3382c
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
	golang.org/x/text v0.3.3
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc v1.29.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6
	gopkg.in/oleiade/lane.v1 v1.0.0
	gopkg.in/urfave/cli.v1 v1.20.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gotest.tools/v3 v3.1.0 // indirect
)
