package tessera

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
)

const apiVersion1 = "1.0"

type versions struct {
	Versions []version `json:"versions"`
}

type version struct {
	Version string `json:"version"`
}

// this method will be removed once quorum will implement a versioned tessera client (in line with tessera API versioning)
func RetrieveTesseraAPIVersion(client *engine.Client) string {
	res, err := client.Get("/version/api")
	if err != nil {
		log.Error("Error invoking the tessera /version/api API: %v.", err)
		return apiVersion1
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Error("Invalid status code returned by the tessera /version/api API: %d.", res.StatusCode)
		return apiVersion1
	}
	versionsObj := new(versions)
	if err := json.NewDecoder(res.Body).Decode(versionsObj); err != nil {
		log.Error("Unable to deserialize the tessera response for /version/api API: %v.", err)
		return apiVersion1
	}
	if len(versionsObj.Versions) == 0 {
		log.Error("Expecting at least one API version to be returned by the tessera /version/api API.")
		return apiVersion1
	}
	latestVersion := versionsObj.Versions[len(versionsObj.Versions)-1]
	if len(latestVersion.Version) == 0 {
		log.Error("Invalid version object returned by the tessera /version/api API.")
		return apiVersion1
	}
	log.Info("Tessera API version: %s", latestVersion.Version)
	return latestVersion.Version
}
