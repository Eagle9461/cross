package keeper_test

import (
	"fmt"

	"github.com/bluele/crossccc/x/ibc/contract"
	"github.com/bluele/crossccc/x/ibc/crossccc"
	"github.com/bluele/crossccc/x/ibc/store/lock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	connection "github.com/cosmos/cosmos-sdk/x/ibc/03-connection"
	connectionexported "github.com/cosmos/cosmos-sdk/x/ibc/03-connection/exported"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	channelexported "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	tendermint "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint"
	commitment "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// define constants used for testing
const (
	testClientType     = clientexported.Tendermint
	testChannelOrder   = channelexported.UNORDERED
	testChannelVersion = "1.0"
)

func (suite *KeeperTestSuite) createClient(actx *appContext, clientID string) {
	actx.app.Commit()
	commitID := actx.app.LastCommitID()

	actx.app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: actx.app.LastBlockHeight() + 1}})
	actx.ctx = actx.app.BaseApp.NewContext(false, abci.Header{})

	consensusState := tendermint.ConsensusState{
		Root:             commitment.NewRoot(commitID.Hash),
		ValidatorSetHash: actx.valSet.Hash(),
	}

	_, err := actx.app.IBCKeeper.ClientKeeper.CreateClient(actx.ctx, clientID, testClientType, consensusState)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) updateClient(actx *appContext, clientID string) {
	// always commit and begin a new block on updateClient
	actx.app.Commit()
	commitID := actx.app.LastCommitID()

	actx.app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: actx.app.LastBlockHeight() + 1}})
	actx.ctx = actx.app.BaseApp.NewContext(false, abci.Header{})

	state := tendermint.ConsensusState{
		Root: commitment.NewRoot(commitID.Hash),
	}

	actx.app.IBCKeeper.ClientKeeper.SetClientConsensusState(actx.ctx, clientID, 1, state)
}

func (suite *KeeperTestSuite) createConnection(actx *appContext, clientID, connectionID, counterpartyClientID, counterpartyConnectionID string, state connectionexported.State) {
	connection := connection.ConnectionEnd{
		State:    state,
		ClientID: clientID,
		Counterparty: connection.Counterparty{
			ClientID:     counterpartyClientID,
			ConnectionID: counterpartyConnectionID,
			Prefix:       actx.app.IBCKeeper.ConnectionKeeper.GetCommitmentPrefix(),
		},
		Versions: connection.GetCompatibleVersions(),
	}

	actx.app.IBCKeeper.ConnectionKeeper.SetConnection(actx.ctx, connectionID, connection)
}

func (suite *KeeperTestSuite) createChannel(actx *appContext, portID string, chanID string, connID string, counterpartyPort string, counterpartyChan string, state channelexported.State) {
	ch := channel.Channel{
		State:    state,
		Ordering: testChannelOrder,
		Counterparty: channel.Counterparty{
			PortID:    counterpartyPort,
			ChannelID: counterpartyChan,
		},
		ConnectionHops: []string{connID},
		Version:        testChannelVersion,
	}

	actx.app.IBCKeeper.ChannelKeeper.SetChannel(actx.ctx, portID, chanID, ch)
}

func (suite *KeeperTestSuite) queryProof(actx *appContext, key []byte) (proof commitment.Proof, height int64) {
	res := actx.app.Query(abci.RequestQuery{
		Path:  fmt.Sprintf("store/%s/key", ibctypes.StoreKey),
		Data:  key,
		Prove: true,
	})

	height = res.Height
	proof = commitment.Proof{
		Proof: res.Proof,
	}

	return
}

func (suite *KeeperTestSuite) TestSendInitiate() {
	lock.RegisterCodec(crossccc.ModuleCdc)

	initiator := sdk.AccAddress("initiator")

	signer1 := sdk.AccAddress("signer1")
	ci1 := contract.NewContractInfo("c1", "issue", [][]byte{[]byte("100")})

	signer2 := sdk.AccAddress("signer2")
	ci2 := contract.NewContractInfo("c2", "issue", [][]byte{[]byte("100")})

	app0 := suite.createApp("app0") // coordinator node
	app1 := suite.createApp("app1")
	app2 := suite.createApp("app2")

	ch0to1 := crossccc.NewChannelInfo("testportzeroone", "testchannelzeroone") // app0 -> app1
	ch1to0 := crossccc.NewChannelInfo("testportonezero", "testchannelonezero") // app1 -> app0
	ch0to2 := crossccc.NewChannelInfo("testportonetwo", "testchannelonetwo")   // app0 -> app2
	ch2to0 := crossccc.NewChannelInfo("testporttwozero", "testchanneltwozero") // app2 -> app0

	var err error
	var nonce uint64 = 1
	var tss = []crossccc.StateTransition{
		crossccc.NewStateTransition(
			ch0to1,
			signer1,
			ci1.Bytes(),
			[]crossccc.OP{lock.Read{}, lock.Write{}},
		),
		crossccc.NewStateTransition(
			ch0to2,
			signer2,
			ci2.Bytes(),
			[]crossccc.OP{lock.Read{}, lock.Write{}},
		),
	}

	msg := crossccc.NewMsgInitiate(
		initiator,
		tss,
		nonce,
	)

	err = app0.app.CrosscccKeeper.MulticastInitiatePacket(
		app0.ctx,
		initiator,
		msg,
		msg.StateTransitions,
	)
	suite.Error(err) // channel does not exist

	// Try to open a channel and connection between app0 and app1, app2

	suite.openChannels(
		app1.chainID,
		app0.chainID+app1.chainID,
		ch0to1,
		app0,

		app0.chainID,
		app1.chainID+app0.chainID,
		ch1to0,
		app1,
	)

	suite.openChannels(
		app2.chainID,
		app0.chainID+app2.chainID,
		ch0to2,
		app0,

		app0.chainID,
		app2.chainID+app1.chainID,
		ch2to0,
		app2,
	)

	err = app0.app.CrosscccKeeper.MulticastInitiatePacket(
		app0.ctx,
		initiator,
		msg,
		msg.StateTransitions,
	)
	suite.NoError(err) // successfully executed

	ci, found := app0.app.CrosscccKeeper.GetCoordinator(app0.ctx, msg.GetTxID())
	if suite.True(found) {
		suite.Equal(ci.Status, crossccc.CO_STATUS_INIT)
	}

	nextSeqSend := uint64(1)
	packetCommitment := app0.app.IBCKeeper.ChannelKeeper.GetPacketCommitment(app0.ctx, ch0to1.Port, ch0to1.Channel, nextSeqSend)
	suite.NotNil(packetCommitment)
	packetCommitment = app0.app.IBCKeeper.ChannelKeeper.GetPacketCommitment(app0.ctx, ch0to2.Port, ch0to2.Channel, nextSeqSend)
	suite.NotNil(packetCommitment)
}
