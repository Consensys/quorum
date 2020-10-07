package tessera

import (
	"testing"

	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/stretchr/testify/assert"
)

func checkParseSucceeds(t *testing.T, version []byte, expectedVersion Version) {
	parsedVersion, err := parseVersion(version)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if compareVersions(parsedVersion, expectedVersion) != 0 {
		t.Errorf("unexpected major or middle version missmatch")
	}
}

func TestParseVersion(t *testing.T) {
	checkParseSucceeds(t, []byte("0.10.6"), Version{0, 10, 6})
	checkParseSucceeds(t, []byte("0.10-SNAPSHOT"), Version{0, 10, 0})
	checkParseSucceeds(t, []byte("0.10.1-SNAPSHOT"), Version{0, 10, 1})
	checkParseSucceeds(t, []byte("0.10.0-SNAPSHOT"), Version{0, 10, 0})
	checkParseSucceeds(t, []byte("0.11.12+12234"), Version{0, 11, 12})
	checkParseSucceeds(t, []byte("0.11-SNAPSHOT"), Version{0, 11, 0})
	// leading zeros in version components
	checkParseSucceeds(t, []byte("000.0011-SNAPSHOT"), Version{0, 11, 0})

	checkParseSucceeds(t, []byte("01.012 SNAPSHOT"), Version{1, 12, 0})

	_, err := parseVersion([]byte("garbage"))
	if err == nil {
		t.Errorf("expecting error to be returned when garbage version is supplied")
	}

	_, err = parseVersion([]byte("1.garbage"))
	if err == nil {
		t.Errorf("expecting error to be returned when garbage version is supplied")
	}
}

func TestVersionsComparison(t *testing.T) {
	v1 := Version{1, 1, 1}
	v2 := Version{1, 1, 1}
	v3 := Version{2, 1, 1}
	v4 := Version{1, 2, 1}
	v5 := Version{1, 1, 2}
	assert.Equal(t, 0, compareVersions(v1, v2), "versions should be equal")
	assert.Equal(t, -1, compareVersions(v1, v3), "v1 shold be smaller than v3")
	assert.Equal(t, 1, compareVersions(v3, v1), "v3 should be bigger than v1")
	assert.Equal(t, -1, compareVersions(v1, v4), "v1 shold be smaller than v4")
	assert.Equal(t, 1, compareVersions(v4, v1), "v4 should be bigger than v1")
	assert.Equal(t, -1, compareVersions(v1, v5), "v1 shold be smaller than v5")
	assert.Equal(t, 1, compareVersions(v5, v1), "v5 should be bigger than v1")
}

func TestTesseraVersionFeatures(t *testing.T) {
	res := tesseraVersionFeatures(Version{2, 11, 12})
	assert.Contains(t, res, engine.PrivacyEnhancements)
	res = tesseraVersionFeatures(Version{0, 12, 0})
	assert.NotContains(t, res, engine.PrivacyEnhancements)
	res = tesseraVersionFeatures(Version{0, 11, 15})
	assert.NotContains(t, res, engine.PrivacyEnhancements)
	res = tesseraVersionFeatures(Version{2, 0, 0})
	assert.Contains(t, res, engine.PrivacyEnhancements)
	res = tesseraVersionFeatures(Version{2, 1, 1})
	assert.Contains(t, res, engine.PrivacyEnhancements)
	res = tesseraVersionFeatures(zero)
	assert.NotContains(t, res, engine.PrivacyEnhancements)
	assert.Empty(t, res)
}
