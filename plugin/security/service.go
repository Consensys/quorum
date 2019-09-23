package security

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type TLSConfigurationSource interface {
	Get(ctx context.Context) (*tls.Config, error)
}

type AuthenticationManager interface {
	Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error)
	IsEnabled(ctx context.Context) (bool, error)
}

// TODO plugin: define as proto
// Security configuration attributes which are defined for secure object.
// These attributes have special meaning to various implementations

// For secure account state
type AccountStateSecurityAttribute struct {
	From common.Address // Account Address
	To   common.Address
}

type ContractSecurityAttribute struct {
	*AccountStateSecurityAttribute
	Visibility  string   // public/private
	Action      string   // create/read/write
	PrivateFrom string   // TM Key, only if Visibility is private
	Parties     []string // TM Keys, only if Visibility is private
}

type AccountAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attr *AccountStateSecurityAttribute) (bool, error)
}

type ContractAccessDecisionManager interface {
	IsAuthorized(ctx context.Context, authToken *proto.PreAuthenticatedAuthenticationToken, attributes []*ContractSecurityAttribute) (bool, error)
}

type AuthenticationManagerDeferFunc func() (AuthenticationManager, error)

type DeferredAuthenticationManager struct {
	deferFunc AuthenticationManagerDeferFunc
}

func (d *DeferredAuthenticationManager) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	am, err := d.deferFunc()
	if err != nil {
		return nil, err
	}
	return am.Authenticate(ctx, token)
}

func (d *DeferredAuthenticationManager) IsEnabled(ctx context.Context) (bool, error) {
	am, err := d.deferFunc()
	if err != nil {
		return false, err
	}
	return am.IsEnabled(ctx)
}

func NewDeferredAuthenticationManager(deferFunc AuthenticationManagerDeferFunc) *DeferredAuthenticationManager {
	return &DeferredAuthenticationManager{
		deferFunc: deferFunc,
	}
}

type DisabledAuthenticationManager struct {
}

func (*DisabledAuthenticationManager) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	return nil, errors.New("not supported operation")
}

func (*DisabledAuthenticationManager) IsEnabled(ctx context.Context) (bool, error) {
	return false, nil
}

func NewDisabledAuthenticationManager() AuthenticationManager {
	return &DisabledAuthenticationManager{}
}

// TODO plugin: the below default implementations need to move to the plugin code

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
		case "public":
			switch attr.Action {
			case "read", "write":
				if (attr.To == common.Address{}) {
					query.Set("owned.eoa", strings.ToLower(attr.From.Hex()))
				} else {
					query.Set("owned.eoa", strings.ToLower(attr.To.Hex()))
				}
			}
		case "private":
			switch attr.Action {
			case "read", "write":
				if (attr.To == common.Address{}) {
					query.Set("owned.eoa", strings.ToLower(attr.From.Hex()))
				} else {
					query.Set("owned.eoa", strings.ToLower(attr.To.Hex()))
				}
				for _, tm := range attr.Parties {
					query.Add("party.tm", tm)
				}
			case "create":
				query.Set("from.tm", attr.PrivateFrom)
				for _, tm := range attr.Parties {
					query.Add("for.tm", tm)
				}
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
	isPathMatched, isAnyAction := matchPath(strings.ToLower(ask.Path), strings.ToLower(granted.Path))
	return allowedPublic(askScheme) ||
		(askScheme == strings.ToLower(granted.Scheme) &&
			matchHost(strings.ToLower(ask.Host), strings.ToLower(granted.Host)) &&
			isPathMatched &&
			matchQuery(isAnyAction, attr, ask.Query(), granted.Query()))
}

func allowedPublic(scheme string) bool {
	return scheme == "public"
}

func matchHost(ask string, granted string) bool {
	if strings.HasPrefix(ask, "0x") && strings.HasPrefix(granted, "0x") && granted == "0x0" {
		return true
	}
	return ask == granted
}

func matchPath(ask string, granted string) (bool, bool) {
	if strings.HasPrefix(granted, "/_") {
		return true, true
	}
	return ask == granted, false
}

func matchQuery(isAnyAction bool, attr *ContractSecurityAttribute, ask, granted url.Values) bool {
	for k, askValues := range ask {
		grantedValues := granted[k]
		if isAnyAction {
			if k == "for.tm" {
				continue
			}
			if k == "from.tm" {
				grantedValues = granted["party.tm"]
			}
		}
		if attr != nil && attr.Visibility == "private" && (attr.Action == "write" || attr.Action == "read") && k == "party.tm" {
			if !subset(askValues, grantedValues) && !subset(grantedValues, askValues) {
				return false
			}
		} else {
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
