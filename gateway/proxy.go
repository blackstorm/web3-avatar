package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Proxy struct {
	client   *ethclient.Client
	contract *Web3Avatar
}

func NewProxy(contract string, infuraToken string) (*Proxy, error) {
	cli, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/" + infuraToken)

	if err != nil {
		return nil, err
	}

	instance, err := NewWeb3Avatar(common.HexToAddress(contract), cli)

	if err != nil {
		return nil, err
	}

	return &Proxy{
		client:   cli,
		contract: instance,
	}, nil
}

func (w Proxy) GetAvatar(address string, name string) (string, error) {
	return w.contract.GetAvatar0(nil, common.HexToAddress(address), name)
}
