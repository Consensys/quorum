//go:generate abigen -pkg bind -abi  ./AMC.abi							-bin  ./AMC.bin            				-type AMC   -out ../bind/amc.go
//go:generate abigen -pkg bind -abi  ./AmendRequest.abi              	-bin  ./AmendRequest.bin               	-type AmendRequest   -out ../bind/amendrequest.go
//go:generate abigen -pkg bind -abi  ./Mode.abi              			-bin  ./Mode.bin               			-type Mode   -out ../bind/mode.go
//go:generate abigen -pkg bind -abi  ./RouterService.abi              	-bin  ./RouterService.bin               -type RouterService   -out ../bind/routerservice.go
//go:generate abigen -pkg bind -abi  ./StandardLCFactory.abi            -bin  ./StandardLCFactory.bin           -type StandardLCFactory   -out ../bind/standardlcfactory.go
//go:generate abigen -pkg bind -abi  ./UPASLCFactory.abi           		-bin  ./UPASLCFactory.bin           	-type UPASLCFactory   -out ../bind/upaslcfactory.go

package gen
