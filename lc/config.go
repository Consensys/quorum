package lc

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type Config struct {
	StandardFactoryAddress common.Address `json:"standardFactoryAddress"`
	UPASFactoryAddress     common.Address `json:"upasFactoryAddress"`
	RouterAddress          common.Address `json:"routerAddress"`
}

func ParseLcConfig() (Config, error) {
	stdFacAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_STANDARD_FACTORY_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_STANDARD_FACTORY_ADDRESS", "err", err)
		return Config{}, err
	}
	upasFacAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_UPAS_FACTORY_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_UPAS_FACTORY_ADDRESS", "err", err)
		return Config{}, err
	}
	routerAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_ROUTER_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_ROUTER_ADDRESS", "err", err)
		return Config{}, err
	}

	return Config{
		StandardFactoryAddress: stdFacAddr.Address(),
		UPASFactoryAddress:     upasFacAddr.Address(),
		RouterAddress:          routerAddr.Address(),
	}, err
}
