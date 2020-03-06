package plugin

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTargetURL(t *testing.T) {
	assert.Error(t, isValidTargetURL("https://localhost.com", "http://localhost.com"))
	assert.Error(t, isValidTargetURL("https://localhost", "http://localhost.com"))

	if err := isValidTargetURL("http://localhost.com", "http://localhost.com"); err != nil {
		t.Errorf(err.Error())
	}

	if err := isValidTargetURL("https://localhost.com/../../", "https://localhost.com"); err != nil {
		t.Errorf(err.Error())
	}
}

func TestIsCleanFileName(t *testing.T) {
	assert.True(t, isCleanFileName("filename"), "filename is not valid")
	assert.True(t, isCleanFileName("filename.exe"), "filename with .exe")

	assert.False(t, isCleanFileName(""), "filename is not valid")
	assert.False(t, isCleanFileName("filename/"), "filename with /")
	assert.False(t, isCleanFileName("filename\\u00"), "filename with \\")
	assert.False(t, isCleanFileName("filename$"), "filename with $")
	assert.False(t, isCleanFileName("filename%"), "filename with %")
	assert.False(t, isCleanFileName("filename%00"), "filename with %")
}

func TestIsCleanEntryPoint(t *testing.T) {
	assert.True(t, isCleanEntryPoint("entrypoint"), "entrypoint is not valid")
	assert.True(t, isCleanEntryPoint("entrypoint.exe"), "entrypoint with .exe")

	assert.False(t, isCleanEntryPoint(""), "entrypoint is not valid")
	assert.False(t, isCleanEntryPoint("entrypoint/"), "entrypoint with /")
	assert.False(t, isCleanEntryPoint("entrypoint\\u00"), "entrypoint with \\")
	assert.False(t, isCleanEntryPoint("entrypoint$"), "entrypoint with $")
	assert.False(t, isCleanEntryPoint("entrypoint%"), "entrypoint with %")
	assert.False(t, isCleanEntryPoint("entrypoint%00"), "entrypoint with %")
}

func TestResolveFilePath_whenTypical(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	f, err := ioutil.TempFile(tmpDir, "f-")
	if err != nil {
		t.Fatal(err)
	}
	actualFile, err := resolveFilePath("file://" + f.Name())

	assert.NoError(t, err)
	assert.Equal(t, f.Name(), actualFile)
}

func TestResolveFilePath_whenInvalidFileURI(t *testing.T) {
	_, err := resolveFilePath("://arbitrary non uri")

	assert.Error(t, err)
}

func TestVerify_whenTypical(t *testing.T) {
	err := verify(validSignature, signerPubKey, arbitraryChecksum)

	assert.NoError(t, err)
}

func TestVerify_whenInvalid(t *testing.T) {
	err := verify(validSignature, arbitraryPubKey, arbitraryChecksum)

	assert.Error(t, err)
}

func TestUnpackPlugin_whenTypical(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	tmpZipFile, err := createArbitraryZip(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	workspace, meta, err := unpackPlugin(tmpZipFile)

	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(workspace)
	}()
	assert.NotEmpty(t, workspace)
	assert.NotNil(t, meta)
}

func createArbitraryZip(tmpDir string) (string, error) {
	tmpFile, err := ioutil.TempFile(tmpDir, "f-")
	if err != nil {
		return "", err
	}

	// Create a new zip archive.
	w := zip.NewWriter(tmpFile)
	defer func() {
		_ = w.Close()
	}()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"plugin-meta.json", `
{
	"name": "arbitrary-plugin",
	"version": "1.0.0",
	"entrypoint": "echo",
	"parameters": [
		"hello world"
	]
}
`},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return "", err
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return "", err
		}
	}

	return tmpFile.Name(), nil
}

