package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	crosstypes "github.com/datachainlab/cross/x/core/types"
	"github.com/gogo/protobuf/proto"
)

// TxAuthenticator defines the expected interface of cross-chain authenticator
type TxAuthenticator interface {
	// InitAuthState initializes the state of the tx corresponding to a given txID
	InitAuthState(ctx sdk.Context, txID crosstypes.TxID, signers []Account) error
	// IsCompletedAuth returns a boolean whether the tx corresponding a given txID is completed
	IsCompletedAuth(ctx sdk.Context, txID crosstypes.TxID) (bool, error)
	// Sign executes
	Sign(ctx sdk.Context, txID crosstypes.TxID, signers []Account) (bool, error)
}

// TxManager defines the expected interface of transaction manager
type TxManager interface {
	// IsActive returns a boolean whether the tx corresponding to a given txID is still active
	IsActive(ctx sdk.Context, txID crosstypes.TxID) (bool, error)
	// OnPostAuth represents a callback function is called at post authentication
	OnPostAuth(ctx sdk.Context, txID crosstypes.TxID) error
}

// IsCompleted returns a boolean whether the required authentication is completed
func (s TxAuthState) IsCompleted() bool {
	return len(s.RemainingSigners) == 0
}

// ConsumeSigners removes the signers from required signers
func (s *TxAuthState) ConsumeSigners(signers []Account) (isConsumed bool) {
	before := len(s.RemainingSigners)
	s.RemainingSigners = getRemainingAccounts(signers, s.RemainingSigners)
	return before-len(s.RemainingSigners) > 0
}

func getRemainingAccounts(signers, required []Account) []Account {
	var state = make([]bool, len(required))
	for i, acc := range required {
		for _, s := range signers {
			if acc.Equal(s) {
				state[i] = true
			}
		}
	}
	var remaining []Account
	for i, acc := range required {
		if !state[i] {
			remaining = append(remaining, acc)
		}
	}
	return remaining
}

// AuthExtensionVerifier defines an interface that verifies a tx with an auth extension signature
type AuthExtensionVerifier interface {
	proto.Message
	Verify(ctx sdk.Context, signer Account, signature signing.SignatureV2, tx sdk.Tx) error
}
