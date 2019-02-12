package helper

import (
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/maticnetwork/heimdall/contracts/rootchain"
	"github.com/maticnetwork/heimdall/contracts/stakemanager"
	"github.com/maticnetwork/heimdall/types"
)

type IContractCaller interface {
	GetHeaderInfo(headerID uint64) (root common.Hash, start uint64, end uint64, err error)
	GetValidatorInfo(valID types.ValidatorID) (validator types.Validator, err error)
	CurrentChildBlock() (uint64, error)
	GetBalance(address common.Address) (*big.Int, error)
	SendCheckpoint(voteSignBytes []byte, sigs []byte, txData []byte)
}

type ContractCaller struct {
	rootChainInstance    *rootchain.Rootchain
	stakeManagerInstance *stakemanager.Stakemanager
	mainChainClient      *ethclient.Client
}

func NewContractCaller() (contractCallerObj ContractCaller, err error) {
	rootChainInstance, err := GetRootChainInstance()
	if err != nil {
		Logger.Error("Error creating rootchain instance", "error", err)
		return contractCallerObj, err
	}
	stakeManagerInstance, err := GetStakeManagerInstance()
	if err != nil {
		Logger.Error("Error creating stakeManagerInstance while getting validator info", "error", err)
		return contractCallerObj, err
	}
	contractCallerObj.mainChainClient = GetMainClient()
	contractCallerObj.stakeManagerInstance = stakeManagerInstance
	contractCallerObj.rootChainInstance = rootChainInstance
	return contractCallerObj, nil
}

// GetHeaderInfo get header info from header id
func (c *ContractCaller) GetHeaderInfo(headerID uint64) (root common.Hash, start uint64, end uint64, err error) {
	// get header from rootchain
	headerIDInt := big.NewInt(0)
	headerIDInt = headerIDInt.SetUint64(headerID)
	headerBlock, err := c.rootChainInstance.HeaderBlock(nil, headerIDInt)
	if err != nil {
		Logger.Error("Unable to fetch header block from rootchain", "headerBlockIndex", headerID)
	}

	return headerBlock.Root, headerBlock.Start.Uint64(), headerBlock.End.Uint64(), nil
}

// CurrentChildBlock fetch current child block
func (c *ContractCaller) CurrentChildBlock() (uint64, error) {
	currentChildBlock, err := c.rootChainInstance.CurrentChildBlock(nil)
	if err != nil {
		Logger.Error("Could not fetch current child block from rootchain contract", "Error", err)
		return 0, err
	}
	return currentChildBlock.Uint64(), nil
}

// get balance of account (returns big.Int balance wont fit in uint64)
func (c *ContractCaller) GetBalance(address common.Address) (*big.Int, error) {
	balance, err := c.mainChainClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		Logger.Error("Unable to fetch balance of account from root chain", "Error", err, "Address", address.String())
		return big.NewInt(0), err
	}

	return balance, nil
}

// GetValidatorInfo get validator info
func (c *ContractCaller) GetValidatorInfo(valID types.ValidatorID) (validator types.Validator, err error) {
	amount, startEpoch, endEpoch, signer, err := c.stakeManagerInstance.GetStakerDetails(nil, big.NewInt(int64(valID)))
	if err != nil {
		Logger.Error("Error fetching validator information from stake manager", "Error", err, "ValidatorID", valID)
		return
	}

	validator = types.Validator{
		ID:         valID,
		Power:      amount.Uint64(),
		StartEpoch: startEpoch.Uint64(),
		EndEpoch:   endEpoch.Uint64(),
		Signer:     signer,
	}

	return validator, nil
}
