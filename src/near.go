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
	Address string `json:"address"`
}

func getNodeURL() (string, error) {

	node := ""

	network, ok := config.Networks["mainnet"]
	if !ok {
		return node, fmt.Errorf("unknown network '%s'", "mainnet")
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return node, fmt.Errorf("failed to create rpc client: %w", err)
	}

	res, err := rpc.ContractViewCallFunction(context.Background(), "webrtc.dtelecom.near", "get_nodes", base64.StdEncoding.EncodeToString([]byte("")), block.FinalityFinal())
	if err != nil {
		return node, fmt.Errorf("failed to view get_nodes: %w", err)
	}

	var viewResult ViewResult
	err = json.Unmarshal([]byte(res.Result), &viewResult)
	if err != nil {
		return "", err
	}

	var getNodesResult []GetNodesResult
	err = json.Unmarshal(viewResult.Result, &getNodesResult)
	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(getNodesResult))
	return getNodesResult[randomIndex].Address, nil
}

func getTokenSignature(token *Token) (string, string, error) {

	keyPair, err := key.NewBase58KeyPair(os.Getenv("NEAR_PK"))
	if err != nil {
		return "", "", err
	}

	json, err := json.Marshal(token)
	if err != nil {
		return "", "", err
	}

	sig := keyPair.Sign(json)
	return base64.StdEncoding.EncodeToString(json), base64.StdEncoding.EncodeToString(sig.Value()), nil
}
