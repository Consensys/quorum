package qlight

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/plugin/security"
)

type TLSConfig struct {
	CACertFileName       string
	ClientCACertFileName string
	ClientAuth           int
	CertFileName         string
	KeyFileName          string
	InsecureSkipVerify   bool
	CipherSuites         string
	ServerName           string
}

func NewTLSConfig(config *TLSConfig) (*tls.Config, error) {
	if config.InsecureSkipVerify {
		return &tls.Config{
			InsecureSkipVerify: true,
		}, nil
	}
	var (
		CA_Pool *x509.CertPool
		err     error
	)
	if len(config.CACertFileName) > 0 {
		CA_Pool, err = x509.SystemCertPool()
		if err != nil {
			CA_Pool = x509.NewCertPool()
		}
		cert, err := os.ReadFile(config.CACertFileName)
		if err != nil {
			return nil, err
		}
		CA_Pool.AppendCertsFromPEM(cert)
	}

	var (
		ClientCA_Pool *x509.CertPool
		ClientAuth    tls.ClientAuthType
	)
	if len(config.ClientCACertFileName) > 0 {
		ClientCA_Pool, err = x509.SystemCertPool()
		if err != nil {
			ClientCA_Pool = x509.NewCertPool()
		}
		cert, err := os.ReadFile(config.ClientCACertFileName)
		if err != nil {
			return nil, err
		}
		ClientCA_Pool.AppendCertsFromPEM(cert)
		if config.ClientAuth < 0 || config.ClientAuth > 4 {
			return nil, fmt.Errorf("Invalid ClientAuth value: %d", config.ClientAuth)
		}
		ClientAuth = tls.ClientAuthType(config.ClientAuth)
	}

	var certificates []tls.Certificate

	if len(config.CertFileName) > 0 && len(config.KeyFileName) > 0 {
		cert, err := tls.LoadX509KeyPair(config.CertFileName, config.KeyFileName)
		if err != nil {
			return nil, err
		}
		certificates = []tls.Certificate{cert}
	}

	var CipherSuites []uint16
	if len(config.CipherSuites) > 0 {
		cipherSuitesStrings := strings.FieldsFunc(config.CipherSuites, func(r rune) bool {
			return r == ','
		})
		if len(cipherSuitesStrings) > 0 {
			cipherSuiteList := make(security.CipherSuiteList, len(cipherSuitesStrings))
			for i, s := range cipherSuitesStrings {
				cipherSuiteList[i] = security.CipherSuite(strings.TrimSpace(s))
			}
			CipherSuites, err = cipherSuiteList.ToUint16Array()
			if err != nil {
				return nil, err
			}
		}
	}

	return &tls.Config{
		RootCAs:      CA_Pool,
		Certificates: certificates,
		ServerName:   config.ServerName,
		ClientCAs:    ClientCA_Pool,
		ClientAuth:   ClientAuth,
		CipherSuites: CipherSuites,
	}, nil
}
