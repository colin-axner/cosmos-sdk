package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
)

const (
	chainID                      = "gaia"
	height                       = 4
	trustingPeriod time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod      time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift  time.Duration = time.Second * 10
)

type TendermintTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	cdc     codec.Marshaler
	privVal tmtypes.PrivValidator
	valSet  *tmproto.ValidatorSet
	header  ibctmtypes.Header
	now     time.Time
}

func (suite *TendermintTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)

	suite.cdc = app.AppCodec()
	suite.now = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	suite.privVal = tmtypes.NewMockPV()

	pubKey, err := suite.privVal.GetPubKey()
	suite.Require().NoError(err)

	val := tmtypes.NewValidator(pubKey, 10)
	suite.valSet, err = tmtypes.NewValidatorSet([]*tmtypes.Validator{val}).ToProto()
	suite.Require().NoError(err)
	suite.header = ibctmtypes.CreateTestHeader(chainID, height, suite.now, suite.valSet, []tmtypes.PrivValidator{suite.privVal})
	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1, Time: suite.now})
}

func TestTendermintTestSuite(t *testing.T) {
	suite.Run(t, new(TendermintTestSuite))
}
