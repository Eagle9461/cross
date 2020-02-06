package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(StateTransition{}, "crossccc/StateTransition", nil)
	cdc.RegisterConcrete(StateTransitions{}, "crossccc/StateTransitions", nil)
	cdc.RegisterConcrete(ChannelInfo{}, "crossccc/ChannelInfo", nil)
	cdc.RegisterInterface((*OP)(nil), nil)
}
