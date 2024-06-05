package resolver

import "github.com/carbonable/carbonable-portfolio-backend/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
}
