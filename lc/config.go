package lc

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type Config struct {
	AmendRequestAddress         common.Address `json:"amendRequestAddress"`
	LCManagementAddress         common.Address `json:"managementAddress"`
	StandardFactoryAddress      common.Address `json:"standardFactoryAddress"`
	UPASFactoryAddress          common.Address `json:"upasFactoryAddress"`
	RouterAddress               common.Address `json:"routerAddress"`
	AccessRoleManagementAddress common.Address `json:"accessRoleManagementAddress"`
	ModeAddress                 common.Address `json:"modeAddress"`
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
	accessRoleManagementAddress, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_ACCESS_ROLE_MANAGEMENT_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_ACCESS_ROLE_MANAGEMENT_ADDRESS", "err", err)
		return Config{}, err
	}
	modeAddr, err := common.NewMixedcaseAddressFromString(os.Getenv("LC_MODE_ADDRESS"))
	if err != nil {
		log.Error("error reading environment variable LC_MODE_ADDRESS", "err", err)
		return Config{}, err
	}
	return Config{
		AmendRequestAddress:         amendRequestAddr.Address(),
		LCManagementAddress:         lcManagementAddr.Address(),
		StandardFactoryAddress:      stdFacAddr.Address(),
		UPASFactoryAddress:          upasFacAddr.Address(),
		RouterAddress:               routerAddr.Address(),
		AccessRoleManagementAddress: accessRoleManagementAddress.Address(),
		ModeAddress:                 modeAddr.Address(),
	}, err
}
