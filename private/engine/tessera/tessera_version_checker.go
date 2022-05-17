package tessera

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/private/engine"
)

const versionLength = 3

type Version [versionLength]uint64

var (
	zero                         = Version{0, 0, 0}
	privacyEnhancementsVersion   = Version{2, 0, 0}
	multitenancyVersion          = Version{2, 1, 0}
	multiplePrivateStatesVersion = Version{3, 0, 0}
	mandatoryRecipientsVersion   = Version{4, 0, 0}

	featureVersions = map[engine.PrivateTransactionManagerFeature]Version{
		engine.PrivacyEnhancements:   privacyEnhancementsVersion,
		engine.MultiTenancy:          multitenancyVersion,
		engine.MultiplePrivateStates: multiplePrivateStatesVersion,
		engine.MandatoryRecipients:   mandatoryRecipientsVersion,
	}
)

func tesseraVersionFeatures(version Version) []engine.PrivateTransactionManagerFeature {
	result := make([]engine.PrivateTransactionManagerFeature, 0)
	for feature, featureVersion := range featureVersions {
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
func compareVersions(v1, v2 Version) int {
	for i := 0; i < versionLength; i++ {
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
func parseVersion(version []byte) (res Version, err error) {
	versionMajMidRegExp, _ := regexp.Compile(`([0-9]+)\.([0-9]+)([^0-9].*)?`)
	versionMajMidMinRegExp, _ := regexp.Compile(`([0-9]+)\.([0-9]+)\.([0-9]+)([^0-9].*)?`)

	var submatch [][]byte
	if versionMajMidMinRegExp.Match(version) {
		submatch = versionMajMidMinRegExp.FindSubmatch(version)[1:4]
	} else if versionMajMidRegExp.Match(version) {
		submatch = versionMajMidRegExp.FindSubmatch(version)[1:3]
	} else {
		return zero, fmt.Errorf("input does not match the expected version pattern")
	}

	// res should be initialized with {0,0,0} - thus it is ok for submatch to have a variable length of 2 or 3
	for idx, val := range submatch {
		intVal, err := strconv.ParseUint(string(val), 10, 64)
		if err != nil {
			return zero, err
		}
		res[idx] = intVal
	}
	return res, nil
}
