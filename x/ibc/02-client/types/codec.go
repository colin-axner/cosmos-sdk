package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
)

// RegisterCodec registers the IBC client interfaces and types
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.ClientState)(nil), nil)
	cdc.RegisterInterface((*exported.ConsensusState)(nil), nil)
	cdc.RegisterInterface((*exported.Header)(nil), nil)
	cdc.RegisterInterface((*exported.Misbehaviour)(nil), nil)
	cdc.RegisterInterface((*exported.MsgCreateClient)(nil), nil)
	cdc.RegisterInterface((*exported.MsgUpdateClient)(nil), nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"cosmos.ibc.client.ClientState",
		(*exported.ClientState)(nil),
	)
	registry.RegisterInterface(
		"cosmos.ibc.client.ConsensusState",
		(*exported.ConsensusState)(nil),
	)
	registry.RegisterInterface(
		"cosmos.ibc.client.Header",
		(*exported.Header)(nil),
	)
	registry.RegisterInterface(
		"cosmos.ibc.client.Misbehaviour",
		(*exported.Misbehaviour)(nil),
	)
	registry.RegisterInterface(
		"cosmos.ibc.client.MsgCreateClient",
		(*exported.MsgCreateClient)(nil),
	)
	registry.RegisterInterface(
		"cosmos.ibc.client.MsgUpdateClient",
		(*exported.MsgUpdateClient)(nil),
	)
}

var (
	amino = codec.New()

	// SubModuleCdc references the global x/ibc/02-client module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/ibc/02-client and
	// defined at the application level.
	SubModuleCdc = codec.NewHybridCodec(amino, cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
