package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"math/rand"
	"os"
)

//ViewResult generic
type ViewResult struct {
	Result []byte `json:"result"`
}

// GetNodesResult data
type GetNodesResult struct {
	Address      string `json:"address"`
	StakedAmount int64  `json:"staked_amount"`
	LockedAmount int    `json:"locked_amount"`
}

func getNodeURL() (string, error) {

	node := ""
	keyPair, err := key.NewBase58KeyPair(os.Getenv("NEAR_PK"))
	if err != nil {
		return node, fmt.Errorf("key error: %w", err)
	}

	network, ok := config.Networks["mainnet"]
	if !ok {
		return node, fmt.Errorf("unknown network '%s'", "mainnet")
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return node, fmt.Errorf("failed to create rpc client: %w", err)
	}

	ctx := client.ContextWithKeyPair(context.Background(), keyPair)

	res, err := rpc.ContractViewCallFunction(ctx, "webrtc.dtelecom.near", "get_nodes", base64.StdEncoding.EncodeToString([]byte("")), block.FinalityFinal())
	if err != nil {
		return node, fmt.Errorf("failed to view get_nodes: %w", err)
	}

	var viewResult ViewResult

	json.Unmarshal([]byte(res.Result), &viewResult)

	var getNodesResult []GetNodesResult
	json.Unmarshal(viewResult.Result, &getNodesResult)

	randomIndex := rand.Intn(len(getNodesResult))
	return getNodesResult[randomIndex].Address, nil
}
