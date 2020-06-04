package tessera

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/private/engine"
)

const VERSION_LENGTH = 3

type VERSION [VERSION_LENGTH]int64

var (
	ZERO = VERSION{0, 0, 0}
	// TODO Qurum - Privacy Enhancements - must update this once tessera with privacy enhancements is released (and the version is known)
	PRIVACY_ENHANCEMENTS_VERSION = VERSION{0, 10, 6}

	FEATURE_VERSIONS = map[engine.PrivateTransactionManagerFeature]VERSION{
		engine.PrivacyEnhancements: PRIVACY_ENHANCEMENTS_VERSION,
	}
)

func tesseraVersionFeatures(version VERSION) []engine.PrivateTransactionManagerFeature {
	result := make([]engine.PrivateTransactionManagerFeature, 0)
	for feature, featureVersion := range FEATURE_VERSIONS {
		if compareVersions(version, featureVersion) >= 0 {
			result = append(result, feature)
		}
	}
	return result
}

// compare two versions
// if v1 > v2 - returns 1
// if v1 < v2 - returns -1
// if v1 = v2 - returns 0
func compareVersions(v1, v2 VERSION) int {
	for i := 0; i < VERSION_LENGTH; i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	return 0
}

// The tessera release versions have 3 components: major.mid.minor.
// Snapshot tessera builds may have versions made of 2 components: major.mid-SNAPSHOT.
// parseVersion will assume the minor version to be 0 for versions with only 2 components.
func parseVersion(version []byte) (res VERSION, err error) {
	versionMajMidRegExp, _ := regexp.Compile(`([0-9]+)\.([0-9]+)([^0-9].*)`)
	versionMajMidMinRegExp, _ := regexp.Compile(`([0-9]+)\.([0-9]+)\.([0-9]+)([^0-9].*)`)

	var submatch [][]byte
	if versionMajMidMinRegExp.Match(version) {
		submatch = versionMajMidMinRegExp.FindSubmatch(version)[1:4]
	} else if versionMajMidRegExp.Match(version) {
		submatch = versionMajMidRegExp.FindSubmatch(version)[1:3]
	} else {
		return ZERO, fmt.Errorf("input does not match the expected version pattern")
	}

	// res should be initialized with {0,0,0} - thus it is ok for submatch to have a variable length of 2 or 3
	for idx, val := range submatch {
		intVal, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return ZERO, err
		}
		res[idx] = intVal
	}
	return res, nil
}
