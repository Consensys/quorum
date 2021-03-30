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

// IsAuthorized performs authorization check for security attributes against
// the granted access inside the pre-authenticated access token.
func IsAuthorized(authToken *proto.PreAuthenticatedAuthenticationToken, secAttributes ...*PrivateStateSecurityAttribute) (bool, error) {
	for _, attr := range secAttributes {
		isAuthorized, err := isAuthorized(authToken, attr)
		if err != nil {
			return false, err
		}
		if !isAuthorized {
			return false, nil
		}
	}
	return true, nil
}

// isAuthorized performs authorization check for one security attribute against
// the granted access inside the pre-authenticated access token.
func isAuthorized(authToken *proto.PreAuthenticatedAuthenticationToken, attr *PrivateStateSecurityAttribute) (bool, error) {
	query := url.Values{}
	if attr.nodeEOA != nil {
		query.Set(QueryNodeEOA, toHexAddress(attr.nodeEOA))
	}
	if attr.selfEOA != nil {
		query.Set(QuerySelfEOA, toHexAddress(attr.selfEOA))
	}
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

// IsPSIAuthorized performs only authorization checks for PSI
func IsPSIAuthorized(authToken *proto.PreAuthenticatedAuthenticationToken, psi types.PrivateStateIdentifier) (bool, error) {
	// compare the security attribute with the granted list
	for _, granted := range authToken.GetAuthorities() {
		grantedValue, err := url.Parse(granted.GetRaw())
		if err != nil {
			continue
		}
		// because we care only for PSI so we try to match only PSI
		isMatched := strings.EqualFold(SchemePSI, grantedValue.Scheme) && strings.EqualFold(psi.String(), grantedValue.Host)
		log.Debug("Checking PSI access", "passed", isMatched, "granted", grantedValue, "ask", psi)
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
		grantedPSI := types.PrivateStateIdentifier(grantedValue.Host)
		// already captured
		if grantedPSI == authorizedPSI {
			continue
		}
		if found {
			return "", ErrPSIFoundMultiple
		}
		found = true
		authorizedPSI = grantedPSI
	}
	if !found {
		return "", ErrPSINotFound
	}
	return authorizedPSI, nil
}

func toHexAddress(a *common.Address) string {
	if a == nil {
		return ""
	}
	if (*a == common.Address{}) {
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
	return matchEOA(granted[QueryNodeEOA], ask[QueryNodeEOA]) || matchEOA(granted[QuerySelfEOA], ask[QuerySelfEOA])
}

func matchEOA(grantedEOAs []string, askEOAs []string) bool {
	if len(grantedEOAs) == 0 || len(askEOAs) == 0 {
		return false
	}
	return common.ContainsAll(grantedEOAs, []string{AnyEOAAddress}, askEOAs)
}
