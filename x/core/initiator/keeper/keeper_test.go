package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/golang/protobuf/proto"

	samplemodtypes "github.com/datachainlab/cross/simapp/samplemod/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	initiatortypes "github.com/datachainlab/cross/x/core/initiator/types"
	txtypes "github.com/datachainlab/cross/x/core/tx/types"
	crosstypes "github.com/datachainlab/cross/x/core/types"
	xcctypes "github.com/datachainlab/cross/x/core/xcc/types"
	ibctesting "github.com/datachainlab/cross/x/ibc/testing"
	"github.com/datachainlab/cross/x/packets"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	queryClient transfertypes.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.App.InterfaceRegistry())
	transfertypes.RegisterQueryServer(queryHelper, suite.chainA.App.TransferKeeper)
	suite.queryClient = transfertypes.NewQueryClient(queryHelper)
}

func (suite *KeeperTestSuite) TestInitiateTx() {
	// setup

	_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint, ibctesting.CrossVersion)
	channelA, channelB := suite.coordinator.CreateChannel(suite.chainA, suite.chainB, connA, connB, crosstypes.PortID, crosstypes.PortID, channeltypes.UNORDERED)

	chAB := xcctypes.ChannelInfo{Port: channelA.PortID, Channel: channelA.ID}
	xccB, err := xcctypes.PackCrossChainChannel(&chAB)
	suite.Require().NoError(err)

	chBA := xcctypes.ChannelInfo{Port: channelB.PortID, Channel: channelB.ID}

	xccSelf, err := xcctypes.PackCrossChainChannel(suite.chainA.App.XCCResolver.GetSelfCrossChainChannel(suite.chainA.GetContext()))
	suite.Require().NoError(err)

	txs := []initiatortypes.ContractTransaction{
		{
			CrossChainChannel: xccSelf,
			Signers: []authtypes.Account{
				authtypes.NewLocalAccount(authtypes.AccountID(suite.chainA.SenderAccount.GetAddress())),
			},
			CallInfo: samplemodtypes.NewContractCallRequest("nop").ContractCallInfo(suite.chainA.App.AppCodec()),
		},
		{
			CrossChainChannel: xccB,
			Signers: []authtypes.Account{
				authtypes.NewAccount(authtypes.AccountID(suite.chainB.SenderAccount.GetAddress()), authtypes.NewAuthTypeChannelWithAny(xccB)),
			},
			CallInfo: samplemodtypes.NewContractCallRequest("nop").ContractCallInfo(suite.chainB.App.AppCodec()),
		},
	}

	// InitiateTx on chainA

	msg0 := &initiatortypes.MsgInitiateTx{
		ChainId:              suite.chainA.ChainID,
		Nonce:                0,
		CommitProtocol:       txtypes.COMMIT_PROTOCOL_SIMPLE,
		ContractTransactions: txs,
		Signers: []authtypes.Account{
			authtypes.NewLocalAccount(authtypes.AccountID(suite.chainA.SenderAccount.GetAddress())),
		},
		TimeoutHeight: clienttypes.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height)+100),
	}
	suite.Require().NoError(msg0.ValidateBasic())

	ctx := suite.chainA.GetContext()
	res0, err := suite.chainA.App.CrossKeeper.InitiatorKeeper().InitiateTx(
		sdk.WrapSDKContext(ctx),
		msg0,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(res0.Status, initiatortypes.INITIATE_TX_STATUS_PENDING)
	{
		ps, err := ibctesting.GetPacketsFromEvents(ctx.EventManager().ABCIEvents())
		suite.Require().NoError(err)
		suite.Require().Len(ps, 0)
	}
	suite.chainA.NextBlock()

	// IBCSignTx on chainB
	ps := ibctesting.NewCapturePacketSender(
		packets.NewBasicPacketSender(suite.chainB.App.IBCKeeper.ChannelKeeper),
	)
	err = suite.chainB.App.CrossKeeper.AuthKeeper().SendIBCSignTx(
		suite.chainB.GetContext(),
		ps,
		&chBA,
		res0.TxID,
		[]authtypes.AccountID{suite.chainB.SenderAccount.GetAddress().Bytes()},
		clienttypes.NewHeight(0, uint64(suite.chainB.CurrentHeader.Height)+100),
		0,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(ps.Packets()))
	suite.chainB.NextBlock()

	// Receive PacketIBCSignTx on chainA
	p0 := ps.Packets()[0]
	res1, _, err := suite.chainA.App.CrossKeeper.AuthKeeper().HandlePacket(
		suite.chainA.GetContext(),
		channeltypes.Packet{DestinationPort: p0.GetDestPort(), DestinationChannel: p0.GetDestChannel()},
		p0,
	)
	suite.Require().NoError(err)
	suite.chainA.NextBlock()
	{
		ps, err := ibctesting.GetPacketsFromEvents(res1.GetEvents().ToABCIEvents())
		suite.Require().NoError(err)
		suite.Require().Len(ps, 1)
	}

	// Re-send IBCSignTx to chainB
	ps = ibctesting.NewCapturePacketSender(
		packets.NewBasicPacketSender(suite.chainB.App.IBCKeeper.ChannelKeeper),
	)
	err = suite.chainB.App.CrossKeeper.AuthKeeper().SendIBCSignTx(
		suite.chainB.GetContext(),
		ps,
		&chBA,
		res0.TxID,
		[]authtypes.AccountID{suite.chainB.SenderAccount.GetAddress().Bytes()},
		clienttypes.NewHeight(0, uint64(suite.chainB.CurrentHeader.Height)+100),
		0,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(ps.Packets()))
	suite.chainB.NextBlock()

	// Receive PacketIBCSignTx contains duplicate txID must be fail
	p1 := ps.Packets()[0]
	_, ack, err := suite.chainA.App.CrossKeeper.AuthKeeper().HandlePacket(
		suite.chainA.GetContext(),
		channeltypes.Packet{DestinationPort: p1.GetDestPort(), DestinationChannel: p1.GetDestChannel()},
		p1,
	)
	suite.Require().NoError(err)
	ackBz, err := proto.Marshal(ack)
	suite.Require().NoError(err)
	inAck, err := packets.UnmarshalIncomingPacketAcknowledgement(suite.chainA.App.AppCodec(), ackBz)
	suite.Require().NoError(err)
	suite.Require().Equal(authtypes.IBC_SIGN_TX_STATUS_FAILED, inAck.Payload().(*authtypes.PacketAcknowledgementIBCSignTx).Status)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
