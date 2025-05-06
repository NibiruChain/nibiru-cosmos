package keeper

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
)

type SchemaProvider interface {
	GetSchema() collections.Schema
}

type HasAddressCodec interface {
	AddressCodec() address.Codec
}
