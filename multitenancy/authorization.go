package multitenancy

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

var (
	ErrNotAuthorized    = errors.New("not authorized")
	ErrPSIFoundMultiple = errors.New("found multiple authorized private state identifiers")
	ErrPSINotFound      = errors.New("no private state identifiers found")
)

// Authorizev performs authorization check for one security attribute against
// the granted access inside the pre-authenticated access token.
func Authorize(authToken *proto.PreAuthenticatedAuthenticationToken, attr *PrivateStateSecurityAttribute) (bool, error) {
	query := url.Values{}
	query.Set(QueryEOA, toHexAddress(attr.eoa))
	// construct the request
	askValue, err := url.Parse(fmt.Sprintf("%s://%s?%s", SchemePSI, attr.psi, query.Encode()))
	if err != nil {
		return false, err
	}
	// compare the security attribute with the granted list
	for _, granted := range authToken.GetAuthorities() {
		grantedValue, err := url.Parse(granted.GetRaw())
		if err != nil {
			continue
		}
		isMatched := match(askValue, grantedValue)
		log.Debug("Checking private state access", "passed", isMatched, "granted", grantedValue, "ask", askValue)
		if isMatched {
			return true, nil
		}
	}
	return false, nil
}

// ExtractPSI returns a single PSI if found in the granted scope.
// If there is none or multiple, return error
func ExtractPSI(authToken *proto.PreAuthenticatedAuthenticationToken) (types.PrivateStateIdentifier, error) {
	var (
		found         bool
		authorizedPSI types.PrivateStateIdentifier
	)
	for _, granted := range authToken.GetAuthorities() {
		grantedValue, err := url.Parse(granted.GetRaw())
		if err != nil || grantedValue.Scheme != SchemePSI {
			continue
		}
		if found {
			return "", ErrPSIFoundMultiple
		}
		found = true
		authorizedPSI = types.PrivateStateIdentifier(grantedValue.Host)
	}
	if !found {
		return "", ErrPSINotFound
	}
	return authorizedPSI, nil
}

func toHexAddress(a common.Address) string {
	if (a == common.Address{}) {
		return AnyEOAAddress
	}
	return strings.ToLower(a.Hex())
}

func match(ask, granted *url.URL) bool {
	return strings.EqualFold(ask.Scheme, granted.Scheme) &&
		strings.EqualFold(ask.Host, granted.Host) &&
		matchQuery(ask.Query(), granted.Query())
}

func matchQuery(ask, granted url.Values) bool {
	grantedEOAs := granted[QueryEOA]
	askEOAs := ask[QueryEOA]
	if len(askEOAs) == 0 && len(grantedEOAs) > 0 {
		return false
	}
	if len(grantedEOAs) == 0 { // consider AnyEOAAddress
		return true
	}
	if len(grantedEOAs) == 1 && strings.EqualFold(grantedEOAs[0], AnyEOAAddress) { // explicit
		return true
	}
	return common.ContainsAll(grantedEOAs, askEOAs)
}
