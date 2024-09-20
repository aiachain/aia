package systemcontract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

const (
	SysContractV0 SysContractVersion = iota
	SysContractV1
	SysContractV2
	WAIAContractV1
	TokenUpdateV1
	SysContractV3
)

type SysContractVersion int

type IUpgradeAction interface {
	GetName() string
	Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) error
	Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) error
}

func ApplySystemContractUpgrade(version SysContractVersion, state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {
	if config == nil || header == nil || state == nil {
		return
	}
	height := header.Number

	var sysContracts []IUpgradeAction
	switch version {
	case SysContractV0:
		sysContracts = []IUpgradeAction{
			&hardForkAddressList{},
		}
	case SysContractV1:
		sysContracts = []IUpgradeAction{
			&hardForkSysGov{},
			&hardForkValidatorsV1{},
			&hardForkPunishV1{},
		}
	case SysContractV2:
		sysContracts = []IUpgradeAction{
			&hardForkAddressListV2{},
			&hardForkValidatorsV2{},
		}
	case WAIAContractV1:
		sysContracts = []IUpgradeAction{
			&hardForkWAIA{},
		}
	case TokenUpdateV1:
		sysContracts = []IUpgradeAction{
			&hardForkTokens{},
		}
	case SysContractV3:
		sysContracts = []IUpgradeAction{
			&hardForkSysContractV3{},
			&hardForkReward{},
		}
	default:
		log.Crit("unsupported SysContractVersion", "version", version)
	}

	for _, contract := range sysContracts {
		log.Info("system contract upgrade", "version", version, "name", contract.GetName(), "height", height, "chainId", config.GetChainId(header.Number).String())

		err = contract.Update(config, height, state)
		if err != nil {
			log.Error("Upgrade system contract update error", "version", version, "name", contract.GetName(), "err", err)
			return
		}

		log.Info("system contract upgrade execution", "version", version, "name", contract.GetName(), "height", header.Number, "chainId", config.GetChainId(header.Number).String())

		err = contract.Execute(state, header, chainContext, config)
		if err != nil {
			log.Error("Upgrade system contract execute error", "version", version, "name", contract.GetName(), "err", err)
			return
		}
	}

	return
}
