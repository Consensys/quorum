package build

import (
	"os"
	"testing"

	testifyassert "github.com/stretchr/testify/assert"
)

func TestIgnorePackages_whenTypical(t *testing.T) {
	assert := testifyassert.New(t)

	arbitraryPackages := []string{"abc", "xyz/abc"}

	actual := IgnorePackages(arbitraryPackages)

	assert.Equal(arbitraryPackages, actual)
}

func TestIgnorePackages_whenIgnoreOnePackage(t *testing.T) {
	assert := testifyassert.New(t)

	arbitraryPackages := []string{"abc", "xyz/abc"}
	assert.NoError(os.Setenv("QUORUM_IGNORE_TEST_PACKAGES", "abc"))

	actual := IgnorePackages(arbitraryPackages)

	assert.Equal([]string{arbitraryPackages[1]}, actual)
}

func TestIgnorePackages_whenIgnorePackages(t *testing.T) {
	assert := testifyassert.New(t)

	arbitraryPackages := []string{"abc", "xyz/abc/opq"}
	assert.NoError(os.Setenv("QUORUM_IGNORE_TEST_PACKAGES", "abc, xyz/abc"))

	actual := IgnorePackages(arbitraryPackages)

	assert.Len(actual, 0)
}
