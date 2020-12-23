/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package adchain

import (
	"encoding/hex"
	"github.com/astaxie/beego/config"
	"github.com/eoscanada/eos-go"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/siddontang/go/log"
)

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager(nil)
	//读取配置
	absFile := filepath.Join("conf", "AD.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	wm.Api.Debug = true
	return wm
}

func TestWalletManager_GetInfo(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetInfo()
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetAccount(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetAccount("ad.token")
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetCurrencyBalance(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetCurrencyBalance("adnnzqln3yws", "AD", "ad.token")
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestGetTokenBalanceByAddress(t *testing.T) {
	wm := testNewWalletManager()

	contract := openwallet.SmartContract{
		Address:  "ad.token:AD",
		Symbol:   "AD",
		Name:     "AD",
		Decimals: 4,
	}
	balance, err := wm.ContractDecoder.GetTokenBalanceByAddress(contract, "testabcadcha")
	if err != nil {
		log.Error("GetTokenBalanceByAddress failed, unexpected error:", err)
		return
	}

	log.Info("balance:", balance)

}

func TestABIBinToJson(t *testing.T) {
	wm := testNewWalletManager()
	data, _ := hex.DecodeString("6072a6fed093b1ca6072a65e2193b1cad2040000000000000459535200000000010474657374")
	result, err := wm.Api.ABIBinToJSON(eos.AccountName("ysr.ystar"), eos.Name("yrctransfer"), eos.HexBytes(data))
	if err != nil {
		log.Error("ABIBinToJSON failed, unexpected error:", err)
		return
	}
	log.Infof("result: %+v", result)
}

func TestWalletManager_GetTransaction(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetTransaction("73fe7e4a93eecf0c8ae16a4a97b41f11efcc68c3c1b804e176ff9fcd310c2b42")
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)

}