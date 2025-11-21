package common

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var CommonClient *ethclient.Client
var once sync.Mutex

type EthClient struct {
	chainID *big.Int
	opt     *bind.TransactOpts
	CallOpt *bind.CallOpts
}

type NewContractFunction[T any] func(a common.Address, b bind.ContractBackend) (T, error)

/**
* init eth-client
**/
func Init() {
	once.Lock()
	defer once.Unlock()
	if CommonClient == nil {
		c, err := ethclient.Dial(os.Getenv("SEPOLIA_RPC_URL"))
		if err != nil {
			fmt.Println("eth client error " + err.Error())
			return
		}
		CommonClient = c
	}
}

/**
* get chainId
**/
func (c *EthClient) ChainId() *EthClient {
	if c.chainID == nil {
		chainId, err := CommonClient.ChainID(context.Background())
		fmt.Println(chainId)
		if err != nil {
			fmt.Println("eth get chain id error " + err.Error())
			return nil
		}
		c.chainID = chainId
	}

	return c
}

/**
* get chain opt
**/
func (c *EthClient) Option() *EthClient {
	if c.opt == nil {
		privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x"))

		if err != nil {
			fmt.Println("key error " + err.Error())
			return nil
		}
		//for write state
		opts, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
		opts.GasLimit = uint64(300000)
		opts.Value = big.NewInt(0)
		if err != nil {
			fmt.Println("transaction error " + err.Error())
			return nil
		}
		c.opt = opts
	}
	return c
}

func (c *EthClient) CallOption() *EthClient {
	if c.CallOpt == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		//for pure or read
		c.CallOpt = &bind.CallOpts{
			Pending: true,
			Context: ctx,
		}
	}

	return c
}


func NewContract[T any](method NewContractFunction[T]) (T, error) {
	// bidAddr:=common.HexToAddress("0xDA20a8F646e3288C6505300074BB3787705b9E1f")
	bidAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	return method(bidAddr, CommonClient)
}
