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

type AccountAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attr *AccountStateSecurityAttribute) (bool, error)
}

type ContractAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attributes []*ContractSecurityAttribute) (bool, error)
}

type DefaultAccountAccessDecisionManager struct {
}

func (am *DefaultAccountAccessDecisionManager) IsAuthorized(ctx context.Context,
	authToken *proto.PreAuthenticatedAuthenticationToken, attr *AccountStateSecurityAttribute) (bool, error) {
	panic("implement me")
}

type DefaultContractAccessDecisionManager struct {
}

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
			case ActionRead, ActionWrite:
				if (attr.To == common.Address{}) {
					query.Set("owned.eoa", strings.ToLower(attr.From.Hex()))
				} else {
					query.Set("owned.eoa", strings.ToLower(attr.To.Hex()))
				}
			}
		case VisibilityPrivate:
			switch attr.Action {
			case ActionRead:
				if (attr.To == common.Address{}) {
					query.Set("owned.eoa", strings.ToLower(attr.From.Hex()))
				} else {
					query.Set("owned.eoa", strings.ToLower(attr.To.Hex()))
				}
				for _, tm := range attr.Parties {
					query.Add("from.tm", tm)
				}
			case ActionWrite:
				if (attr.To == common.Address{}) {
					query.Set("owned.eoa", strings.ToLower(attr.From.Hex()))
				} else {
					query.Set("owned.eoa", strings.ToLower(attr.To.Hex()))
				}
				query.Set("from.tm", attr.PrivateFrom)
			case ActionCreate:
				query.Set("from.tm", attr.PrivateFrom)
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
			log.Debug("Checking contract access", "granted", granted, "with", ask)
			if match(attr, request, pi) {
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
	return granted == "0x0" || ask == granted
}

func matchPath(ask string, granted string) bool {
	return strings.HasPrefix(granted, "/_") || ask == granted
}

func matchQuery(attr *ContractSecurityAttribute, ask, granted url.Values) bool {
	// possible scenarios:
	// 1. read -> from.tm -> at least 1 of the same key must appear in both lists
	// 2. read - owned.eoa/to.eoa -> check subset
	// 3. write/create -> from.tm/owned.eoa/to.eoa -> check subset

	for k, askValues := range ask {
		grantedValues := granted[k]
		if attr.Action == ActionRead {
			// Scenario 1
			if k == "from.tm" {
				if isIntersectionEmpty(grantedValues, askValues) {
					return false
				}
			}
			//Scenario 2
			if k == "owned.eoa" || k == "to.eoa" {
				if !subset(grantedValues, askValues) {
					return false
				}
			}
		} else {
			//action is "write" or "create"

			//Scenario 3
			if !subset(grantedValues, askValues) {
				return false
			}
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
			if sanitizedGrantedValue == "0x0" || sanitizedAskValue == sanitizedGrantedValue {
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