var (
	arbitraryChecksum = "bf9a942afca462a9fb45f471f8d4db8c79cf332d"
	// signature of the signed arbitraryChecksum
	validSignature = []byte(`
-----BEGIN PGP SIGNATURE-----

iQJBBAEBCAArFiEEHGpboPTpUoYceZX2PIgUS38YTSgFAl1C/8ENHGFiY0B0ZXN0
LmNvbQAKCRA8iBRLfxhNKHeBEACs14x1+UoVEVNVDNSJORsQy6nthHiwrb5l66dW
KPcEt96y7KXJObSF7TWfmGjIgQXmDnrwMY78bKcbWVK90siDwA0SajUwmwmCbCeC
nMTIza1a64KblJRVGal9D5EWLdAOuQkAV2tddyWMqdvv2ef46y+2zmoKE3bOQLXj
sCi5e8myuh5ottfrf5Tkxi7QHrWICxYjAMEUkvke/jbYUFi1787VnHZ8LDG1x5WN
yz3KysyaraMiOstk5PcACU+bsvEIXFppJsgx9eNqdyfQ0/oMzKlqlHhss/W5osyq
LeVY9dcMXUSNGmB6deJde93pv3kYnLarhEM5Ovm5BxYMyzudk9hUy3wXyb51EPaL
z/hYViGpBVSwKY6q47s8duXruOA0TzYu5jYmJd+CzqBkDbJfh7JG9iJkdG4Q30ui
D2wvTBJfz6wu1qYj0semX4l4ntpJ6OcIvD0BpP1wz3eC9rt+3RzrjVWPbVoTOyB0
V7vVQPMJowoPvluIUP0eInc+jDue2Z/8DHjWDu1k4jmZbO7r/5Hib79JtA3LIGqq
CizH/cFWXLJAh6n7tFBREKgCgrsSQDIppdMFNc8GyRIh2qIkGexcWOBdiO6iU41t
anKh+gD3mcn637Hzn0p2AA0TK0D/HzPX9ZCwgGVyoQkXoMa1zWqzQ7QEbk8DmH9M
5rr3iw==
=4cZl
-----END PGP SIGNATURE-----
`)
	signerPubKey = []byte(`
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
=8vbI
-----END PGP PUBLIC KEY BLOCK-----
`)
	arbitraryPubKey = []byte(`
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBF1DABYBCADaEO4PuQjDl91EMSClZeGT6ohkf99BuJpLd+Qfpj5rnJEFwxEr
CiykzwQv3vaJR48NrEe2Sa4U6iqGKzI0maDZrWFi8q/4j2hFO4QM8Sa2IWAZKeDa
FKR+csrYX0f1PbiTspr+XjdvYKZtaOOs2qkFo5qOscN2rU7rLtK+NBDUR8sx+wDP
YI0+B0EpkQ0zSB/se918i2APpleqCXL15E3Ie1u+pBdgLiD5ZN1/iE5Tf+lPCcUT
O/stDXzlqz06zVwtSfsX381rz1r+wCsOsTQvpd8d7ztYyMnUwwEOF3b4FukeM+pw
TZHfF8yzjSD5rkYMlL4zP4SpROxyJooNApPbABEBAAG0IUFyYml0cmFyeUtleSA8
YXJiaXRyYXJ5QHRlc3QuY29tPokBTgQTAQgAOBYhBCHupLgewVEnyziIcws53EZa
WzsXBQJdQwAWAhsDBQsJCAcCBhUKCQgLAgQWAgMBAh4BAheAAAoJEAs53EZaWzsX
4RsIALTjYGwJpf5ZDAIx5v84miI7ArP3J3GFYQnBwIwpEmtKh0Y7RcJFNnHlt5Af
Xx4d2HrSxCBD9hiMecLfT73ZQDNLrUeIsv+UJb5JJe/JvQZyD1oWENVHLFzxDC70
IJXEd+5goftcsGoXCz+tbk+NAvx1kRG3xbsx6PrKjU4w1d04aBrvcO8rCwYbd1lS
yOXWYtRJFK4rGlbReKH50onF65g2Sdzaqx6MPT1gPPjyi3LpuPeTypEdGwrK6eok
UqvhLOGiMMXh7s9t6UqfodQ6ayJsimmRw8+tIsSmlMy5cyhcmYQE3+VF2VfTmXv6
ELLG0zVBWAeiDYo8AN7hy/MZM6q5AQ0EXUMAFgEIAMLrdMg/FHHFoZV8hv53Bsvd
Cr64kx1wxMW72rw3k1Onb4pXmoDNCkkTJqNX+9ocxkgf8eMUJKagKLjF/9c8M6oB
SfpL2XfI2WmY3HqhBE8p5WfXY819chH7qhzeDuy36q51CVngDJbl5sS1SIz2xMeC
BW+oqSXc59a6KQ+qXQg0iYUCIvfH+3Yi+wlmWQ1QQjKXGmmvJTR6vTx7pwX6awVd
HIBUw14A2xwC1uqHD+dC0GMNPTNlT1bP/SJ8F/W8uQxadCFyhjEaFWqnAFSpNwT8
mG98DumXbjfhgKPVbSt9uBbMFUXfKpj6uECgvtqVfCpHlQf/tze7gpNsnKgEvdcA
EQEAAYkBNgQYAQgAIBYhBCHupLgewVEnyziIcws53EZaWzsXBQJdQwAWAhsMAAoJ
EAs53EZaWzsXkW4H/RMaBVx3cR+GQMLJ38MxRuIksV6Fi46AGJJeIp9vqNgYQBdq
J3pyPtW0rnBLuqRTZZ+cQOp9mDuaqrblqD9jRm8vKL6vhzRmS1affHD4NPhh1WKH
Avi26TEFE1Y/xQ630mcm5K8CF3ItKqO56MSALzpNdc6tDdDflNd7JhkC6iSVKjaE
BCR8hH8opNGta0cX0isOVLN1z1bRt/xJTOxjXqoJFcmIuHIOCQzk7ODvqyphaeuV
ys/n9RSyDZF01sXnU6LUWsHau5MdFmOZC5oPdWVjB6GIEpZtIccrhm6vH12TVhA3
ERUWZPUImhIpQS8TsuwLMkccr7OEXjUayamdBKw=
=1rdD
-----END PGP PUBLIC KEY BLOCK-----
`)
)
