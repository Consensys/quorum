package tessera

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
)

// TODO Qurum - Privacy Enhancements - must update these once tessera with privacy enhancements is released (and the version is known)
const PE_MAJOR_VER = 0
const PE_MID_VER = 11

func tesseraVersionFeatures(version []byte) []engine.PrivateTransactionManagerFeature {
	result := make([]engine.PrivateTransactionManagerFeature, 0)
	if tesseraVersionSupportsPrivacyEnhancements(version) {
		result = append(result, engine.PrivacyEnhancements)
	}
	return result
}

// The tessera release versions have 3 components: major.mid.minor.
// Snapshot tessera builds may have versions made of 2 components: major.mid-SNAPSHOT.
// As a result the tessera version check for privacy enhancements will only take into consideration the major and mid
// components of the version string.
func tesseraVersionSupportsPrivacyEnhancements(version []byte) bool {
	maj, mid, err := getMajorMidFromVersion(version)
	if err != nil {
		log.Error("Unable to extract major and mid components from tessera version: %s. Assuming tessera does not have privacy enhancements.", version)
		return false
	}
	if maj < PE_MAJOR_VER {
		return false
	} else if maj > PE_MAJOR_VER {
		return true
	}
	// major versions are equal - compare mid versions
	if mid < PE_MID_VER {
		return false
	}
	return true
}

func getMajorMidFromVersion(version []byte) (major int64, mid int64, err error) {
	versionRegExp, _ := regexp.Compile("([0-9]+)\\.([0-9]+)([^0-9].*)")

	submatch := versionRegExp.FindSubmatch(version)
	if len(submatch) < 4 {
		return 0, 0, fmt.Errorf("input does not match the expected version pattern")
	}

	major, err = strconv.ParseInt(string(submatch[1]), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	mid, err = strconv.ParseInt(string(submatch[2]), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return major, mid, nil
}
