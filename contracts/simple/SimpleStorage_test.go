package simple

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestSimple(t *testing.T) {

	key := `{"address":"ed9d02e382b34818e88b88a309c7fe71e65f419d","crypto":{"cipher":"aes-128-ctr","ciphertext":"4e77046ba3f699e744acb4a89c36a3ea1158a1bd90a076d36675f4c883864377","cipherparams":{"iv":"a8932af2a3c0225ee8e872bc0e462c11"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8ca49552b3e92f79c51f2cd3d38dfc723412c212e702bd337a3724e8937aff0f"},"mac":"6d1354fef5aa0418389b1a5d1f5ee0050d7273292a1171c51fd02f9ecff55264"},"id":"a65d1ac3-db7e-445d-a1cc-b6c5eeaa05e0","version":3}`
	auth, err := bind.NewTransactor(strings.NewReader(key), "")

	conn, err := ethclient.Dial("/Users/angelapratt/projects/go/src/github.com/quorum-examples/examples/7nodes/qdata/dd1/geth.ipc")

	privateFor := make([]string, 2)
	privateFor[0] = "1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg="
	privateFor[1] = "QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc="

	transactOpts := bind.TransactOpts{
		From:        auth.From,
		Signer:      auth.Signer,
		GasLimit:    3558096384,
		GasPrice:    big.NewInt(0),
		PrivateFrom: "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=",
		PrivateFor:  privateFor,
	}
	//set privacy fields

	fmt.Println("hi")
	addr, _, _, err := DeploySimpleStorage(&transactOpts, conn, big.NewInt(11))
	fmt.Println(addr.Hex())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}
