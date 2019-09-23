package security

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
	msg            string
	rawAuthorities []string
	attributes     []*ContractSecurityAttribute
	isAuthorized   bool
}

func TestMatch_whenTypical(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/create/contracts?from.tm=A/&for.tm=B&for.tm=C")
	ask, _ := url.Parse("private://0xa1b1c1/create/contracts?for.tm=B&from.tm=A%2F")

	assert.True(t, match(nil, ask, granted))
}

func TestMatch_whenAnyAction(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/_/contracts?owned.eoa=0x0&party.tm=A1")
	ask, _ := url.Parse("private://0xa1b1c1/create/contracts?for.tm=B&from.tm=A1")

	assert.True(t, match(nil, ask, granted))

	ask, _ = url.Parse("private://0xa1b1c1/read/contracts?owned.eoa=0x0&party.tm=A1&party.tm=B1")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "read",
	}, ask, granted))

	ask, _ = url.Parse("private://0xa1b1c1/write/contracts?owned.eoa=0x0&party.tm=A1&party.tm=B1")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "write",
	}, ask, granted))
}

func TestMatch_whenContractWritePermission_AskIsTheSuperSet(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&party.tm=A")
	ask, _ := url.Parse("private://0xa1b1c1/write/contracts?owned.eoa=0xc1d1e1&party.tm=A&party.tm=B")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "write",
	}, ask, granted), "with write permission")

	granted, _ = url.Parse("private://0x0/read/contracts?owned.eoa=0x0&party.tm=A")
	ask, _ = url.Parse("private://0xa1b1c1/read/contracts?owned.eoa=0xc1d1e1&party.tm=A&party.tm=B")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "read",
	}, ask, granted), "with read permission")
}

func TestMatch_whenContractWritePermission_GrantedIsTheSuperSet(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&party.tm=A&party.tm=B")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&party.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "write",
	}, ask, granted), "with write permission")

	granted, _ = url.Parse("private://0x0/read/contracts?owned.eoa=0x0&party.tm=A&party.tm=B")
	ask, _ = url.Parse("private://0x0/read/contracts?owned.eoa=0x0&party.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "read",
	}, ask, granted), "with read permission")
}

func TestMatch_whenContractWritePermission_Same(t *testing.T) {
	granted, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&party.tm=A")
	ask, _ := url.Parse("private://0x0/write/contracts?owned.eoa=0x0&party.tm=A")

	assert.True(t, match(&ContractSecurityAttribute{
		Visibility: "private",
		Action:     "write",
	}, ask, granted))
}

func TestMatch_whenUsingWildcardAccount(t *testing.T) {
	granted, _ := url.Parse("private://0x0/create/contracts?from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D&for.tm=EiEm6FntRU6LaD0gNIZbDKC4HcwYvl3c2XoViPur%2BxM%3D")
	ask, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?for.tm=EiEm6FntRU6LaD0gNIZbDKC4HcwYvl3c2XoViPur%2BxM%3D&from.tm=dLHrFQpbSda0EhJnLonsBwDjks%2Bf724NipfI5zK5RSs%3D")

	assert.True(t, match(nil, ask, granted))

	granted, _ = url.Parse("private://0x0/read/contract?owned.eoa=0x0")
	ask, _ = url.Parse("private://0xa1b1c1/read/contract?owned.eoa=0x1234")

	assert.True(t, match(nil, ask, granted))
}

func TestMatch_whenPublic(t *testing.T) {
	granted, _ := url.Parse("private://0xa1b1c1/create/contract?from.tm=A/&for.tm=B&for.tm=C")
	ask, _ := url.Parse("public://0x0/create/contract")

	assert.True(t, match(nil, ask, granted))
}

