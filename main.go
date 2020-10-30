package main

import (
	"fmt"
	"github.com/xuyp1991/thresholdSign/ethutil"
	// "github.com/xuyp1991/thresholdSign/tss"
	"github.com/xuyp1991/thresholdSign/operator"
	"github.com/ethereum/go-ethereum/crypto"
	"encoding/hex"

)

const (
	strPwd = "12345678"
	testKeyFile1 = "/home/xuyapeng/ethereum/testnet/data/keystore/UTC--2020-10-28T06-37-48.082796475Z--e82bb272c6fac7f74a148dcc8a21f6fc46f3c761"

)

func main() {
	key1, err := ethutil.DecryptKeyFile(testKeyFile1,strPwd)
	if  err != nil {
		fmt.Println("DecryptKeyFile error ",err)
	}
	fmt.Println("DecryptKeyFile key succ ",key1)
	prikey, _ := operator.EthereumKeyToOperatorKey(key1)
	strkey := hex.EncodeToString(crypto.FromECDSA(prikey))
		fmt.Println("get pri key ",strkey)

}