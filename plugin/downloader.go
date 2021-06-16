package plugin

import (
	"fmt"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// get plugin zip file from local or remote
type Downloader struct {
	pm *PluginManager
}

func NewDownloader(pm *PluginManager) *Downloader {
	return &Downloader{
		pm: pm,
	}
}

func (d *Downloader) Download(definition *PluginDefinition) (string, error) {
	// check if plugin is already in the local
	pluginFile := path.Join(d.pm.pluginBaseDir, definition.DistFileName())
	exist := common.FileExist(pluginFile)
	log.Debug("checking plugin zip file", "path", pluginFile, "exist", exist)
	if exist {
		return pluginFile, nil
	}
	if err := d.pm.centralClient.PluginDistribution(definition, pluginFile); err != nil {
		return "", fmt.Errorf("can't download from Plugin Central due to: %s. Please download the plugin manually and copy it to %s", err, d.pm.pluginBaseDir)
	}
	return pluginFile, nil
}
