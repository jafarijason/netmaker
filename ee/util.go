package ee

import (
	"encoding/base64"

	"github.com/gravitl/netmaker/logic"
)

var isEnterprise bool

// setIsEnterprise - sets server to use enterprise features
func setIsEnterprise() {
	isEnterprise = true
	logic.SetEEForTelemetry(isEnterprise)
}

// base64encode - base64 encode helper function
func base64encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// base64decode - base64 decode helper function
func base64decode(input string) []byte {

	bytes, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		return nil
	}

	return bytes
}
func getCurrentServerLimit() (limits LicenseLimits) {
	limits.SetDefaults()
	hosts, err := logic.GetAllHosts()
	if err == nil {
		limits.Hosts = len(hosts)
	}
	clients, err := logic.GetAllExtClients()
	if err == nil {
		limits.Clients = len(clients)
	}
	users, err := logic.GetUsers()
	if err == nil {
		limits.Users = len(users)
	}
	networks, err := logic.GetNetworks()
	if err == nil {
		limits.Networks = len(networks)
	}
	return
}
