package cli

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmtypes "github.com/tendermint/tendermint/types"

	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/cross/x/core/initiator/types"
	txtypes "github.com/datachainlab/cross/x/core/tx/types"
	xcctypes "github.com/datachainlab/cross/x/core/xcc/types"
)

// NewInitiateTxCmd returns the command to create a NewMsgInitiateTx transaction
func NewInitiateTxCmd() *cobra.Command {
	const (
		flagContractTransactions = "contract-txs"
	)

	cmd := &cobra.Command{
		Use:   "initiate-tx",
		Short: "Create a NewMsgInitiateTx transaction",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			clientCtx = clientCtx.WithOutputFormat("json")
			sender := authtypes.AccountIDFromAccAddress(clientCtx.GetFromAddress())
			ctxs, err := readContractTransactions(clientCtx.JSONCodec, viper.GetStringSlice(flagContractTransactions))
			if err != nil {
				return err
			}

			h, height, err := QueryTendermintHeader(clientCtx)
			if err != nil {
				return err
			}
			version := clienttypes.ParseChainID(h.Header.ChainID)

			msg := types.NewMsgInitiateTx(
				[]authtypes.Account{authtypes.NewAccount(sender, authtypes.NewAuthTypeLocal())},
				clientCtx.ChainID,
				uint64(time.Now().Unix()),
				txtypes.COMMIT_PROTOCOL_SIMPLE,
				ctxs,
				clienttypes.NewHeight(version, uint64(height)+100),
				0,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().StringSlice(flagContractTransactions, nil, "A file path to includes a contract transaction")
	cmd.MarkFlagRequired(flagContractTransactions)

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func readContractTransactions(m codec.JSONCodec, pathList []string) ([]types.ContractTransaction, error) {
	var cTxs []types.ContractTransaction
	for _, path := range pathList {
		bz, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var cTx types.ContractTransaction
		if err := m.UnmarshalJSON(bz, &cTx); err != nil {
			return nil, err
		}
		cTxs = append(cTxs, cTx)
	}
	return cTxs, nil
}

func resolveXCC(queryClient channeltypes.QueryClient, s string) (*codectypes.Any, error) {
	ci, err := parseChannelInfoFromString(s)
	if err != nil {
		return nil, err
	}
	return xcctypes.PackCrossChainChannel(ci)
}

// QueryTendermintHeader takes a client context and returns the appropriate
// tendermint header
// Original implementation(but has a little) is here: https://github.com/cosmos/cosmos-sdk/blob/300b7393addba8c162cae929db90b083dcf93bd0/x/ibc/core/02-client/client/utils/utils.go#L123
func QueryTendermintHeader(clientCtx client.Context) (ibctmtypes.Header, int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return ibctmtypes.Header{}, 0, err
	}

	info, err := node.ABCIInfo(context.Background())
	if err != nil {
		return ibctmtypes.Header{}, 0, err
	}

	height := info.Response.LastBlockHeight

	commit, err := node.Commit(context.Background(), &height)
	if err != nil {
		return ibctmtypes.Header{}, 0, err
	}

	page := 1
	count := 10_000

	validators, err := node.Validators(context.Background(), &height, &page, &count)
	if err != nil {
		return ibctmtypes.Header{}, 0, err
	}

	protoCommit := commit.SignedHeader.ToProto()
	protoValset, err := tmtypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		return ibctmtypes.Header{}, 0, err
	}

	header := ibctmtypes.Header{
		SignedHeader: protoCommit,
		ValidatorSet: protoValset,
	}

	return header, height, nil
}
