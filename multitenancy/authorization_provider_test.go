package multitenancy

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(false)))
}

type testCase struct {
	msg          string
	granted      []string
	ask          []*ContractSecurityAttribute
	isAuthorized bool
}

func TestMatch_whenTypical(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/create/contracts?from.tm=A/")
	ask, _ := url.Parse("private://0xa1b1c1/create/contracts?from.tm=A%2F")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenAskNothing(t *testing.T) {
	granted, _ := url.Parse("private://0x0/_/contracts?from.tm=A&owned.eoa=0x0")
	ask, _ := url.Parse("private://0xa1b1c1/write/contracts?owned.eoa=0xe1e1e1")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))

	ask, _ = url.Parse("private://0xa1b1c1/write/contracts")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenGrantNothing(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/write/contracts")
	ask, _ := url.Parse("private://0xa1b1c1/write/contracts?from.tm=A")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenAnyAction(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/_/contracts?owned.eoa=0x0&from.tm=A1")
	ask, _ := url.Parse("private://0xa1b1c1/read/contracts?from.tm=A1")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))

	ask, _ = url.Parse("private://0xa1b1c1/read/contracts?owned.eoa=0x0&from.tm=A1&from.tm=B1")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))

	ask, _ = url.Parse("private://0xa1b1c1/write/contracts?owned.eoa=0x0&from.tm=A1")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionWrite,
	}, ask, granted))
}

func TestMatch_whenPathNotMatched(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/write/contracts?owned.eoa=0x0&from.tm=A1")
	ask, _ := url.Parse("private://0xa1b1c1/read/contracts?from.tm=A1")

	assert.False(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenSchemeIsNotEqual(t *testing.T) {
	granted, _ := url.Parse("unknown://0xa1b1c1/create/contracts?from.tm=A")
	ask, _ := url.Parse("private://0xa1b1c1/create/contracts?from.tm=A")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenContractWritePermission_GrantedIsTheSuperSet(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A&from.tm=B")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionWrite,
	}, ask, granted), "with write permission")

	granted, _ = url.Parse("private://0x0/read/contracts?owned.eoa=0x0&from.tm=A&from.tm=B")
	ask, _ = url.Parse("private://0x0/read/contracts?owned.eoa=0x0&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted), "with read permission")
}

func TestMatch_whenContractReadPermission_AnyAction(t *testing.T) {
	granted, _ := url.Parse("private://0x1234/_/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x1234&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractReadPermission_AnyEoa(t *testing.T) {
	granted, _ := url.Parse("private://0x1234/_/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x0&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractReadPermission_EoaDifferent(t *testing.T) {
	granted, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x095e7baea6a6c7c4c2dfeb977efac326af552d87&from.tm=A")
	ask, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x945304eb96065b2a98b57a48a06ae28d285a71b5&from.tm=A")

	assert.False(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractReadPermission_EoaSame(t *testing.T) {
	granted, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x095e7baea6a6c7c4c2dfeb977efac326af552d87&from.tm=A")
	ask, _ := url.Parse("private://0x0/read/contracts?owned.eoa=0x095e7baea6a6c7c4c2dfeb977efac326af552d87&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractReadPermission_TmKeysIntersect(t *testing.T) {
	granted, _ := url.Parse("private://0x0/read/contracts?from.tm=A&from.tm=B")
	ask, _ := url.Parse("private://0x0/read/contracts?from.tm=B&from.tm=C")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractReadPermission_TmKeysDontIntersect(t *testing.T) {
	granted, _ := url.Parse("private://0x0/read/contracts?from.tm=A&from.tm=B")
	ask, _ := url.Parse("private://0x0/read/contracts?from.tm=C&from.tm=D")

	assert.False(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionRead,
	}, ask, granted))
}

func TestMatch_whenContractWritePermission_Same(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionWrite,
	}, ask, granted))
}

func TestMatch_whenContractWritePermission_Different(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=B")

	assert.False(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionWrite,
	}, ask, granted))
}

func TestMatch_whenContractWritePermission_AskIsSuperSet(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&from.tm=B&from.tm=C&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionWrite,
	}, ask, granted))
}

func TestMatch_whenContractCreatePermission_Same(t *testing.T) {
	granted, _ := url.Parse("private://0x0/create/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/create/contracts?owned.eoa=0x0&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionCreate,
	}, ask, granted))
}

