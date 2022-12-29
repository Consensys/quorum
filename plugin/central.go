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
	"runtime"
	"strings"
	"text/template"
	"time"

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

// Get the public key from central. PublicKeyURI can be relative to the base URL
// so we need to parse and make sure finally URL is resolved.
func (cc *CentralClient) PublicKey() ([]byte, error) {
	target, err := cc.toURL(cc.config.PublicKeyURI)
	if err != nil {
		return nil, err
	}
	log.Debug("downloading public key", "url", target)
	buf := new(bytes.Buffer)
	if err := cc.download(target, buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// retrieve plugin signature
func (cc *CentralClient) PluginSignature(definition *PluginDefinition) ([]byte, error) {
	target, err := cc.toURLFromTemplate(cc.config.PluginSigPathTemplate, definition)
	if err != nil {
		return nil, err
	}
	log.Debug("downloading plugin signature file", "url", target)
	buf := new(bytes.Buffer)
	if err := cc.download(target, buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// retrieve plugin distribution file
func (cc *CentralClient) PluginDistribution(definition *PluginDefinition, outFilePath string) error {
	target, err := cc.toURLFromTemplate(cc.config.PluginDistPathTemplate, definition)
	if err != nil {
		return err
	}
	log.Debug("downloading plugin zip file", "url", target)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = outFile.Close()
	}()
	return cc.download(target, outFile)
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

// return full URL using config.BaseURL
func (cc *CentralClient) toURL(relativePath string) (string, error) {
	base, err := url.Parse(cc.config.BaseURL)
	if err != nil {
		return "", err
	}
	u, err := base.Parse(relativePath)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// return full URL using config.BaseURL from given template
func (cc *CentralClient) toURLFromTemplate(pathTemplate string, definition *PluginDefinition) (string, error) {
	t, err := template.New("").Parse(pathTemplate)
	if err != nil {
		return "", err
	}
	path := new(bytes.Buffer)
	if err := t.Execute(path, struct {
		Name    string
		Version string
		OS      string
		Arch    string
	}{
		Name:    definition.Name,
		Version: string(definition.Version),
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
	}); err != nil {
		return "", err
	}
	return cc.toURL(path.String())
}

// perform http GET to the target URL and write output to out
func (cc *CentralClient) download(target string, out io.Writer) (err error) {
	defer func(start time.Time) {
		log.Debug("download completed", "url", target, "err", err, "took", time.Since(start))
	}(time.Now())
	var readCloser io.ReadCloser
	readCloser, err = cc.get(target)
	if err != nil {
		return
	}
	defer func() {
		_ = readCloser.Close()
	}()
	_, err = io.Copy(out, readCloser)
	return err
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
