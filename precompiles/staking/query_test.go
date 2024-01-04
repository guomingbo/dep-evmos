package staking_test

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v16/precompiles/authorization"
	cmn "github.com/evmos/evmos/v16/precompiles/common"
	"github.com/evmos/evmos/v16/precompiles/staking"
	testutiltx "github.com/evmos/evmos/v16/testutil/tx"
)

func (s *PrecompileTestSuite) TestDelegation() {
	method := s.precompile.Methods[staking.DelegationMethod]

	testCases := []struct {
		name        string
		malleate    func(operatorAddress common.Address) []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 2, 0),
		},
		{
			"fail - invalid delegator address",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					"invalid",
					operatorAddress,
				}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidDelegator, "invalid"),
		},
		{
			"fail - invalid operator address",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					s.address,
					"invalid",
				}
			},
			func(bz []byte) {},
			100000,
			true,
			"invalid validator addres",
		},
		{
			"success - empty delegation",
			func(operatorAddress common.Address) []interface{} {
				addr, _ := testutiltx.NewAddrKey()
				return []interface{}{
					addr,
					operatorAddress,
				}
			},
			func(bz []byte) {
				var delOut staking.DelegationOutput
				err := s.precompile.UnpackIntoInterface(&delOut, staking.DelegationMethod, bz)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Equal(delOut.Shares.Int64(), big.NewInt(0).Int64())
			},
			100000,
			false,
			"",
		},
		{
			"success",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					s.address,
					operatorAddress,
				}
			},
			func(bz []byte) {
				var delOut staking.DelegationOutput
				err := s.precompile.UnpackIntoInterface(&delOut, staking.DelegationMethod, bz)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Equal(delOut.Shares, big.NewInt(1e18))
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			operatorAddress, _ := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)

			bz, err := s.precompile.Delegation(s.ctx, contract, &method, tc.malleate(common.BytesToAddress(operatorAddress.Bytes())))

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotEmpty(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestUnbondingDelegation() {
	method := s.precompile.Methods[staking.UnbondingDelegationMethod]

	testCases := []struct {
		name        string
		malleate    func(operatorAddress common.Address) []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 2, 0),
		},
		{
			"fail - invalid delegator address",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					"invalid",
					operatorAddress,
				}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidDelegator, "invalid"),
		},
		{
			"success - no unbonding delegation found",
			func(operatorAddress common.Address) []interface{} {
				addr, _ := testutiltx.NewAddrKey()
				return []interface{}{
					addr,
					operatorAddress,
				}
			},
			func(data []byte) {
				var ubdOut staking.UnbondingDelegationOutput
				err := s.precompile.UnpackIntoInterface(&ubdOut, staking.UnbondingDelegationMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(ubdOut.UnbondingDelegation.Entries, 0)
			},
			100000,
			false,
			"",
		},
		{
			"success",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					s.address,
					operatorAddress,
				}
			},
			func(data []byte) {
				var ubdOut staking.UnbondingDelegationOutput
				err := s.precompile.UnpackIntoInterface(&ubdOut, staking.UnbondingDelegationMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(ubdOut.UnbondingDelegation.Entries, 1)
				s.Require().Equal(ubdOut.UnbondingDelegation.Entries[0].CreationHeight, s.ctx.BlockHeight())
				s.Require().Equal(ubdOut.UnbondingDelegation.Entries[0].Balance, big.NewInt(1e18))
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			_, err := s.app.StakingKeeper.Undelegate(s.ctx, s.address.Bytes(), s.validators[0].GetOperator(), math.LegacyNewDec(1))
			s.Require().NoError(err)

			operatorAddress, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
			s.Require().NoError(err)

			bz, err := s.precompile.UnbondingDelegation(s.ctx, contract, &method, tc.malleate(common.BytesToAddress(operatorAddress.Bytes())))

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestValidator() {
	method := s.precompile.Methods[staking.ValidatorMethod]

	testCases := []struct {
		name        string
		malleate    func(operatorAddress common.Address) []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{}
			},
			func(_ []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 1, 0),
		},
		{
			"success",
			func(operatorAddress common.Address) []interface{} {
				return []interface{}{
					operatorAddress,
				}
			},
			func(data []byte) {
				var valOut staking.ValidatorOutput
				err := s.precompile.UnpackIntoInterface(&valOut, staking.ValidatorMethod, data)
				s.Require().NoError(err, "failed to unpack output")

				operatorAddress, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
				s.Require().NoError(err)

				s.Require().Equal(valOut.Validator.OperatorAddress, common.BytesToAddress(operatorAddress.Bytes()))
			},
			100000,
			false,
			"",
		},
		{
			name: "success - empty validator",
			malleate: func(_ common.Address) []interface{} {
				newAddr, _ := testutiltx.NewAccAddressAndKey()
				newValAddr := sdk.ValAddress(newAddr)
				return []interface{}{
					common.BytesToAddress(newValAddr.Bytes()),
				}
			},
			postCheck: func(data []byte) {
				var valOut staking.ValidatorOutput
				err := s.precompile.UnpackIntoInterface(&valOut, staking.ValidatorMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Equal(valOut.Validator.OperatorAddress.String(), common.Address{}.String())
				s.Require().Equal(valOut.Validator.Status, uint8(0))
			},
			gas: 100000,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			operatorAddress, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
			s.Require().NoError(err)

			bz, err := s.precompile.Validator(s.ctx, &method, contract, tc.malleate(common.BytesToAddress(operatorAddress.Bytes())))

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestValidators() {
	method := s.precompile.Methods[staking.ValidatorsMethod]

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			func(_ []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 2, 0),
		},
		{
			"fail - invalid number of arguments",
			func() []interface{} {
				return []interface{}{
					stakingtypes.Bonded.String(),
				}
			},
			func(_ []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 2, 1),
		},
		{
			"success - bonded status & pagination w/countTotal",
			func() []interface{} {
				return []interface{}{
					stakingtypes.Bonded.String(),
					query.PageRequest{
						Limit:      1,
						CountTotal: true,
					},
				}
			},
			func(data []byte) {
				const expLen = 1
				var valOut staking.ValidatorsOutput
				err := s.precompile.UnpackIntoInterface(&valOut, staking.ValidatorsMethod, data)
				s.Require().NoError(err, "failed to unpack output")

				s.Require().Len(valOut.Validators, expLen)
				// passed CountTotal = true
				s.Require().Equal(len(s.validators), int(valOut.PageResponse.Total))
				s.Require().NotEmpty(valOut.PageResponse.NextKey)
				s.assertValidatorsResponse(valOut.Validators, expLen)
			},
			100000,
			false,
			"",
		},
		{
			"success - bonded status & pagination w/countTotal & key is []byte{0}",
			func() []interface{} {
				return []interface{}{
					stakingtypes.Bonded.String(),
					query.PageRequest{
						Key:        []byte{0},
						Limit:      1,
						CountTotal: true,
					},
				}
			},
			func(data []byte) {
				const expLen = 1
				var valOut staking.ValidatorsOutput
				err := s.precompile.UnpackIntoInterface(&valOut, staking.ValidatorsMethod, data)
				s.Require().NoError(err, "failed to unpack output")

				s.Require().Len(valOut.Validators, expLen)
				// passed CountTotal = true
				s.Require().Equal(len(s.validators), int(valOut.PageResponse.Total))
				s.Require().NotEmpty(valOut.PageResponse.NextKey)
				s.assertValidatorsResponse(valOut.Validators, expLen)
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			bz, err := s.precompile.Validators(s.ctx, &method, contract, tc.malleate())

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestRedelegation() {
	method := s.precompile.Methods[staking.RedelegationMethod]
	redelegateMethod := s.precompile.Methods[staking.RedelegateMethod]

	testCases := []struct {
		name        string
		malleate    func(srcOperatorAddr, destOperatorAddr common.Address) []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func(srcOperatorAddr, destOperatorAddr common.Address) []interface{} {
				return []interface{}{}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 0),
		},
		{
			"fail - invalid delegator address",
			func(srcOperatorAddr, destOperatorAddr common.Address) []interface{} {
				return []interface{}{
					"invalid",
					srcOperatorAddr,
					destOperatorAddr,
				}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidDelegator, "invalid"),
		},
		{
			"fail - empty src validator addr",
			func(srcOperatorAddr, destOperatorAddr common.Address) []interface{} {
				return []interface{}{
					s.address,
					common.Address{},
					destOperatorAddr,
				}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidValidator, common.Address{}),
		},
		{
			"fail - empty destination addr",
			func(srcOperatorAddr, destOperatorAddr common.Address) []interface{} {
				return []interface{}{
					s.address,
					srcOperatorAddr,
					common.Address{},
				}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidValidator, common.Address{}),
		},
		{
			"success",
			func(srcOperatorAddr, destOperatorAddr common.Address) []interface{} {
				return []interface{}{
					s.address,
					srcOperatorAddr,
					destOperatorAddr,
				}
			},
			func(data []byte) {
				var redOut staking.RedelegationOutput
				err := s.precompile.UnpackIntoInterface(&redOut, staking.RedelegationMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(redOut.Redelegation.Entries, 1)
				s.Require().Equal(redOut.Redelegation.Entries[0].CreationHeight, s.ctx.BlockHeight())
				s.Require().Equal(redOut.Redelegation.Entries[0].SharesDst, big.NewInt(1e18))
			},
			100000,
			false,
			"",
		},
		{
			name: "success - no redelegation found",
			malleate: func(srcOperatorAddr, _ common.Address) []interface{} {
				nonExistentOperator := sdk.ValAddress([]byte("non-existent-operator"))
				return []interface{}{
					s.address,
					srcOperatorAddr,
					common.BytesToAddress(nonExistentOperator),
				}
			},
			postCheck: func(data []byte) {
				var redOut staking.RedelegationOutput
				err := s.precompile.UnpackIntoInterface(&redOut, staking.RedelegationMethod, data)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Len(redOut.Redelegation.Entries, 0)
			},
			gas: 100000,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			operatorAddr0, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
			s.Require().NoError(err)
			operatorAddress0 := common.BytesToAddress(operatorAddr0.Bytes())

			operatorAddr1, err := sdk.ValAddressFromBech32(s.validators[1].OperatorAddress)
			s.Require().NoError(err)
			operatorAddress1 := common.BytesToAddress(operatorAddr1.Bytes())

			delegationArgs := []interface{}{
				s.address,
				operatorAddress0,
				operatorAddress1,
				big.NewInt(1e18),
			}

			err = s.CreateAuthorization(s.address, staking.RedelegateAuthz, nil)
			s.Require().NoError(err)

			_, err = s.precompile.Redelegate(s.ctx, s.address, contract, s.stateDB, &redelegateMethod, delegationArgs)
			s.Require().NoError(err)

			bz, err := s.precompile.Redelegation(s.ctx, &method, contract, tc.malleate(operatorAddress0, operatorAddress1))

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestRedelegations() {
	var (
		delAmt                 = big.NewInt(3e17)
		redelTotalCount uint64 = 2
		method                 = s.precompile.Methods[staking.RedelegationsMethod]
	)

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 4, 0),
		},
		{
			"fail - invalid delegator address",
			func() []interface{} {
				operatorAddr0, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress0 := common.BytesToAddress(operatorAddr0.Bytes())

				operatorAddr1, err := sdk.ValAddressFromBech32(s.validators[1].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress1 := common.BytesToAddress(operatorAddr1.Bytes())

				return []interface{}{
					common.BytesToAddress([]byte("invalid")),
					operatorAddress0,
					operatorAddress1,
					query.PageRequest{},
				}
			},
			func(bz []byte) {},
			100000,
			true,
			"redelegation not found",
		},
		{
			"fail - invalid query | all empty args ",
			func() []interface{} {
				return []interface{}{
					common.Address{},
					common.Address{},
					common.Address{},
					query.PageRequest{},
				}
			},
			func(data []byte) {},
			100000,
			true,
			"invalid query. Need to specify at least a source validator address or delegator address",
		},
		{
			"fail - invalid query | only destination validator address",
			func() []interface{} {
				operatorAddr1, err := sdk.ValAddressFromBech32(s.validators[1].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress1 := common.BytesToAddress(operatorAddr1.Bytes())

				return []interface{}{
					common.Address{},
					common.Address{},
					operatorAddress1,
					query.PageRequest{},
				}
			},
			func(data []byte) {},
			100000,
			true,
			"invalid query. Need to specify at least a source validator address or delegator address",
		},
		{
			"success - specified delegator, source & destination",
			func() []interface{} {
				operatorAddr0, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress0 := common.BytesToAddress(operatorAddr0.Bytes())

				operatorAddr1, err := sdk.ValAddressFromBech32(s.validators[1].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress1 := common.BytesToAddress(operatorAddr1.Bytes())

				return []interface{}{
					s.address,
					operatorAddress0,
					operatorAddress1,
					query.PageRequest{},
				}
			},
			func(data []byte) {
				s.assertRedelegationsOutput(data, 0, delAmt, 2, false)
			},
			100000,
			false,
			"",
		},
		{
			"success - specifying only source w/pagination",
			func() []interface{} {
				operatorAddr0, err := sdk.ValAddressFromBech32(s.validators[0].OperatorAddress)
				s.Require().NoError(err)
				operatorAddress0 := common.BytesToAddress(operatorAddr0.Bytes())

				return []interface{}{
					common.Address{},
					operatorAddress0,
					common.Address{},
					query.PageRequest{
						Limit:      1,
						CountTotal: true,
					},
				}
			},
			func(data []byte) {
				s.assertRedelegationsOutput(data, redelTotalCount, delAmt, 2, true)
			},
			100000,
			false,
			"",
		},
		{
			"success - get all existing redelegations for a delegator w/pagination",
			func() []interface{} {
				return []interface{}{
					s.address,
					common.Address{},
					common.Address{},
					query.PageRequest{
						Limit:      1,
						CountTotal: true,
					},
				}
			},
			func(data []byte) {
				s.assertRedelegationsOutput(data, redelTotalCount, delAmt, 2, true)
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			err := s.setupRedelegations(delAmt)
			s.Require().NoError(err)

			// query redelegations
			bz, err := s.precompile.Redelegations(s.ctx, &method, contract, tc.malleate())

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}

func (s *PrecompileTestSuite) TestAllowance() {
	approvedCoin := sdk.Coin{Denom: s.bondDenom, Amount: math.NewInt(1e18)}
	granteeAddr := testutiltx.GenerateAddress()
	method := s.precompile.Methods[authorization.AllowanceMethod]

	testCases := []struct {
		name        string
		malleate    func() []interface{}
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"fail - empty input args",
			func() []interface{} {
				return []interface{}{}
			},
			func(bz []byte) {},
			100000,
			true,
			fmt.Sprintf(cmn.ErrInvalidNumberOfArgs, 3, 0),
		},
		{
			"success - query delegate method allowance",
			func() []interface{} {
				err := s.CreateAuthorization(granteeAddr, staking.DelegateAuthz, &approvedCoin)
				s.Require().NoError(err)

				return []interface{}{
					granteeAddr,
					s.address,
					staking.DelegateMsg,
				}
			},
			func(bz []byte) {
				var amountsOut *big.Int
				err := s.precompile.UnpackIntoInterface(&amountsOut, authorization.AllowanceMethod, bz)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Equal(big.NewInt(1e18), amountsOut, "expected different allowed amount")
			},
			100000,
			false,
			"",
		},
		{
			"success - return empty allowance if authorization is not found",
			func() []interface{} {
				return []interface{}{
					granteeAddr,
					s.address,
					staking.UndelegateMsg,
				}
			},
			func(bz []byte) {
				var amountsOut *big.Int
				err := s.precompile.UnpackIntoInterface(&amountsOut, authorization.AllowanceMethod, bz)
				s.Require().NoError(err, "failed to unpack output")
				s.Require().Equal(int64(0), amountsOut.Int64(), "expected no allowance")
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest() // reset
			contract := vm.NewContract(vm.AccountRef(s.address), s.precompile, big.NewInt(0), tc.gas)

			args := tc.malleate()
			bz, err := s.precompile.Allowance(s.ctx, &method, contract, args)

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}
