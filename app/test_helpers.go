// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package app

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"

	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/cosmos/ibc-go/v6/testing/mock"

	abci "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/version"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/evmos/evmos/v11/cmd/config"
	"github.com/evmos/evmos/v11/encoding"
	"github.com/evmos/evmos/v11/tests"
	evmostypes "github.com/evmos/evmos/v11/types"
	"github.com/evmos/evmos/v11/utils"
	evmtypes "github.com/evmos/evmos/v11/x/evm/types"
	feemarkettypes "github.com/evmos/evmos/v11/x/feemarket/types"
)

func init() {
	cfg := sdk.GetConfig()
	config.SetBech32Prefixes(cfg)
	config.SetBip44CoinType(cfg)
}

// DefaultTestingAppInit defines the IBC application used for testing
var DefaultTestingAppInit func() (ibctesting.TestingApp, map[string]json.RawMessage) = SetupTestingApp

// DefaultConsensusParams defines the default Tendermint consensus params used in
// Evmos testing.
var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   -1, // no limit
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func init() {
	feemarkettypes.DefaultMinGasPrice = sdk.ZeroDec()
	cfg := sdk.GetConfig()
	config.SetBech32Prefixes(cfg)
	config.SetBip44CoinType(cfg)
}

type AccountType int8

const (
	AccountType_EOA = iota
	AccountType_Contract
	AccountType_Validator
)

type Account struct {
	Address    sdk.AccAddress
	AddressHex common.Address
	PubKey     cryptotypes.PubKey
	TmPubKey   tmcrypto.PubKey
	PrivKey    cryptotypes.PrivKey
	Type       AccountType
}

func NewEOAAccount(t testing.TB) Account {
	addr, privKey := tests.NewAddrKey()
	return Account{
		Address:    addr.Bytes(),
		AddressHex: addr,
		PubKey:     privKey.PubKey(),
		PrivKey:    privKey,
		Type:       AccountType_EOA,
	}
}

func NewValidatorAccount(t testing.TB) Account {
	privVal := mock.NewPV()
	pubKey := privVal.PrivKey.PubKey()
	tmPubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	addr := sdk.AccAddress(pubKey.Address())
	return Account{
		Address:    addr,
		AddressHex: common.BytesToAddress(addr.Bytes()),
		PubKey:     privVal.PrivKey.PubKey(),
		PrivKey:    privVal.PrivKey,
		Type:       AccountType_Validator,
		TmPubKey:   tmPubKey,
	}
}

func SetupWithOptions(
	t testing.TB,
	numValidators,
	numAccounts uint64,
) (
	app *Evmos,
	ctx sdk.Context,
	accounts,
	validatorAccounts []Account,
) {
	options := simapp.SetupOptions{
		Logger:             log.NewNopLogger(),
		DB:                 dbm.NewMemDB(),
		InvCheckPeriod:     0,
		HomePath:           DefaultNodeHome,
		SkipUpgradeHeights: nil,
		EncConfig:          encoding.MakeConfig(ModuleBasics),
		AppOpts:            simapp.EmptyAppOptions{},
	}

	genesis := NewDefaultGenesisState()
	return CustomSetup(t, numValidators, numAccounts, options, genesis)
}

