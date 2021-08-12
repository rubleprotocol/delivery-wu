// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/maticnetwork/bor/common"
	erc20 "github.com/maticnetwork/heimdall/contracts/erc20"

	heimdalltypes "github.com/maticnetwork/heimdall/types"

	mock "github.com/stretchr/testify/mock"

	rootchain "github.com/maticnetwork/heimdall/contracts/rootchain"

	slashmanager "github.com/maticnetwork/heimdall/contracts/slashmanager"

	stakemanager "github.com/maticnetwork/heimdall/contracts/stakemanager"

	stakinginfo "github.com/maticnetwork/heimdall/contracts/stakinginfo"

	statereceiver "github.com/maticnetwork/heimdall/contracts/statereceiver"

	statesender "github.com/maticnetwork/heimdall/contracts/statesender"

	types "github.com/maticnetwork/bor/core/types"

	validatorset "github.com/maticnetwork/heimdall/contracts/validatorset"
)

// IContractCaller is an autogenerated mock type for the IContractCaller type
type IContractCaller struct {
	mock.Mock
}

// ApproveTokens provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *IContractCaller) ApproveTokens(_a0 *big.Int, _a1 common.Address, _a2 common.Address, _a3 *erc20.Erc20) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, common.Address, common.Address, *erc20.Erc20) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckIfBlocksExist provides a mock function with given fields: end
func (_m *IContractCaller) CheckIfBlocksExist(end uint64) bool {
	ret := _m.Called(end)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(end)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CurrentAccountStateRoot provides a mock function with given fields: stakingInfoInstance
func (_m *IContractCaller) CurrentAccountStateRoot(stakingInfoInstance *stakinginfo.Stakinginfo) ([32]byte, error) {
	ret := _m.Called(stakingInfoInstance)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*stakinginfo.Stakinginfo) [32]byte); ok {
		r0 = rf(stakingInfoInstance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*stakinginfo.Stakinginfo) error); ok {
		r1 = rf(stakingInfoInstance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentHeaderBlock provides a mock function with given fields: rootChainInstance, childBlockInterval
func (_m *IContractCaller) CurrentHeaderBlock(rootChainInstance *rootchain.Rootchain, childBlockInterval uint64) (uint64, error) {
	ret := _m.Called(rootChainInstance, childBlockInterval)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*rootchain.Rootchain, uint64) uint64); ok {
		r0 = rf(rootChainInstance, childBlockInterval)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rootchain.Rootchain, uint64) error); ok {
		r1 = rf(rootChainInstance, childBlockInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentSpanNumber provides a mock function with given fields: _a0
func (_m *IContractCaller) CurrentSpanNumber(_a0 *validatorset.Validatorset) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*validatorset.Validatorset) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// CurrentStateCounter provides a mock function with given fields: stateSenderInstance
func (_m *IContractCaller) CurrentStateCounter(stateSenderInstance *statesender.Statesender) *big.Int {
	ret := _m.Called(stateSenderInstance)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*statesender.Statesender) *big.Int); ok {
		r0 = rf(stateSenderInstance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// DecodeNewHeaderBlockEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeNewHeaderBlockEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*rootchain.RootchainNewHeaderBlock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *rootchain.RootchainNewHeaderBlock
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *rootchain.RootchainNewHeaderBlock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rootchain.RootchainNewHeaderBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeSignerUpdateEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeSignerUpdateEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoSignerChange, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoSignerChange
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoSignerChange); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoSignerChange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeSlashedEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeSlashedEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoSlashed, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoSlashed
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoSlashed); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoSlashed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeStateSyncedEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeStateSyncedEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*statesender.StatesenderStateSynced, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *statesender.StatesenderStateSynced
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *statesender.StatesenderStateSynced); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*statesender.StatesenderStateSynced)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeUnJailedEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeUnJailedEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoUnJailed, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoUnJailed
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoUnJailed); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoUnJailed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorExitEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeValidatorExitEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoUnstakeInit, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoUnstakeInit
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoUnstakeInit); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoUnstakeInit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorJoinEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeValidatorJoinEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoStaked, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoStaked
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoStaked); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoStaked)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorStakeUpdateEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeValidatorStakeUpdateEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoStakeUpdate, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoStakeUpdate
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoStakeUpdate); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoStakeUpdate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeValidatorTopupFeesEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *IContractCaller) DecodeValidatorTopupFeesEvent(_a0 common.Address, _a1 *types.Receipt, _a2 uint64) (*stakinginfo.StakinginfoTopUpFee, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *stakinginfo.StakinginfoTopUpFee
	if rf, ok := ret.Get(0).(func(common.Address, *types.Receipt, uint64) *stakinginfo.StakinginfoTopUpFee); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.StakinginfoTopUpFee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, *types.Receipt, uint64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: address
func (_m *IContractCaller) GetBalance(address common.Address) (*big.Int, error) {
	ret := _m.Called(address)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockNumberFromTxHash provides a mock function with given fields: _a0
func (_m *IContractCaller) GetBlockNumberFromTxHash(_a0 common.Hash) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Hash) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCheckpointSign provides a mock function with given fields: txHash
func (_m *IContractCaller) GetCheckpointSign(txHash common.Hash) ([]byte, []byte, []byte, error) {
	ret := _m.Called(txHash)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Hash) []byte); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(common.Hash) []byte); ok {
		r1 = rf(txHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 []byte
	if rf, ok := ret.Get(2).(func(common.Hash) []byte); ok {
		r2 = rf(txHash)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]byte)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(common.Hash) error); ok {
		r3 = rf(txHash)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetConfirmedTxReceipt provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) GetConfirmedTxReceipt(_a0 common.Hash, _a1 uint64) (*types.Receipt, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash, uint64) *types.Receipt); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash, uint64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTronEventsByContractAddress provides a mock function with given fields: address, from, to
