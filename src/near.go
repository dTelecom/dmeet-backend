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
	NodeID  string `json:"node_id"`
}

func getNodeURL() (string, string, error) {

	network, ok := config.Networks["mainnet"]
	if !ok {
		return "", "", fmt.Errorf("unknown network '%s'", "mainnet")
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to create rpc client: %w", err)
	}

	res, err := rpc.ContractViewCallFunction(context.Background(), "webrtc.dtelecom.near", "get_nodes", base64.StdEncoding.EncodeToString([]byte("")), block.FinalityFinal())
	if err != nil {
		return "", "", fmt.Errorf("failed to view get_nodes: %w", err)
	}

	var viewResult ViewResult
	err = json.Unmarshal([]byte(res.Result), &viewResult)
	if err != nil {
		return "", "", err
	}

	var getNodesResult []GetNodesResult
	err = json.Unmarshal(viewResult.Result, &getNodesResult)
	if err != nil {
		return "", "", err
	}

	randomIndex := rand.Intn(len(getNodesResult))
	return getNodesResult[randomIndex].Address, getNodesResult[randomIndex].NodeID, nil
}

func getTokenSignature(token *Token) (string, string, error) {

	keyPair, err := key.NewBase58KeyPair(os.Getenv("NEAR_PK"))
	if err != nil {
		return "", "", err
	}

	j, err := json.Marshal(token)
	if err != nil {
		return "", "", err
	}

	sig := keyPair.Sign(j)
	return base64.StdEncoding.EncodeToString(j), base64.StdEncoding.EncodeToString(sig.Value()), nil
}
