// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/rpc"
	"gopkg.in/urfave/cli.v1"
)

var (
	consoleFlags   = []cli.Flag{utils.JSpathFlag, utils.ExecFlag, utils.PreloadJSFlag}
	rpcClientFlags = []cli.Flag{utils.RPCClientToken, utils.RPCClientTLSCert, utils.RPCClientTLSCaCert, utils.RPCClientTLSCipherSuites, utils.RPCClientTLSInsecureSkipVerify}

	consoleCommand = cli.Command{
		Action:   utils.MigrateFlags(localConsole),
		Name:     "console",
		Usage:    "Start an interactive JavaScript environment",
		Flags:    append(append(append(nodeFlags, rpcFlags...), consoleFlags...), whisperFlags...),
		Category: "CONSOLE COMMANDS",
		Description: `
The Geth console is an interactive shell for the JavaScript runtime environment
which exposes a node admin interface as well as the Ðapp JavaScript API.
See https://github.com/ethereum/go-ethereum/wiki/JavaScript-Console.`,
	}

	attachCommand = cli.Command{
		Action:    utils.MigrateFlags(remoteConsole),
		Name:      "attach",
		Usage:     "Start an interactive JavaScript environment (connect to node)",
		ArgsUsage: "[endpoint]",
		Flags:     append(append(consoleFlags, utils.DataDirFlag), rpcClientFlags...),
		Category:  "CONSOLE COMMANDS",
		Description: `
The Geth console is an interactive shell for the JavaScript runtime environment
which exposes a node admin interface as well as the Ðapp JavaScript API.
See https://github.com/ethereum/go-ethereum/wiki/JavaScript-Console.
This command allows to open a console on a running geth node.`,
	}

	javascriptCommand = cli.Command{
		Action:    utils.MigrateFlags(ephemeralConsole),
		Name:      "js",
		Usage:     "Execute the specified JavaScript files",
		ArgsUsage: "<jsfile> [jsfile...]",
		Flags:     append(nodeFlags, consoleFlags...),
		Category:  "CONSOLE COMMANDS",
		Description: `
The JavaScript VM exposes a node admin interface as well as the Ðapp
JavaScript API. See https://github.com/ethereum/go-ethereum/wiki/JavaScript-Console`,
	}
)

// Quorum
//
// read tls client configuration from command line arguments
//
// only for HTTPS or WSS
func readTLSClientConfig(endpoint string, ctx *cli.Context) (*tls.Config, bool, error) {
	if !strings.HasPrefix(endpoint, "https://") && !strings.HasPrefix(endpoint, "wss://") {
		return nil, false, nil
	}
	hasCustomTls := false
	insecureSkipVerify := ctx.Bool(utils.RPCClientTLSInsecureSkipVerify.Name)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	var certFile, caFile string
	if !insecureSkipVerify {
		var certPem, caPem []byte
		certFile, caFile = ctx.String(utils.RPCClientTLSCert.Name), ctx.String(utils.RPCClientTLSCaCert.Name)
		var err error
		if certFile != "" {
			if certPem, err = ioutil.ReadFile(certFile); err != nil {
				return nil, true, err
			}
		}
		if caFile != "" {
			if caPem, err = ioutil.ReadFile(caFile); err != nil {
				return nil, true, err
			}
		}
		if len(certPem) != 0 || len(caPem) != 0 {
			certPool, err := x509.SystemCertPool()
			if err != nil {
				certPool = x509.NewCertPool()
			}
			if len(certPem) != 0 {
				certPool.AppendCertsFromPEM(certPem)
			}
			if len(caPem) != 0 {
				certPool.AppendCertsFromPEM(caPem)
			}
			tlsConfig.RootCAs = certPool
			hasCustomTls = true
		}
	} else {
		hasCustomTls = true
	}
	cipherSuitesInput := ctx.String(utils.RPCClientTLSCipherSuites.Name)
	cipherSuitesStrings := strings.FieldsFunc(cipherSuitesInput, func(r rune) bool {
		return r == ','
	})
	if len(cipherSuitesStrings) > 0 {
		cipherSuiteList := make(security.CipherSuiteList, len(cipherSuitesStrings))
		for i, s := range cipherSuitesStrings {
			cipherSuiteList[i] = security.CipherSuite(strings.TrimSpace(s))
		}
		cipherSuites, err := cipherSuiteList.ToUint16Array()
		if err != nil {
			return nil, true, err
		}
		tlsConfig.CipherSuites = cipherSuites
		hasCustomTls = true
	}
	if !hasCustomTls {
		return nil, false, nil
	}
	return tlsConfig, hasCustomTls, nil
}

