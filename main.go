package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"fmt"
	"math"
	"go-erc20-token-transaction/gethToken"
)

func main() {
	t := InitTranns("http://127.0.0.1", "0xdd974D5C2e2928deA5F71b9825b8b646686BD200")
	err := t.Transaction("input your address you want transc", "key file path", "your pwd", 0.0001)
	if err != nil {
		panic(err)
	}
}


type TokenTransaction struct {
	client  *ethclient.Client
	contractAddress string
}

func InitTranns(url,contractAddress string) *TokenTransaction {
	rpcDial, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}

	client := ethclient.NewClient(rpcDial)
	return &TokenTransaction{client:client,contractAddress:contractAddress}
}


func (s *TokenTransaction) Transaction(toAddress,keyfile,pwd string, tokenAmount float64) (err error){
	i,err := ioutil.ReadFile(keyfile)
	if err != nil {
		return
	}

	auth,err := bind.NewTransactor(strings.NewReader(string(i)), pwd)
	if err != nil {
		return
	}

	token,err := gethToken.NewToken(common.HexToAddress(s.contractAddress),s.client)
	if err != nil {
		return
	}


	amount := big.NewFloat(tokenAmount)
	tenDecimal := big.NewFloat(math.Pow(10, float64(18)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
	auth.GasLimit = 200000
	txs, err := token.Transfer(auth, common.HexToAddress(toAddress), convertAmount)
	if err != nil {
		return
	}

	fmt.Println("chainId---->", txs.ChainId())
	return
}