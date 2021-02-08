package tessera

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
)

const apiVersion1 = "1.0"

// this method will be removed once quorum will implement a versioned tessera client (in line with tessera API versioning)
func RetrieveTesseraAPIVersion(client *engine.Client) string {
	res, err := client.Get("/version/api")
	if err != nil {
		log.Error("Error invoking the tessera /version/api API:", "err", err)
		return apiVersion1
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Error(fmt.Sprintf("Invalid status code returned by the tessera /version/api API: %d.", res.StatusCode))
		return apiVersion1
	}
	var versions []string
	if err := json.NewDecoder(res.Body).Decode(&versions); err != nil {
		log.Error("Unable to deserialize the tessera response for /version/api API:", "err", err)
		return apiVersion1
	}
	if len(versions) == 0 {
		log.Error("Expecting at least one API version to be returned by the tessera /version/api API.")
		return apiVersion1
	}
	// pick the latest version from the versions array
	latestVersion := apiVersion1
	latestParsedVersion, _ := parseVersion([]byte(latestVersion))
	for _, ver := range versions {
		if len(ver) == 0 {
			log.Error("Invalid (empty) version returned by the tessera /version/api API. Skipping value.")
			continue
		}
		parsedVer, err := parseVersion([]byte(ver))
		if err != nil {
			log.Error(fmt.Sprintf("Unable to parse version returned by the tessera /version/api API: %s. Skipping value.", ver))
			continue
		}
		if compareVersions(parsedVer, latestParsedVersion) > 0 {
			latestVersion = ver
			latestParsedVersion = parsedVer
		}
	}
	log.Info(fmt.Sprintf("Tessera API version: %s", latestVersion))
	return latestVersion
}