func TestMatch_whenContractCreatePermission_Different(t *testing.T) {
	granted, _ := url.Parse("private://0x0/create/contracts?owned.eoa=0x0&from.tm=A")
	ask, _ := url.Parse("private://0x0/create/contracts?owned.eoa=0x0&from.tm=B")

	assert.False(t, match(&ContractSecurityAttribute{
		Visibility: VisibilityPrivate,
		Action:     ActionCreate,
	}, ask, granted))
}

func TestMatch_whenUsingWildcardAccount(t *testing.T) {
	granted, _ := url.Parse("private://0x0/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")
	ask, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))

	granted, _ = url.Parse("private://0x0/read/contract?owned.eoa=0x0&from.tm=A")
	ask, _ = url.Parse("private://0xa1b1c1/read/contract?owned.eoa=0x1234&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionRead}, ask, granted))
}

func TestMatch_whenNotUsingWildcardAccount(t *testing.T) {
	granted, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")
	ask, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))

	granted, _ = url.Parse("private://0x0/read/contract?owned.eoa=0x0&from.tm=A")
	ask, _ = url.Parse("private://0xa1b1c1/read/contract?owned.eoa=0x1234&from.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionRead}, ask, granted))
}

func TestMatch_failsWhenAccountsDiffer(t *testing.T) {
	granted, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")
	ask, _ := url.Parse("private://0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenPublic(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/create/contract?from.tm=A/")
	ask, _ := url.Parse("public://0x0/create/contract")

	assert.True(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func TestMatch_whenNotEscaped(t *testing.T) {
	// query not escaped probably in the granted authority resource identitifer
	granted, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=")
	ask, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=BULeR8JyUWhiuuCMU%2FHLA0Q5pzkYT%2BcHII3ZKBey3Bo%3D")

	assert.False(t, match(&ContractSecurityAttribute{Action: ActionCreate}, ask, granted))
}

func runTestCases(t *testing.T, testCases []*testCase) {
	testObject := &DefaultContractAuthorizationProvider{}
	for _, tc := range testCases {
		log.Debug("--> Running test case: " + tc.msg)
		authorities := make([]*proto.GrantedAuthority, 0)
		for _, a := range tc.granted {
			authorities = append(authorities, &proto.GrantedAuthority{Raw: a})
		}
		b, err := testObject.IsAuthorized(
			context.Background(),
			&proto.PreAuthenticatedAuthenticationToken{Authorities: authorities},
			tc.ask...)
		if !assert.NoError(t, err, tc.msg) {
			return
		}
		if !assert.Equal(t, tc.isAuthorized, b, tc.msg) {
			return
		}
	}
}

func TestDefaultAccountAccessDecisionManager_IsAuthorized_forPublicContracts(t *testing.T) {
	runTestCases(t, []*testCase{
		canCreatePublicContracts,
		// canNotCreatePublicContracts,
		canReadOwnedPublicContracts,
		canReadOtherPublicContracts,
		// canNotReadOtherPublicContracts,
		canWriteOwnedPublicContracts,
		canWriteOtherPublicContracts1,
		canWriteOtherPublicContracts2,
		// canNotWriteOtherPublicContracts,
		canCreatePublicContractsAndWriteToOthers,
	})
}

func TestDefaultAccountAccessDecisionManager_IsAuthorized_forPrivateContracts(t *testing.T) {
	runTestCases(t, []*testCase{
		canCreatePrivateContracts,
		canNotCreatePrivateContracts,
		canReadOwnedPrivateContracts,
		canReadOtherPrivateContracts,
		canNotReadOtherPrivateContracts,
		canNotReadOtherPrivateContractsNoPrivy,
		canWriteOwnedPrivateContracts,
		canWriteOtherPrivateContracts,
		canWriteOtherPrivateContractsWithOverlappedScope,
		canNotWriteOtherPrivateContracts,
		canNotWriteOtherPrivateContractsNoPrivy,
	})
}

func TestDefaultAccountAccessDecisionManager_IsAuthorized_forPrivateContracts_wildcards_whenCreate(t *testing.T) {
	fullAccessToX := []string{
		"private://0x0/_/contracts?owned.eoa=0x0&from.tm=X",
	}
	runTestCases(t, []*testCase{
		{
			msg:          "X has full access to a private contract when create",
			isAuthorized: true,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				// create
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionCreate,
					PrivateFrom: "X",
					Parties:     []string{},
				},
			},
		},
		{
			msg:          "X can't creat private contract with other TM key",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				// create
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionCreate,
					PrivateFrom: "A",
					Parties:     []string{},
				},
			},
		},
	})
}