func TestMatch_whenNotEscaped(t *testing.T) {
	// query not escaped probably in the granted authority resource identitifer
	granted, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?from.tm=BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=&for.tm=BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=")
	ask, _ := url.Parse("private://0xed9d02e382b34818e88b88a309c7fe71e65f419d/create/contracts?for.tm=BULeR8JyUWhiuuCMU%2FHLA0Q5pzkYT%2BcHII3ZKBey3Bo%3D&from.tm=BULeR8JyUWhiuuCMU%2FHLA0Q5pzkYT%2BcHII3ZKBey3Bo%3D")

	assert.False(t, match(nil, ask, granted))
}

func runTestCases(t *testing.T, testCases []*testCase) {
	testObject := &DefaultContractAccessDecisionManager{}
	for _, tc := range testCases {
		log.Debug("--> Running test case: " + tc.msg)
		authorities := make([]*proto.GrantedAuthority, 0)
		for _, a := range tc.rawAuthorities {
			authorities = append(authorities, &proto.GrantedAuthority{Raw: a})
		}
		b, err := testObject.IsAuthorized(
			context.Background(),
			&proto.PreAuthenticatedAuthenticationToken{Authorities: authorities},
			tc.attributes)
		assert.NoError(t, err, tc.msg)
		assert.Equal(t, tc.isAuthorized, b, tc.msg)
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
		canCreatePrivateContracts2,
		canNotCreatePrivateContracts,
		canReadOwnedPrivateContracts,
		canReadOtherPrivateContracts,
		canNotReadOtherPrivateContracts,
		canNotReadOtherPrivateContractsNoPrivy,
		canWriteOwnedPrivateContracts,
		canWriteOtherPrivateContracts,
		canNotWriteOtherPrivateContracts,
		canNotWriteOtherPrivateContractsNoPrivy,
	})
}

