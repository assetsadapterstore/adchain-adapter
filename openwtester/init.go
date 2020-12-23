package openwtester

import (
	"github.com/assetsadapterstore/adchain-adapter/adchain"
	"github.com/blocktree/eosio-adapter/eosio"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	cache := eosio.NewCacheManager()

	openw.RegAssets(adchain.Symbol, adchain.NewWalletManager(&cache))
}
