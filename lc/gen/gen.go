//go:generate abigen -pkg bind -abi  ./AMC.abi					-type AMC   -out ../bind/amc.go
//go:generate abigen -pkg bind -abi  ./AmendRequest.abi			-type AmendRequest   -out ../bind/amendrequest.go
//go:generate abigen -pkg bind -abi  ./Mode.abi              	-type Mode   -out ../bind/mode.go
//go:generate abigen -pkg bind -abi  ./RouterService.abi        -type RouterService   -out ../bind/routerservice.go
//go:generate abigen -pkg bind -abi  ./StandardLCFactory.abi    -type StandardLCFactory   -out ../bind/standardlcfactory.go
//go:generate abigen -pkg bind -abi  ./UPASLCFactory.abi        -type UPASLCFactory   -out ../bind/upaslcfactory.go

package gen
