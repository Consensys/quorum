package multitenancy

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

// AccountAccessDecisionManager performs authorization checks for Ethereum Account
// based on what is entitled in the proto.PreAuthenticatedAuthenticationToken
// and what is asked in ContractSecurityAttribute list.
// Note: place holder for future, this is to protect Value Transfer between accounts.
type AccountAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attr *AccountStateSecurityAttribute) (bool, error)
}

// ContractAccessDecisionManager performs authorization checks for contract
// based on what is entitled in the proto.PreAuthenticatedAuthenticationToken
// and what is asked in ContractSecurityAttribute list.
type ContractAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attributes []*ContractSecurityAttribute) (bool, error)
}

type DefaultContractAccessDecisionManager struct {
}

// IsAuthorized performs authorization check for each security attribute against
// the granted access inside the pre-authenticated access token.
//
// All security attributes must pass.
func (cm *DefaultContractAccessDecisionManager) IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attributes []*ContractSecurityAttribute) (bool, error) {
	matchCount := 0
	if len(attributes) == 0 {
		return false, nil
	}
	for _, attr := range attributes {
		query := url.Values{}
		switch attr.Visibility {
		case VisibilityPublic:
			switch attr.Action {
			case ActionRead, ActionWrite, ActionCreate:
				if (attr.To == common.Address{}) {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.From.Hex()))
				} else {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.To.Hex()))
				}
			}
		case VisibilityPrivate:
			switch attr.Action {
			case ActionRead:
				if (attr.To == common.Address{}) {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.From.Hex()))
				} else {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.To.Hex()))
				}
				for _, tm := range attr.Parties {
					query.Add(QueryFromTM, tm)
				}
			case ActionWrite:
				if (attr.To == common.Address{}) {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.From.Hex()))
				} else {
					query.Set(QueryOwnedEOA, strings.ToLower(attr.To.Hex()))
				}
				query.Add(QueryFromTM, attr.PrivateFrom)
			case ActionCreate:
				query.Set(QueryFromTM, attr.PrivateFrom)
			}
		}
		// construct request permission identifier
		request, err := url.Parse(fmt.Sprintf("%s://%s/%s/%s?%s", attr.Visibility, strings.ToLower(attr.From.Hex()), attr.Action, "contracts", query.Encode()))
		if err != nil {
			return false, err
		}
		// compare the contract security attribute with the consolidate list
		for _, granted := range authToken.GetAuthorities() {
			pi, err := url.Parse(granted.GetRaw())
			if err != nil {
				continue
			}
			granted := pi.String()
			ask := request.String()
			isMatched := match(attr, request, pi)
			log.Debug("Checking contract access", "passed", isMatched, "granted", granted, "ask", ask)
			if isMatched {
				matchCount++
				break
			}
		}
	}
	return matchCount == len(attributes), nil
}

func match(attr *ContractSecurityAttribute, ask, granted *url.URL) bool {
	askScheme := strings.ToLower(ask.Scheme)
	if allowedPublic(askScheme) {
		return true
	}

	isPathMatched := matchPath(strings.ToLower(ask.Path), strings.ToLower(granted.Path))
	return askScheme == strings.ToLower(granted.Scheme) && //Note: "askScheme" here is "private" since we checked VisibilityPublic above.
		matchHost(strings.ToLower(ask.Host), strings.ToLower(granted.Host)) && //whether i have permission to execute using this ethereum address
		isPathMatched && //is our permission for the same action (read, write, deploy)
		matchQuery(attr, ask.Query(), granted.Query())
}

func allowedPublic(scheme string) bool {
	return scheme == string(VisibilityPublic)
}

func matchHost(ask string, granted string) bool {
	return granted == AnyEOAAddress || ask == granted
}

func matchPath(ask string, granted string) bool {
	return strings.HasPrefix(granted, "/_") || ask == granted
}

func matchQuery(attr *ContractSecurityAttribute, ask, granted url.Values) bool {
	// if asking nothing, we should bail out
	if len(ask) == 0 || len(ask[QueryFromTM]) == 0 {
		return false
	}
	// possible scenarios:
	// 1. read/write -> from.tm -> at least 1 of the same key must appear in both lists
	// 2. read/write - owned.eoa/to.eoa -> check subset
	// 3. create -> from.tm/owned.eoa/to.eoa -> check subset
	for k, askValues := range ask {
		grantedValues := granted[k]
		switch attr.Action {
		case ActionRead, ActionWrite:
			// Scenario 1
			if k == QueryFromTM {
				if isIntersectionEmpty(grantedValues, askValues) {
					return false
				}
			}
			//Scenario 2
			if k == QueryOwnedEOA || k == QueryToEOA {
				if !subset(grantedValues, askValues) {
					return false
				}
			}
		case ActionCreate:
			//Scenario 3
			if !subset(grantedValues, askValues) {
				return false
			}
		default:
			// we don't know, better reject
			log.Error("unsupported action", "action", attr.Action)
			return false
		}
	}
	return true
}

func subset(grantedValues, askValues []string) bool {
	for _, askValue := range askValues {
		found := false
		sanitizedAskValue := askValue
		if strings.HasPrefix(askValue, "0x") {
			sanitizedAskValue = strings.ToLower(askValue)
		}
		for _, grantedValue := range grantedValues {
			sanitizedGrantedValue := grantedValue
			if strings.HasPrefix(grantedValue, "0x") {
				sanitizedGrantedValue = strings.ToLower(grantedValue)
			}
			if sanitizedGrantedValue == AnyEOAAddress || sanitizedAskValue == sanitizedGrantedValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func isIntersectionEmpty(grantedValues, askValues []string) bool {
	grantedMap := make(map[string]bool)
	for _, grantedVal := range grantedValues {
		grantedMap[grantedVal] = true
	}
	for _, askVal := range askValues {
		if grantedMap[askVal] {
			return false
		}
	}
	return true
}
