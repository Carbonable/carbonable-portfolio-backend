package resolver

import (
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
	Rpc    *rpc.Provider
}
