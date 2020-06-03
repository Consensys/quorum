package tessera

import (
	"testing"
)

func checkMajorMidSucceeds(t *testing.T, version []byte, major int64, middle int64) {
	maj, mid, err := getMajorMidFromVersion(version)
	if err != nil {
		t.Errorf("unexpected error")
	}
	if maj != major || mid != middle {
		t.Errorf("unexpected major or middle version missmatch")
	}
}

func TestGetMajorMidFromVersion(t *testing.T) {
	checkMajorMidSucceeds(t, []byte("0.11.12+12234"), 0, 11)
	checkMajorMidSucceeds(t, []byte("0.11-SNAPSHOT"), 0, 11)
	// leading zeros in major and mid components
	checkMajorMidSucceeds(t, []byte("000.0011-SNAPSHOT"), 0, 11)

	checkMajorMidSucceeds(t, []byte("01.12 SNAPSHOT"), 1, 12)

	_, _, err := getMajorMidFromVersion([]byte("garbage"))
	if err == nil {
		t.Errorf("expecting error to be returned when garbage version is supplied")
	}
}

func TestTesseraVersionSupportsPrivacyEnhancements(t *testing.T) {
	res := tesseraVersionSupportsPrivacyEnhancements([]byte("0.11.12+12234"))
	if !res {
		t.Errorf("supplied version should support privacy enhancements")
	}

	res = tesseraVersionSupportsPrivacyEnhancements([]byte("0.12-SNAPSHOT"))
	if !res {
		t.Errorf("supplied version should support privacy enhancements")
	}

	res = tesseraVersionSupportsPrivacyEnhancements([]byte("1.6.12"))
	if !res {
		t.Errorf("supplied version should support privacy enhancements")
	}

	res = tesseraVersionSupportsPrivacyEnhancements([]byte("0.10.5"))
	if res {
		t.Errorf("supplied version should not support privacy enhancements")
	}

	res = tesseraVersionSupportsPrivacyEnhancements([]byte("garbage"))
	if res {
		t.Errorf("supplied version should not support privacy enhancements")
	}
}