var (
	canCreatePublicContracts = &testCase{
		msg: "0x0a1a1a1 can create public contracts",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
		},
		attributes: []*ContractSecurityAttribute{
			{
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
				},
				Visibility: "public",
				Action:     "create",
			},
		},
		isAuthorized: true,
	}
	canCreatePublicContractsAndWriteToOthers = &testCase{
		msg: "0x0a1a1a1 can create public contracts and write to contracts created by 0xb1b1b1",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		attributes: []*ContractSecurityAttribute{
			{
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
				},
				Visibility: "public",
				Action:     "create",
			}, {
				AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
					From: common.HexToAddress("0xa1a1a1"),
					To:   common.HexToAddress("0xb1b1b1"),
				},
				Visibility: "public",
				Action:     "write",
			},
		},
		isAuthorized: true,
	}
	//
	//canNotCreatePublicContracts = &testCase{
	//	msg: "0xb1b1b1 can not create public contracts",
	//	rawAuthorities: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/create/contracts",
	//		"public://0x0000000000000000000000000000000000b1b1b1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	attributes: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: "public",
	//		Action:     "create",
	//	}},
	//	isAuthorized: false,
	//}
	canReadOwnedPublicContracts = &testCase{
		msg: "0x0a1a1a1 can read public contracts created by self",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: "public",
			Action:     "read",
		}},
		isAuthorized: true,
	}
	canReadOtherPublicContracts = &testCase{
		msg: "0x0a1a1a1 can read public contracts created by 0xb1b1b1",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "public",
			Action:     "read",
		}},
		isAuthorized: true,
	}
	//canNotReadOtherPublicContracts = &testCase{
	//	msg: "0x0a1a1a1 can only read public contracts created by self",
	//	rawAuthorities: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	attributes: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xa1a1a1"),
	//			To:   common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: "public",
	//		Action:     "read",
	//	}},
	//	isAuthorized: false,
	//}
	canWriteOwnedPublicContracts = &testCase{
		msg: "0x0a1a1a1 can send transactions to public contracts created by self",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: "public",
			Action:     "write",
		}},
		isAuthorized: true,
	}
	canWriteOtherPublicContracts1 = &testCase{
		msg: "0xa1a1a1 can send transactions to public contracts created by 0xb1b1b1",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "public",
			Action:     "write",
		}},
		isAuthorized: true,
	}
	canWriteOtherPublicContracts2 = &testCase{
		msg: "0xa1a1a1 can send transactions to public contracts created by 0xb1b1b1",
		rawAuthorities: []string{
			"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&owned.eoa=0x0000000000000000000000000000000000c1c1c1",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xc1c1c1"),
			},
			Visibility: "public",
			Action:     "write",
		}},
		isAuthorized: true,
	}
	//canNotWriteOtherPublicContracts = &testCase{
	//	msg: "0x0a1a1a1 can only send transactions to public contracts created by self",
	//	rawAuthorities: []string{
	//		"public://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//		"public://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1",
	//	},
	//	attributes: []*ContractSecurityAttribute{{
	//		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
	//			From: common.HexToAddress("0xa1a1a1"),
	//			To:   common.HexToAddress("0xb1b1b1"),
	//		},
	//		Visibility: "public",
	//		Action:     "write",
	//	}},
	//	isAuthorized: false,
	//}
	// private contracts
	canCreatePrivateContracts = &testCase{
		msg: "0x0a1a1a1 can create private contracts with sender key A to receiver keys B and C",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=A&for.tm=B&for.tm=C",
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=A&for.tm=D",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  "private",
			Action:      "create",
			PrivateFrom: "A",
			Parties:     []string{"B", "C"},
		}},
		isAuthorized: true,
	}
	canCreatePrivateContracts2 = &testCase{
		msg: "0x0a1a1a1 can create private contracts with sender key A to receiver key B",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=A&for.tm=B&for.tm=C",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  "private",
			Action:      "create",
			PrivateFrom: "A",
			Parties:     []string{"B"},
		}},
		isAuthorized: true,
	}
	canNotCreatePrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT create private contracts with sender key A to receiver keys (B, C and D)",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/create/contracts?from.tm=A&for.tm=B&for.tm=C",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility:  "private",
			Action:      "create",
			PrivateFrom: "A",
			Parties:     []string{"B", "C", "D"},
		}},
		isAuthorized: false,
	}
	canReadOwnedPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can read private contracts created by self and was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1&party.tm=A&party.tm=B",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: "private",
			Action:     "read",
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canReadOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can read private contracts created by 0xb1b1b1 and was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&party.tm=A",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "read",
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canNotReadOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT read private contracts created by 0xb1b1b1 even it was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000c1c1c1&party.tm=A",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "read",
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canNotReadOtherPrivateContractsNoPrivy = &testCase{
		msg: "0x0a1a1a1 can NOT read private contracts created by 0xb1b1b1 as it was privy to a key B",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/read/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&party.tm=B",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "read",
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canWriteOwnedPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can write private contracts created by self and was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000a1a1a1&party.tm=A&party.tm=B",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
			},
			Visibility: "private",
			Action:     "write",
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canWriteOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can write private contracts created by 0xb1b1b1 and was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&party.tm=A",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "write",
			Parties:    []string{"A"},
		}},
		isAuthorized: true,
	}
	canNotWriteOtherPrivateContracts = &testCase{
		msg: "0x0a1a1a1 can NOT write private contracts created by 0xb1b1b1 even it was privy to a key A",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000c1c1c1&party.tm=A",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "write",
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
	canNotWriteOtherPrivateContractsNoPrivy = &testCase{
		msg: "0x0a1a1a1 can NOT write private contracts created by 0xb1b1b1 as it was privy to a key B",
		rawAuthorities: []string{
			"private://0x0000000000000000000000000000000000a1a1a1/write/contracts?owned.eoa=0x0000000000000000000000000000000000b1b1b1&party.tm=B",
		},
		attributes: []*ContractSecurityAttribute{{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
				From: common.HexToAddress("0xa1a1a1"),
				To:   common.HexToAddress("0xb1b1b1"),
			},
			Visibility: "private",
			Action:     "write",
			Parties:    []string{"A"},
		}},
		isAuthorized: false,
	}
)