func TestDefaultAccountAccessDecisionManager_IsAuthorized_forPrivateContracts_wildcards_whenRead(t *testing.T) {
	fullAccessToX := []string{
		"private://0x0/_/contracts?owned.eoa=0x0&from.tm=X",
	}
	runTestCases(t, []*testCase{
		{
			msg:          "X has full access to a private contract when read as one of the participants",
			isAuthorized: true,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "X",
					Parties:                       []string{"X", "Y"},
				},
			},
		},
		{
			msg:          "X has full access to a private contract when read as a single participant",
			isAuthorized: true,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "X",
					Parties:                       []string{"X"},
				},
			},
		},
		{
			msg:          "X can't read other private contracts",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "X",
					Parties:                       []string{"A", "B"},
				},
			},
		},
		{
			msg:          "X can't read other private contracts by faking the read",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "A",
					Parties:                       []string{"A", "B"},
				},
			},
		},
		{
			msg:          "X can't read other private contracts when proxy-read",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				// read its own contract
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "X",
					Parties:                       []string{"X"},
				},
				// but using it as proxy to read other contract
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
					Visibility:                    VisibilityPrivate,
					Action:                        ActionRead,
					PrivateFrom:                   "X",
					Parties:                       []string{"A", "B"},
				},
			},
		},
	})
}

func TestDefaultAccountAccessDecisionManager_IsAuthorized_forPrivateContracts_wildcards_whenWrite(t *testing.T) {
	fullAccessToX := []string{
		"private://0x0/_/contracts?owned.eoa=0x0&from.tm=X",
	}
	runTestCases(t, []*testCase{
		{
			msg:          "X has full access to a private contract when write as a single participant",
			isAuthorized: true,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionWrite,
					PrivateFrom: "X",
					Parties:     []string{"X"},
				},
			},
		},
		{
			msg:          "X has full access to a private contract when write as one of the participants",
			isAuthorized: true,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionWrite,
					PrivateFrom: "X",
					Parties:     []string{"X", "Y"},
				},
			},
		},
		{
			msg:          "X must not access other private contracts when faking write",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
						To:   common.HexToAddress("0xb1b1b1"), // creator EOA address
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionWrite,
					PrivateFrom: "A",
					Parties:     []string{"A", "B"},
				},
			},
		},
		{
			msg:          "X can not write to a private contract not privy to X",
			isAuthorized: false,
			granted:      fullAccessToX,
			ask: []*ContractSecurityAttribute{
				{
					AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
						From: common.HexToAddress("0xa1a1a1"),
						To:   common.HexToAddress("0xb1b1b1"), // creator EOA address
					},
					Visibility:  VisibilityPrivate,
					Action:      ActionWrite,
					PrivateFrom: "X",
					Parties:     []string{"A", "B"},
				},
			},
		},
	})
}

