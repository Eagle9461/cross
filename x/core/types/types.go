package types

import (
	"bytes"
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

type TxID = []byte

type TxIndex = uint8

type AccountAddress []byte

func (ac AccountAddress) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(ac)
}

func (tx ContractTransaction) GetChainID(m codec.Marshaler) (ChainID, error) {
	var chainID ChainID
	if err := m.UnpackAny(&tx.ChainId, &chainID); err != nil {
		return nil, err
	}
	return chainID, nil
}

func NewReturnValue(v []byte) *ReturnValue {
	rv := ReturnValue{Value: v}
	return &rv
}

func (rv *ReturnValue) IsNil() bool {
	if rv == nil {
		return true
	}
	return false
}

func (rv *ReturnValue) Equal(other *ReturnValue) bool {
	if rv.IsNil() && other.IsNil() {
		return true
	} else if rv.IsNil() && !other.IsNil() {
		return false
	} else if !rv.IsNil() && other.IsNil() {
		return false
	} else {
		return bytes.Equal(rv.Value, other.Value)
	}
}

func (lk Link) ValidateBasic() error {
	if lk.SrcIndex > math.MaxUint8 {
		return fmt.Errorf("src_index value is overflow: %v", lk.SrcIndex)
	}
	return nil
}

func (lk Link) GetSrcIndex() TxIndex {
	return TxIndex(lk.SrcIndex)
}

type ContractCallInfo []byte

type ContractRuntimeInfo struct {
	CommitMode             CommitMode
	StateConstraintType    StateConstraintType
	ExternalObjectResolver ObjectResolver
}

type StateConstraintType = uint32

const (
	NoStateConstraint         StateConstraintType = iota // NoStateConstraint indicates that no constraints on the state before and after the precommit is performed
	ExactMatchStateConstraint                            // ExactMatchStateConstraint indicates the constraint on state state before and after the precommit is performed
	PreStateConstraint                                   // PreStateConstraint indicates the constraint on state before the precommit is performed
	PostStateConstraint                                  // PostStateConstraint indicates the constraint on state after the precommit is performed
)

func NewStateConstraint(tp StateConstraintType, opItems []OP) StateConstraint {
	ops, err := PackOPs(opItems)
	if err != nil {
		panic(err)
	}
	return StateConstraint{Type: tp, Ops: ops}
}

type OP interface {
	proto.Message
}

// ChainID represents an ID of chain that contains a contract function to be called
type ChainID interface {
	Type() string
	Equal(ChainID) bool
	String() string
}

var _ ChainID = (*ChannelInfo)(nil)

// Type implements ChainID.Type
func (ci ChannelInfo) Type() string {
	return "channelinfo"
}

func (ci *ChannelInfo) Equal(other ChainID) bool {
	return ci == other
}

// ChannelResolver defines the interface of resolver resolves chainID to ChannelInfo
type ChannelResolver interface {
	Resolve(ctx sdk.Context, chainID ChainID) (*ChannelInfo, error)
	Capabilities() ChannelResolverCapabilities
}

// ChannelResolverCapabilities defines the capabilities for the ChannelResolver
type ChannelResolverCapabilities interface {
	// CrossChainCalls returns true if support for cross-chain calls is enabled.
	CrossChainCalls() bool
}

type channelResolverCapabilities struct {
	crossChainCalls bool
}

// CrossChainCalls implements ChannelResolverCapabilities.CrossChainCalls
func (c channelResolverCapabilities) CrossChainCalls() bool {
	return c.crossChainCalls
}

// ChannelInfoResolver just returns a given ChannelInfo as is.
type ChannelInfoResolver struct{}

var _ ChannelResolver = (*ChannelInfoResolver)(nil)

// Resolve implements ChannelResolver.ResResolve
func (r ChannelInfoResolver) Resolve(ctx sdk.Context, chainID ChainID) (*ChannelInfo, error) {
	ci, ok := chainID.(*ChannelInfo)
	if !ok {
		return nil, fmt.Errorf("cannot resolve '%v'", chainID)
	}
	return ci, nil
}

// Capabilities implements ChannelResolver.Capabilities
func (r ChannelInfoResolver) Capabilities() ChannelResolverCapabilities {
	return channelResolverCapabilities{crossChainCalls: false}
}

// TODO set txInfo to linkObjects
func NewContractTransactionInfo(tx ContractTransaction, linkObjects []Object) ContractTransactionInfo {
	return ContractTransactionInfo{
		Tx: tx,
	}
}

func (ti ContractTransactionInfo) ValidateBasic() error {
	return nil
}
