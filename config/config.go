package config

import (
	_ "embed"
	"encoding/json"
	"os"

	"github.com/carbonable-labs/indexer.sdk/sdk"
)

//go:embed mainnet.json
var Mainnet []byte

//go:embed sepolia.json
var Sepolia []byte

func GetContracts() sdk.Config {
	var config sdk.Config
	var data []byte
	if os.Getenv("NETWORK") == "mainnet" {
		data = Mainnet
	} else {
		data = Sepolia
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}