func (_m *IContractCaller) GetTronEventsByContractAddress(address []string, from int64, to int64) ([]types.Log, error) {
	ret := _m.Called(address, from, to)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func([]string, int64, int64) []types.Log); ok {
		r0 = rf(address, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, int64, int64) error); ok {
		r1 = rf(address, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHeaderInfo provides a mock function with given fields: headerID, rootChainInstance, childBlockInterval
func (_m *IContractCaller) GetHeaderInfo(headerID uint64, rootChainInstance *rootchain.Rootchain, childBlockInterval uint64) (common.Hash, uint64, uint64, uint64, heimdalltypes.HeimdallAddress, error) {
	ret := _m.Called(headerID, rootChainInstance, childBlockInterval)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(uint64, *rootchain.Rootchain, uint64) common.Hash); ok {
		r0 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(uint64, *rootchain.Rootchain, uint64) uint64); ok {
		r1 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 uint64
	if rf, ok := ret.Get(2).(func(uint64, *rootchain.Rootchain, uint64) uint64); ok {
		r2 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	var r3 uint64
	if rf, ok := ret.Get(3).(func(uint64, *rootchain.Rootchain, uint64) uint64); ok {
		r3 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		r3 = ret.Get(3).(uint64)
	}

	var r4 heimdalltypes.HeimdallAddress
	if rf, ok := ret.Get(4).(func(uint64, *rootchain.Rootchain, uint64) heimdalltypes.HeimdallAddress); ok {
		r4 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).(heimdalltypes.HeimdallAddress)
		}
	}

	var r5 error
	if rf, ok := ret.Get(5).(func(uint64, *rootchain.Rootchain, uint64) error); ok {
		r5 = rf(headerID, rootChainInstance, childBlockInterval)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// GetLastChildBlock provides a mock function with given fields: rootChainInstance
func (_m *IContractCaller) GetLastChildBlock(rootChainInstance *rootchain.Rootchain) (uint64, error) {
	ret := _m.Called(rootChainInstance)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*rootchain.Rootchain) uint64); ok {
		r0 = rf(rootChainInstance)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rootchain.Rootchain) error); ok {
		r1 = rf(rootChainInstance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainChainBlock provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMainChainBlock(_a0 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainStakingSyncNonce provides a mock function with given fields: validatorID, stakingManagerInstance
func (_m *IContractCaller) GetMainStakingSyncNonce(validatorID uint64, stakingManagerInstance *stakemanager.Stakemanager) uint64 {
	ret := _m.Called(validatorID, stakingManagerInstance)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64, *stakemanager.Stakemanager) uint64); ok {
		r0 = rf(validatorID, stakingManagerInstance)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetMainTxReceipt provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMainTxReceipt(_a0 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticChainBlock provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMaticChainBlock(_a0 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*big.Int) *types.Header); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticTokenInstance provides a mock function with given fields: maticTokenAddress
func (_m *IContractCaller) GetMaticTokenInstance(maticTokenAddress common.Address) (*erc20.Erc20, error) {
	ret := _m.Called(maticTokenAddress)

	var r0 *erc20.Erc20
	if rf, ok := ret.Get(0).(func(common.Address) *erc20.Erc20); ok {
		r0 = rf(maticTokenAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*erc20.Erc20)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(maticTokenAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaticTxReceipt provides a mock function with given fields: _a0
func (_m *IContractCaller) GetMaticTxReceipt(_a0 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootChainInstance provides a mock function with given fields: rootchainAddress
func (_m *IContractCaller) GetRootChainInstance(rootchainAddress common.Address) (*rootchain.Rootchain, error) {
	ret := _m.Called(rootchainAddress)

	var r0 *rootchain.Rootchain
	if rf, ok := ret.Get(0).(func(common.Address) *rootchain.Rootchain); ok {
		r0 = rf(rootchainAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rootchain.Rootchain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(rootchainAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootHash provides a mock function with given fields: start, end, checkpointLength
func (_m *IContractCaller) GetRootHash(start uint64, end uint64, checkpointLength uint64) ([]byte, error) {
	ret := _m.Called(start, end, checkpointLength)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(uint64, uint64, uint64) []byte); ok {
		r0 = rf(start, end, checkpointLength)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64, uint64) error); ok {
		r1 = rf(start, end, checkpointLength)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlashManagerInstance provides a mock function with given fields: slashManagerAddress
func (_m *IContractCaller) GetSlashManagerInstance(slashManagerAddress common.Address) (*slashmanager.Slashmanager, error) {
	ret := _m.Called(slashManagerAddress)

	var r0 *slashmanager.Slashmanager
	if rf, ok := ret.Get(0).(func(common.Address) *slashmanager.Slashmanager); ok {
		r0 = rf(slashManagerAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slashmanager.Slashmanager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(slashManagerAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpanDetails provides a mock function with given fields: id, _a1
func (_m *IContractCaller) GetSpanDetails(id *big.Int, _a1 *validatorset.Validatorset) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(id, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *validatorset.Validatorset) *big.Int); ok {
		r0 = rf(id, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(*big.Int, *validatorset.Validatorset) *big.Int); ok {
		r1 = rf(id, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 *big.Int
	if rf, ok := ret.Get(2).(func(*big.Int, *validatorset.Validatorset) *big.Int); ok {
		r2 = rf(id, _a1)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*big.Int, *validatorset.Validatorset) error); ok {
		r3 = rf(id, _a1)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetStakeManagerInstance provides a mock function with given fields: stakingManagerAddress
func (_m *IContractCaller) GetStakeManagerInstance(stakingManagerAddress common.Address) (*stakemanager.Stakemanager, error) {
	ret := _m.Called(stakingManagerAddress)

	var r0 *stakemanager.Stakemanager
	if rf, ok := ret.Get(0).(func(common.Address) *stakemanager.Stakemanager); ok {
		r0 = rf(stakingManagerAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakemanager.Stakemanager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(stakingManagerAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakingInfoInstance provides a mock function with given fields: stakingInfoAddress
func (_m *IContractCaller) GetStakingInfoInstance(stakingInfoAddress common.Address) (*stakinginfo.Stakinginfo, error) {
	ret := _m.Called(stakingInfoAddress)

	var r0 *stakinginfo.Stakinginfo
	if rf, ok := ret.Get(0).(func(common.Address) *stakinginfo.Stakinginfo); ok {
		r0 = rf(stakingInfoAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stakinginfo.Stakinginfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(stakingInfoAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateReceiverInstance provides a mock function with given fields: stateReceiverAddress
func (_m *IContractCaller) GetStateReceiverInstance(stateReceiverAddress common.Address) (*statereceiver.Statereceiver, error) {
	ret := _m.Called(stateReceiverAddress)

	var r0 *statereceiver.Statereceiver
	if rf, ok := ret.Get(0).(func(common.Address) *statereceiver.Statereceiver); ok {
		r0 = rf(stateReceiverAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*statereceiver.Statereceiver)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(stateReceiverAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateSenderInstance provides a mock function with given fields: stateSenderAddress
func (_m *IContractCaller) GetStateSenderInstance(stateSenderAddress common.Address) (*statesender.Statesender, error) {
	ret := _m.Called(stateSenderAddress)

	var r0 *statesender.Statesender
	if rf, ok := ret.Get(0).(func(common.Address) *statesender.Statesender); ok {
		r0 = rf(stateSenderAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*statesender.Statesender)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(stateSenderAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTronTransactionReceipt provides a mock function with given fields: txID
func (_m *IContractCaller) GetTronTransactionReceipt(txID string) (*types.Receipt, error) {
	ret := _m.Called(txID)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(string) *types.Receipt); ok {
		r0 = rf(txID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(txID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTronStakingSyncNonce provides a mock function with given fields: validatorID, stakingManagerAddress
func (_m *IContractCaller) GetTronStakingSyncNonce(validatorID uint64, stakingManagerAddress string) uint64 {
	ret := _m.Called(validatorID, stakingManagerAddress)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64, string) uint64); ok {
		r0 = rf(validatorID, stakingManagerAddress)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetValidatorInfo provides a mock function with given fields: valID, stakingInfoInstance
func (_m *IContractCaller) GetValidatorInfo(valID heimdalltypes.ValidatorID, stakingInfoInstance *stakinginfo.Stakinginfo) (heimdalltypes.Validator, error) {
	ret := _m.Called(valID, stakingInfoInstance)

	var r0 heimdalltypes.Validator
	if rf, ok := ret.Get(0).(func(heimdalltypes.ValidatorID, *stakinginfo.Stakinginfo) heimdalltypes.Validator); ok {
		r0 = rf(valID, stakingInfoInstance)
	} else {
		r0 = ret.Get(0).(heimdalltypes.Validator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(heimdalltypes.ValidatorID, *stakinginfo.Stakinginfo) error); ok {
		r1 = rf(valID, stakingInfoInstance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidatorSetInstance provides a mock function with given fields: validatorSetAddress
func (_m *IContractCaller) GetValidatorSetInstance(validatorSetAddress common.Address) (*validatorset.Validatorset, error) {
	ret := _m.Called(validatorSetAddress)

	var r0 *validatorset.Validatorset
	if rf, ok := ret.Get(0).(func(common.Address) *validatorset.Validatorset); ok {
		r0 = rf(validatorSetAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validatorset.Validatorset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(validatorSetAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTxConfirmed provides a mock function with given fields: _a0, _a1
func (_m *IContractCaller) IsTxConfirmed(_a0 common.Hash, _a1 uint64) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Hash, uint64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SendCheckpoint provides a mock function with given fields: sigedData, sigs, rootchainAddress, rootChainInstance
func (_m *IContractCaller) SendCheckpoint(sigedData []byte, sigs [][3]*big.Int, rootchainAddress common.Address, rootChainInstance *rootchain.Rootchain) error {
	ret := _m.Called(sigedData, sigs, rootchainAddress, rootChainInstance)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, [][3]*big.Int, common.Address, *rootchain.Rootchain) error); ok {
		r0 = rf(sigedData, sigs, rootchainAddress, rootChainInstance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMainStakingSync provides a mock function with given fields: stakingType, sigedData, sigs, stakingManagerAddress, stakingManagerInstance
func (_m *IContractCaller) SendMainStakingSync(stakingType string, sigedData []byte, sigs [][3]*big.Int, stakingManagerAddress common.Address, stakingManagerInstance *stakemanager.Stakemanager) error {
	ret := _m.Called(stakingType, sigedData, sigs, stakingManagerAddress, stakingManagerInstance)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, [][3]*big.Int, common.Address, *stakemanager.Stakemanager) error); ok {
		r0 = rf(stakingType, sigedData, sigs, stakingManagerAddress, stakingManagerInstance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendTick provides a mock function with given fields: sigedData, sigs, slashManagerAddress, slashManagerInstance
func (_m *IContractCaller) SendTick(sigedData []byte, sigs []byte, slashManagerAddress common.Address, slashManagerInstance *slashmanager.Slashmanager) error {
	ret := _m.Called(sigedData, sigs, slashManagerAddress, slashManagerInstance)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte, common.Address, *slashmanager.Slashmanager) error); ok {
		r0 = rf(sigedData, sigs, slashManagerAddress, slashManagerInstance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendTronCheckpoint provides a mock function with given fields: signedData, sigs, rootChainAddress
func (_m *IContractCaller) SendTronCheckpoint(signedData []byte, sigs [][3]*big.Int, rootChainAddress string) (string, error) {
	ret := _m.Called(signedData, sigs, rootChainAddress)

	var r0 string
	if rf, ok := ret.Get(0).(func([]byte, [][3]*big.Int, string) string); ok {
		r0 = rf(signedData, sigs, rootChainAddress)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, [][3]*big.Int, string) error); ok {
		r1 = rf(signedData, sigs, rootChainAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTronStakingSync provides a mock function with given fields: stakingType, sigedData, sigs, stakingManagerAddress
func (_m *IContractCaller) SendTronStakingSync(stakingType string, sigedData []byte, sigs [][3]*big.Int, stakingManagerAddress string) error {
	ret := _m.Called(stakingType, sigedData, sigs, stakingManagerAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, [][3]*big.Int, string) error); ok {
		r0 = rf(stakingType, sigedData, sigs, stakingManagerAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StakeFor provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *IContractCaller) StakeFor(_a0 common.Address, _a1 *big.Int, _a2 *big.Int, _a3 bool, _a4 common.Address, _a5 *stakemanager.Stakemanager) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, *big.Int, *big.Int, bool, common.Address, *stakemanager.Stakemanager) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
