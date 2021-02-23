package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
)

const (
	node1 = "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@127.0.0.1:21000?discport=0&raftport=50401"
	node2 = "enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402"
	node3 = "enode://579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778@127.0.0.1:21002?discport=0&raftport=50403"
)

func TestIsNodePermissioned(t *testing.T) {
	type args struct {
		nodename    string
		currentNode string
		datadir     string
		direction   string
	}
	d, _ := ioutil.TempDir("", "qdata")
	defer os.RemoveAll(d)
	writeNodeToFile(d, params.PERMISSIONED_CONFIG, node1)
	writeNodeToFile(d, params.PERMISSIONED_CONFIG, node3)
	writeNodeToFile(d, params.BLACKLIST_CONFIG, node3)
	n1, _ := enode.ParseV4(node1)
	n2, _ := enode.ParseV4(node2)
	n3, _ := enode.ParseV4(node3)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "node present",
			args: args{n1.ID().String(), n2.EnodeID(), d, "INWARD"},
			want: true,
		},
		{
			name: "node not present",
			args: args{n2.ID().String(), n1.EnodeID(), d, "OUTWARD"},
			want: false,
		},
		{
			name: "blacklisted node",
			args: args{n3.ID().String(), n1.EnodeID(), d, "INWARD"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNodePermissioned(tt.args.nodename, tt.args.currentNode, tt.args.datadir, tt.args.direction); got != tt.want {
				t.Errorf("IsNodePermissioned() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_isNodeBlackListed(t *testing.T) {
	type args struct {
		nodeName string
		dataDir  string
	}

	d, _ := ioutil.TempDir("", "qdata")
	defer os.RemoveAll(d)
	writeNodeToFile(d, params.BLACKLIST_CONFIG, node1)
	n1, _ := enode.ParseV4(node1)
	n2, _ := enode.ParseV4(node2)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "blacklisted node",
			args: args{n1.ID().String(), d},
			want: true,
		},
		{
			name: "blacklisted node",
			args: args{n2.ID().String(), d},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNodeBlackListed(tt.args.nodeName, tt.args.dataDir); got != tt.want {
				t.Errorf("isNodeBlackListed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func writeNodeToFile(dataDir, fileName, url string) {
	fileExists := true
	path := filepath.Join(dataDir, fileName)

	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		if _, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			return
		}
		fileExists = false
	}

	var nodeList []string
	var blob []byte
	if fileExists {
		blob, err := ioutil.ReadFile(path)
		if err == nil {
			if err := json.Unmarshal(blob, &nodeList); err != nil {
				return
			}
		}
	}
	nodeList = append(nodeList, url)
	blob, _ = json.Marshal(nodeList)
	_ = ioutil.WriteFile(path, blob, 0644)

}
