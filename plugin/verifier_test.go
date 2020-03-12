package plugin

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVerifier_whenResolvingDefaultPublicKeyLocation(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	if err := ioutil.WriteFile(path.Join(tmpDir, DefaultPublicKeyFile), []byte("foo"), 0644); err != nil {
		t.Fatal(err)
	}
	arbitraryPM := &PluginManager{
		pluginBaseDir: tmpDir,
	}

	testObject, err := NewVerifier(arbitraryPM, true, "")

	assert.NoError(t, err)
	assert.IsType(t, &LocalVerifier{}, testObject)
}

func TestNewVerifier_whenUsingOnlineVerifier(t *testing.T) {
	arbitraryPM := &PluginManager{}

	testObject, err := NewVerifier(arbitraryPM, false, "")

	assert.NoError(t, err)
	assert.IsType(t, &OnlineVerifier{}, testObject)
}
