package keeper

import (
	"cosmossdk.io/collections"
)

type SchemaProvider interface {
	GetSchema() collections.Schema
}