var (
	canCreatePublicContracts = &testCase{
		msg: "0x0a1a1a1 can create public contracts",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
		},
		ask: []*ContractSecurityAttribute{
			{
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
				},
				Visibility: VisibilityPublic,
				Action:     ActionCreate,
			},
		},
		isAuthorized: true,
	}
	canCreatePublicContractsAndWriteToOthers = &testCase{
		msg: "0x0a1a1a1 can create public contracts and write to contracts created by 0xb1b1b1",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		ask: []*ContractSecurityAttribute{
			{
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
				},
				Visibility: VisibilityPublic,
				Action:     ActionCreate,
			}, {
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
					To:   common.HexToAddress("0xb1b1b1"),
				},
				Visibility: VisibilityPublic,
				Action:     ActionWrite,
			},
		},
		isAuthorized: true,
	}
	//
	//canNotCreatePublicContracts = &testCase{
	//	msg: "0xb1b1b1 can not create public contracts",
	//	granted: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
	//		"public://0x0000000000000000000000000000000000b1b1b1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	ask: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: VisibilityPublic,
	//		Action:     ActionCreate,
	//	}},
	//	isAuthorized: false,
	//}
	canReadOwnedPublicContracts = &testCase{
		msg: "0x0a1a1a1 can read public contracts created by self",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: VisibilityPublic,
			Action:     ActionRead,
		}},
		isAuthorized: true,
	}
	canReadOtherPublicContracts = &testCase{
		msg: "0x0a1a1a1 can read public contracts created by 0xb1b1b1",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPublic,
			Action:     ActionRead,
		}},
		isAuthorized: true,
	}
	//canNotReadOtherPublicContracts = &testCase{
	//	msg: "0x0a1a1a1 can only read public contracts created by self",
	//	granted: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	ask: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xa1a1a1"),
	//			To:   common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: VisibilityPublic,
	//		Action:     ActionRead,
	//	}},
	//	isAuthorized: false,
	//}
	canWriteOwnedPublicContracts = &testCase{
		msg: "0x0a1a1a1 can send transactions to public contracts created by self",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: VisibilityPublic,
			Action:     ActionWrite,
		}},
		isAuthorized: true,
	}
	canWriteOtherPublicContracts1 = &testCase{
		msg: "0xa1a1a1 can send transactions to public contracts created by 0xb1b1b1",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPublic,
			Action:     ActionWrite,
		}},
		isAuthorized: true,
	}
	canWriteOtherPublicContracts2 = &testCase{
		msg: "0xa1a1a1 can send transactions to public contracts created by 0xb1b1b1",
		granted: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xc1c1c1"),
			},
			Visibility: VisibilityPublic,
			Action:     ActionWrite,
		}},
		isAuthorized: true,
	}
	//canNotWriteOtherPublicContracts = &testCase{
	//	msg: "0x0a1a1a1 can only send transactions to public contracts created by self",
	//	granted: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//		"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	ask: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xa1a1a1"),
	//			To:   common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: VisibilityPublic,
	//		Action:     ActionWrite,
	//	}},
	//	isAuthorized: false,
	//}
	// private contracts
	canCreatePrivateContracts = &testCase{
		msg: "0x0a1a1a1 can create private contracts with sender key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=A",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  VisibilityPrivate,
			Action:      ActionCreate,
			PrivateFrom: "A",
			Parties:     []string{},
		}},
		isAuthorized: true,
	}
	canNotCreatePrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT create private contracts with sender key A if only own key B",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  VisibilityPrivate,
			Action:      ActionCreate,
			PrivateFrom: "A",
			Parties:     []string{},
		}},
		isAuthorized: false,
	}
	canReadOwnedPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can read private contracts created by self and was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1&from.tm=A&from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionRead,
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canReadOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can read private contracts created by 0xb1b1b1 and was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=A",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionRead,
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canNotReadOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT read private contracts created by 0xb1b1b1 even it was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000c1c1c1&from.tm=A",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionRead,
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canNotReadOtherPrivateContractsNoPrivy = &testCase{
		msg: "0x0a1a1a1 can NOT read private contracts created by 0xb1b1b1 as it was privy to a key B",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionRead,
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canWriteOwnedPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can write private contracts created by self and was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1&from.tm=A&from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  VisibilityPrivate,
			Action:      ActionWrite,
			PrivateFrom: "A",
			Parties:     []string{"A"},
		}},
		isAuthorized: true,
	}
	canWriteOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can write private contracts created by 0xb1b1b1 and was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=A",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility:  VisibilityPrivate,
			Action:      ActionWrite,
			PrivateFrom: "A",
			Parties:     []string{"A"},
		}},
		isAuthorized: true,
	}
	canWriteOtherPrivateContractsWithOverlappedScope = &testCase{
		msg: "0x0a1a1a1 can write private contracts created by 0xb1b1b1 and was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=A",
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=A&from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility:  VisibilityPrivate,
			Action:      ActionWrite,
			PrivateFrom: "A",
			Parties:     []string{"A"},
		}},
		isAuthorized: true,
	}
	canNotWriteOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT write private contracts created by 0xb1b1b1 even it was privy to a key A",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000c1c1c1&from.tm=A",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionWrite,
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canNotWriteOtherPrivateContractsNoPrivy = &testCase{
		msg: "0x0a1a1a1 can NOT write private contracts created by 0xb1b1b1 as it was privy to a key B",
		granted: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&from.tm=B",
		},
		ask: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: VisibilityPrivate,
			Action:     ActionWrite,
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
)
