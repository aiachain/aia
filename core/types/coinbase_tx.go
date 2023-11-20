package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

var (
	coinbaseRewardAmount = new(big.Int).Mul(big.NewInt(20), big.NewInt(params.Ether))
	CoinbaseAddr         = common.HexToAddress("0x000000000000000000000000000000000000A001")
)

func IsCoinbase(tx *Transaction) bool {
	if tx.Gas() == 0 && tx.GasPrice().Sign() == 0 {
		return true
	}
	return false
}

func CoinbaseReward(c *params.ChainConfig, num *big.Int) *big.Int {
	if c.IsTulip(num) {
		return new(big.Int).Mul(coinbaseRewardAmount, big.NewInt(10))
	}
	return coinbaseRewardAmount
}