// Setup initializes a new Evmos. A Nop logger is set in Evmos.
func CustomSetup(
	t testing.TB,
	numValidators,
	numAccounts uint64,
	options simapp.SetupOptions,
	genesis simapp.GenesisState,
) (
	app *Evmos,
	ctx sdk.Context,
	accounts,
	validatorAccounts []Account,
) {
	validatorAccounts = make([]Account, numValidators)
	tmValidators := make([]*tmtypes.Validator, numValidators)

	for i := 0; i < int(numValidators); i++ {
		validatorAcc := NewValidatorAccount(t)
		tmValidator := tmtypes.NewValidator(validatorAcc.TmPubKey, sdk.TokensToConsensusPower(sdk.OneInt(), evmostypes.PowerReduction))
		validatorAccounts[i] = validatorAcc
		tmValidators[i] = tmValidator
	}

	valSet := tmtypes.NewValidatorSet(tmValidators)

	accounts = make([]Account, numAccounts)
	genAccounts := make([]authtypes.GenesisAccount, numAccounts)
	balances := make([]banktypes.Balance, numAccounts)

	for i := uint64(0); i < numAccounts; i++ {
		acc := NewEOAAccount(t)
		accounts[i] = acc

		baseAcc := authtypes.NewBaseAccount(acc.Address, acc.PubKey, i, 0)
		genAccounts[i] = &evmostypes.EthAccount{BaseAccount: baseAcc, CodeHash: common.BytesToHash(evmtypes.EmptyCodeHash).Hex()}

		balances[i] = banktypes.Balance{
			Address: acc.Address.String(),
			Coins:   sdk.Coins{evmostypes.NewAEvmosCoin(sdk.NewInt(1).Mul(evmostypes.PowerReduction))},
		}
	}

	app = NewEvmos(
		options.Logger,
		options.DB,
		nil,
		true,
		options.SkipUpgradeHeights,
		options.HomePath,
		options.InvCheckPeriod,
		options.EncConfig,
		options.AppOpts,
		// baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(options.AppOpts.Get(sdkserver.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(options.AppOpts.Get(sdkserver.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(options.AppOpts.Get(sdkserver.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(options.AppOpts.Get(sdkserver.FlagMinRetainBlocks))),
		// baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(options.AppOpts.Get(sdkserver.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(options.AppOpts.Get(sdkserver.FlagIndexEvents))),
		// baseapp.SetSnapshot(snapshotStore, snapshotOptions),
		baseapp.SetIAVLCacheSize(cast.ToInt(options.AppOpts.Get(sdkserver.FlagIAVLCacheSize))),
		baseapp.SetIAVLDisableFastNode(cast.ToBool(options.AppOpts.Get(sdkserver.FlagDisableIAVLFastNode))),
	)

	// init chain must be called to stop deliverState from being nil
	genesis = GenesisStateWithValSet(app, genesis, valSet, genAccounts, balances...)

	// // Verify feeMarket genesis
	// if feemarketGenesis != nil {
	// 	err := feemarketGenesis.Validate()
	// 	require.NoError(t, err)

	// 	genesis[feemarkettypes.ModuleName] = app.AppCodec().MustMarshalJSON(feemarketGenesis)
	// }

	stateBytes, err := json.MarshalIndent(genesis, "", " ")
	require.NoError(t, err)

	// Initialize the chain
	req := abci.RequestInitChain{
		ChainId:         utils.MainnetChainID + "-1",
		Validators:      []abci.ValidatorUpdate{},
		ConsensusParams: DefaultConsensusParams,
		AppStateBytes:   stateBytes,
	}

	res := app.InitChain(req)
	header := NewHeader(1, time.Now().UTC(), req.ChainId, validatorAccounts[0].Address.Bytes(), res.AppHash)
	ctx = app.NewContext(false, header)
	return app, ctx, accounts, validatorAccounts
}

func NewHeader(
	height int64,
	blockTime time.Time,
	chainID string,
	proposer sdk.ConsAddress,
	appHash []byte,
) tmproto.Header {
	return tmproto.Header{
		Height:          height,
		ChainID:         chainID,
		Time:            blockTime,
		ProposerAddress: proposer.Bytes(),
		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            appHash,
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	}
}

func GenesisStateWithValSet(
	app *Evmos,
	genesis simapp.GenesisState,
	valSet *tmtypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) simapp.GenesisState {
	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesis[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, _ := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		pkAny, _ := codectypes.NewAnyWithValue(pk)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

	}
	// set validators and delegations
	stakingparams := stakingtypes.DefaultParams()
	stakingparams.BondDenom = utils.BaseDenom
	stakingGenesis := stakingtypes.NewGenesisState(stakingparams, validators, delegations)
	genesis[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(utils.BaseDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(utils.BaseDenom, bondAmt)},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesis[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesis
}

// SetupTestingApp initializes the IBC-go testing application
func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	// FIXME: update to use the new testing setup
	db := dbm.NewMemDB()
	cfg := encoding.MakeConfig(ModuleBasics)
	app := NewEvmos(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 5, cfg, simapp.EmptyAppOptions{})
	return app, NewDefaultGenesisState()
}
