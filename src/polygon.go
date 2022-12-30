package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"main/dmeet"
	"math/big"
	"os"
	"strconv"
)

func getMembershipOwner(id string) (string, error) {

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(os.Getenv("POLYGON_RPC"))
	if err != nil {
		return "", err
	}

	address := common.HexToAddress(os.Getenv("POLYGON_ADDRESS"))
	instance, err := dmeet.NewDmeet(address, client)
	if err != nil {
		return "", err
	}

	addr, err := instance.Creators(nil, big.NewInt(i))
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

func getMembershipBalance(account string, id string) (string, error) {

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}

	client, err := ethclient.Dial(os.Getenv("POLYGON_RPC"))
	if err != nil {
		return "", err
	}

	address := common.HexToAddress(os.Getenv("POLYGON_ADDRESS"))
	instance, err := dmeet.NewDmeet(address, client)
	if err != nil {
		return "", err
	}

	balance, err := instance.BalanceOf(nil, common.HexToAddress(account), big.NewInt(i))
	if err != nil {
		return "", err
	}

	return balance.String(), nil
}
