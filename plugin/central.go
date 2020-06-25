package plugin

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

// Central https centralClient communicating with Plugin Central
type CentralClient struct {
	config     *PluginCentralConfiguration
	httpClient *http.Client
}

// Create New Central Client
func NewPluginCentralClient(config *PluginCentralConfiguration) *CentralClient {
	c := &CentralClient{
		config: config,
	}
	c.httpClient = &http.Client{}
	c.httpClient.Transport = &http.Transport{
		DialTLS: c.getNewSecureDialer(),
	}
	return c
}

// Builds a Dialer that supports CA Verification & Certificate Pinning.
func (cc *CentralClient) getNewSecureDialer() Dialer {
	return func(network, addr string) (net.Conn, error) {
		c, err := tls.Dial(network, addr, &tls.Config{InsecureSkipVerify: cc.config.InsecureSkipTLSVerify})
		if err != nil {
			return c, err
		}
		// support certificate pinning?
		if cc.config.CertFingerprint != "" {
			conState := c.ConnectionState()
			for _, peercert := range conState.PeerCertificates {
				if bytes.Equal(peercert.Signature[0:], []byte(cc.config.CertFingerprint)) {
					return c, nil
				}
			}
			return nil, fmt.Errorf("certificate pinning failed")
		}
		return c, nil
	}
}

// Get the public key from central
func (cc *CentralClient) PublicKey() ([]byte, error) {
	target := fmt.Sprintf("%s/%s", cc.config.BaseURL, cc.config.PublicKeyURI)
	log.Debug("downloading public key", "url", target)
	readCloser, err := cc.get(target)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = readCloser.Close()
	}()
	return ioutil.ReadAll(readCloser)
}

// retrieve plugin signature
func (cc *CentralClient) PluginSignature(definition *PluginDefinition) ([]byte, error) {
	target := fmt.Sprintf("%s/%s/%s", cc.config.BaseURL, definition.RemotePath(), definition.SignatureFileName())
	log.Debug("downloading plugin signature file", "url", target)
	readCloser, err := cc.get(target)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = readCloser.Close()
	}()
	return ioutil.ReadAll(readCloser)
}

// retrieve plugin distribution file
func (cc *CentralClient) PluginDistribution(definition *PluginDefinition, outFilePath string) error {
	target := fmt.Sprintf("%s/%s/%s", cc.config.BaseURL, definition.RemotePath(), definition.DistFileName())
	log.Debug("downloading plugin zip file", "url", target)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = outFile.Close()
	}()
	readCloser, err := cc.get(target)
	if err != nil {
		return err
	}
	defer func() {
		_ = readCloser.Close()
	}()
	_, err = io.Copy(outFile, readCloser)
	return err
}

// perform HTTP GET
//
// caller needs to close the reader
func (cc *CentralClient) get(target string) (io.ReadCloser, error) {
	if err := isValidTargetURL(cc.config.BaseURL, target); err != nil {
		return nil, err
	}
	res, err := cc.httpClient.Get(target)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer func() {
			_ = res.Body.Close()
		}()
		data, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("HTTP GET error: code=%d, status=%s, body=%s", res.StatusCode, res.Status, string(data))
	}
	return res.Body, nil
}

// An adapter function for tls.Dial with CA verification & SSL Pinning support.
type Dialer func(network, addr string) (net.Conn, error)

// Validate the target url is well formed and match base.
func isValidTargetURL(base string, target string) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	t, err := url.Parse(base)
	if err != nil {
		return err
	}
	if strings.Compare(t.Host, u.Host) != 0 || strings.Compare(t.Scheme, u.Scheme) != 0 {
		return fmt.Errorf("target host doesnt match base host")
	}
	return nil
}
