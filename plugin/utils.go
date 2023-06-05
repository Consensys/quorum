package plugin

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/openpgp" // nolint:staticcheck
)

// Returns true if provided string contains only alphanumeric characters with the exception of the character (.) otherwise false.
func isCleanFileName(s string) bool {
	if s == "" {
		return false
	}
	return regexp.MustCompile(`^[\w.-]+$`).MatchString(s)
}

// Returns true if provided string contains only alphanumeric characters otherwise false
func isCleanEntryPoint(s string) bool {
	if s == "" {
		return false
	}
	return regexp.MustCompile(`^[\w-_.]+$`).MatchString(s)
}

func unzipFile(output string, input *zip.File) error {
	inputFile, err := input.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = inputFile.Close()
	}()
	outputFile, err := os.OpenFile(
		output,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		input.Mode(),
	)
	if err != nil {
		return err
	}
	defer func() {
		_ = outputFile.Close()
	}()
	_, err = io.Copy(outputFile, inputFile)
	return err
}

// Unzip src path to dest. Creates dest if the file doesnt exists.
func unzip(src string, dest string) error {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	targetDir := dest
	for _, file := range zipReader.Reader.File {
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(extractedFilePath, file.Mode()); err != nil {
				return err
			}
		} else {
			if err := unzipFile(extractedFilePath, file); err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns Hex encoded value of sha256(binary content of filepath)
func getSha256Checksum(filePath string) (string, error) {
	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Open a new hash interface to write to
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Unpack pluginPath and returns plugin random generated unpacking path & plugin metadata.
func unpackPlugin(pluginPath string) (string, *MetaData, error) {
	// Unpack pluginMeta
	// Reduce TOC/TOU risk
	unpackDir := path.Join(os.TempDir(), uuid.New(), uuid.New())

	err := os.MkdirAll(unpackDir, os.ModePerm)
	if err != nil {
		return unpackDir, nil, err
	}

	// Unzip to new Dir
	err = unzip(pluginPath, unpackDir)
	if err != nil {
		return unpackDir, nil, err
	}

	// Make Plugin
	pluginMeta := MetaData{}
	// Verify Plugin Structure
	jsonFile, err := os.Open(path.Join(unpackDir, "plugin-meta.json"))
	if err != nil {
		return unpackDir, nil, err
	}
	defer jsonFile.Close()

	if err := json.NewDecoder(jsonFile).Decode(&pluginMeta); err != nil {
		return unpackDir, nil, err
	}

	if pluginMeta.EntryPoint == "" {
		return unpackDir, nil, fmt.Errorf("plugin-meta.json entry point not set")
	}

	if !isCleanEntryPoint(pluginMeta.EntryPoint) {
		return unpackDir, nil, fmt.Errorf("entrypoint must be only alphanumeric value")
	}
	return unpackDir, &pluginMeta, nil
}

func verify(signature, pubkey []byte, checksum string) error {
	// verify file signature
	keyring, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(pubkey))
	if err != nil {
		return err
	}
	entity, err := openpgp.CheckArmoredDetachedSignature(keyring, strings.NewReader(checksum), bytes.NewReader(signature))
	if err != nil {
		log.Debug("unable to verify signature with original checksum. Now add \\n to the end and try", "checksum", checksum, "error", err)
		entity, err = openpgp.CheckArmoredDetachedSignature(keyring, strings.NewReader(checksum+"\n"), bytes.NewReader(signature))
		if err != nil {
			return err
		}
	}
	if entity == nil {
		return fmt.Errorf("verification failed")
	}
	return nil
}

// resolve URL-based value to file path
func resolveFilePath(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return filepath.Abs(filepath.Join(u.Host, u.Path))
}
