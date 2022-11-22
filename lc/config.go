package lc

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type Config struct {
	AmendRequestAddress    common.Address `json:"AmendRequestAddress"`
	LCManagementAddress    common.Address `json:"ManagementAddress"`
	StandardFactoryAddress common.Address `json:"StandardFactoryAddress"`
	UPASFactoryAddress     common.Address `json:"UpasFactoryAddress"`
	RouterAddress          common.Address `json:"RouterAddress"`
	// AccessRoleManagementAddress common.Address `json:"accessRoleManagementAddress"`
	ModeAddress common.Address `json:"ModeAddress"`
}

func ParseLcConfig() (Config, error) {
	amendRequestAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_ADMEND_REQUEST_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_ADMEND_REQUEST_ADDRESS", "err", err)
		return Config{}, err
	}
	lcManagementAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_MANAGEMENT_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_MANAGEMENT_ADDRESS", "err", err)
		return Config{}, err
	}
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
	modeAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_MODE_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_MODE_ADDRESS", "err", err)
		return Config{}, err
	}
	return Config{
		AmendRequestAddress:    amendRequestAddr.Address(),
		LCManagementAddress:    lcManagementAddr.Address(),
		StandardFactoryAddress: stdFacAddr.Address(),
		UPASFactoryAddress:     upasFacAddr.Address(),
		RouterAddress:          routerAddr.Address(),
		ModeAddress:            modeAddr.Address(),
	}, err
}