// localConsole starts a new geth node, attaching a JavaScript console to it at the
// same time.
func localConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	prepare(ctx)
	node := makeFullNode(ctx)
	startNode(ctx, node)
	defer node.Close()

	// Attach to the newly started node and start the JavaScript console
	client, err := node.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to the inproc geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	console, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer console.Stop(false)

	// If only a short execution was requested, evaluate and return
	if script := ctx.GlobalString(utils.ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}
	// Otherwise print the welcome screen and enter interactive mode
	console.Welcome()
	console.Interactive()

	return nil
}

// remoteConsole will connect to a remote geth instance, attaching a JavaScript
// console to it.
func remoteConsole(ctx *cli.Context) error {
	// Attach to a remotely running geth instance and start the JavaScript console
	endpoint := ctx.Args().First()
	if endpoint == "" {
		path := node.DefaultDataDir()
		if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
			path = ctx.GlobalString(utils.DataDirFlag.Name)
		}
		if path != "" {
			if ctx.GlobalBool(utils.TestnetFlag.Name) {
				path = filepath.Join(path, "testnet")
			} else if ctx.GlobalBool(utils.RinkebyFlag.Name) {
				path = filepath.Join(path, "rinkeby")
			}
		}
		endpoint = fmt.Sprintf("%s/geth.ipc", path)
	}
	client, err := dialRPC(endpoint, ctx)
	if err != nil {
		utils.Fatalf("Unable to attach to remote geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	consl, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer consl.Stop(false)

	if script := ctx.GlobalString(utils.ExecFlag.Name); script != "" {
		consl.Evaluate(script)
		return nil
	}

	// Otherwise print the welcome screen and enter interactive mode
	consl.Welcome()
	consl.Interactive()

	return nil
}

// dialRPC returns a RPC client which connects to the given endpoint.
// The check for empty endpoint implements the defaulting logic
// for "geth attach" and "geth monitor" with no argument.
//
// Quorum: passing the cli context to build security-aware client:
// 1. Custom TLS configuration
// 2. Access Token awareness via rpc.HttpCredentialsProviderFunc
func dialRPC(endpoint string, ctx *cli.Context) (*rpc.Client, error) {
	if endpoint == "" {
		endpoint = node.DefaultIPCEndpoint(clientIdentifier)
	} else if strings.HasPrefix(endpoint, "rpc:") || strings.HasPrefix(endpoint, "ipc:") {
		// Backwards compatibility with geth < 1.5 which required
		// these prefixes.
		endpoint = endpoint[4:]
	}
	var (
		client  *rpc.Client
		err     error
		dialCtx = context.Background()
	)
	tlsConfig, hasCustomTls, tlsReadErr := readTLSClientConfig(endpoint, ctx)
	if tlsReadErr != nil {
		return nil, tlsReadErr
	}
	if token := ctx.String(utils.RPCClientToken.Name); token != "" {
		var f rpc.HttpCredentialsProviderFunc = func(ctx context.Context) (string, error) {
			return token, nil
		}
		// it's important that f MUST BE OF TYPE rpc.HttpCredentialsProviderFunc
		dialCtx = context.WithValue(dialCtx, rpc.CtxCredentialsProvider, f)
	}
	if hasCustomTls {
		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
		switch u.Scheme {
		case "https":
			customHttpClient := &http.Client{
				Transport: http.DefaultTransport,
			}
			customHttpClient.Transport.(*http.Transport).TLSClientConfig = tlsConfig
			client, err = rpc.DialHTTPWithClient(endpoint, customHttpClient)
		case "wss":
			client, err = rpc.DialWebsocketWithCustomTLS(dialCtx, endpoint, "", tlsConfig)
		default:
			log.Warn("unsupported scheme for custom TLS which is only for HTTPS/WSS", "scheme", u.Scheme)
			client, err = rpc.DialContext(dialCtx, endpoint)
		}
	} else {
		client, err = rpc.DialContext(dialCtx, endpoint)
	}
	if f, ok := dialCtx.Value(rpc.CtxCredentialsProvider).(rpc.HttpCredentialsProviderFunc); ok && err == nil {
		client, err = client.WithHTTPCredentials(f)
	}
	return client, err
}

// ephemeralConsole starts a new geth node, attaches an ephemeral JavaScript
// console to it, executes each of the files specified as arguments and tears
// everything down.
func ephemeralConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	node := makeFullNode(ctx)
	startNode(ctx, node)
	defer node.Close()

	// Attach to the newly started node and start the JavaScript console
	client, err := node.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to the inproc geth: %v", err)
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	console, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer console.Stop(false)

	// Evaluate each of the specified JavaScript files
	for _, file := range ctx.Args() {
		if err = console.Execute(file); err != nil {
			utils.Fatalf("Failed to execute %s: %v", file, err)
		}
	}
	// Wait for pending callbacks, but stop for Ctrl-C.
	abort := make(chan os.Signal, 1)
	signal.Notify(abort, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-abort
		os.Exit(0)
	}()
	console.Stop(true)

	return nil
}
