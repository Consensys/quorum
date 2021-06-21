package http

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

// Load Root CA certificate(s).
// Path can be a single certificate file, or a comma separated list containing a combination of
// certificate files or directories containing certificate files.
func loadRootCaCerts(rootCAPath string) (*x509.CertPool, error) {
	rootCAPool, err := x509.SystemCertPool()
	if err != nil {
		rootCAPool = x509.NewCertPool()
	}
	if len(rootCAPath) == 0 {
		return rootCAPool, nil
	}

	list := strings.Split(rootCAPath, ",")
	for _, thisFileOrDirEntry := range list {
		info, err := os.Lstat(thisFileOrDirEntry)
		if err != nil {
			return nil, fmt.Errorf("unable to check whether RootCA entry '%v' is a file or directory, due to: %s", thisFileOrDirEntry, err)
		}

		if info.Mode()&os.ModeDir != 0 {
			fileList, err := ioutil.ReadDir(thisFileOrDirEntry)
			if err != nil {
				return nil, fmt.Errorf("unable to read contents of RootCA directory '%v', due to: %s", thisFileOrDirEntry, err)
			}

			for _, fileinfo := range fileList {
				if err := loadRootCAFromFile(thisFileOrDirEntry+"/"+fileinfo.Name(), rootCAPool); err != nil {
					return nil, err
				}
			}
		} else if err := loadRootCAFromFile(thisFileOrDirEntry, rootCAPool); err != nil {
			return nil, err
		}
	}

	return rootCAPool, nil
}

func loadRootCAFromFile(file string, roots *x509.CertPool) error {
	log.Debug("loading RootCA certificate for connection to private transaction manager", "file", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("unable to read contents of RootCA certificate file '%v', due to: %s", file, err)
	}
	if !roots.AppendCertsFromPEM(data) {
		return fmt.Errorf("failed to add TlsRootCA certificate to pool, check that '%v' contains a valid RootCA certificate", file)
	}
	return nil
}
