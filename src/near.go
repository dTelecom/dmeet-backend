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
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
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
	PK      string `json:"pk"`
}

func getNodeURL() (string, string, string, error) {

	network, ok := config.Networks["mainnet"]
	if !ok {
		return "", "", "", fmt.Errorf("unknown network '%s'", "mainnet")
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to create rpc client: %w", err)
	}

	res, err := rpc.ContractViewCallFunction(context.Background(), "webrtc.dtelecom.near", "get_nodes", base64.StdEncoding.EncodeToString([]byte("")), block.FinalityFinal())
	if err != nil {
		return "", "", "", fmt.Errorf("failed to view get_nodes: %w", err)
	}

	var viewResult ViewResult
	err = json.Unmarshal([]byte(res.Result), &viewResult)
	if err != nil {
		return "", "", "", err
	}

	var getNodesResult []GetNodesResult
	err = json.Unmarshal(viewResult.Result, &getNodesResult)
	if err != nil {
		return "", "", "", err
	}

	randomIndex := rand.Intn(len(getNodesResult))
	return getNodesResult[randomIndex].Address, getNodesResult[randomIndex].NodeID, getNodesResult[randomIndex].PK, nil
}

//GetTokenSignature token
func GetTokenSignature(token *Token) (string, string, error) {

	j, err := json.Marshal(token)
	if err != nil {
		return "", "", err
	}

	sig, err := GetSignature(j)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(j), base64.StdEncoding.EncodeToString(sig), nil
}

//GetSignature msg
func GetSignature(data []byte) ([]byte, error) {

	keyPair, err := key.NewBase58KeyPair(os.Getenv("NEAR_PK"))
	if err != nil {
		return []byte{}, err
	}

	sig := keyPair.Sign(data)
	return sig.Value(), nil
}

//VerifySignature msg
func VerifySignature(message []byte, pk string, sign []byte) (bool, error) {

	pubKey, err := key.NewBase58PublicKey(pk)
	if err != nil {
		return false, err
	}

	signature := signature.NewSignatureED25519(sign)

	publicKey := pubKey.ToPublicKey()
	ok, err := publicKey.Verify(message, signature)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func getEpochHeight() (uint64, error) {
	var height uint64 = 0

	network, ok := config.Networks["mainnet"]
	if !ok {
		return height, fmt.Errorf("unknown network '%s'", "mainnet")
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return height, fmt.Errorf("failed to create rpc client: %w", err)
	}

	res, err := rpc.ContractViewCallFunction(context.Background(), "webrtc.dtelecom.near", "get_epoch_height", base64.StdEncoding.EncodeToString([]byte("")), block.FinalityFinal())
	if err != nil {
		return height, fmt.Errorf("failed to view get_epoch_height: %w", err)
	}

	var viewResult ViewResult

	json.Unmarshal([]byte(res.Result), &viewResult)

	json.Unmarshal(viewResult.Result, &height)

	return height, nil
}
