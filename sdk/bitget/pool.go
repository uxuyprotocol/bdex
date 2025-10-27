package bitget

import (
	"context"
)

// /bgw-pro/market/poolList

const (
	PoolListPath = "/bgw-pro/market/poolList"
)

// PoolList represents the pool data structure
type PoolList struct {
	PoolAddr        string `json:"poolAddr"`
	PoolSymbol      string `json:"poolSymbol"`
	Protocol        string `json:"protocol"`
	ProtocolAddress string `json:"protocolAddress"`
	ProtocolIcon    string `json:"protocolIcon"`
	TotalUsd        string `json:"totalUsd"`
	Change          string `json:"change"`
	Token0Symbol    string `json:"token0Symbol"`
	Token1Symbol    string `json:"token1Symbol"`
	Token0Contract  string `json:"token0Contract"`
	Token1Contract  string `json:"token1Contract"`
	Reserve0        string `json:"reserve0"`
	Reserve1        string `json:"reserve1"`
	PriceRate       string `json:"priceRate"`
	PriceRateText   string `json:"priceRateText"`
	Token0Icon      string `json:"token0Icon"`
	Token1Icon      string `json:"token1Icon"`
	ActivityList    []struct {
		Side            string `json:"side"`
		Token0Symbol    string `json:"token0Symbol"`
		Token1Symbol    string `json:"token1Symbol"`
		Amount0         string `json:"amount0"`
		Amount1         string `json:"amount1"`
		Time            string `json:"time"`
		TxId            string `json:"txId"`
		TxUrl           string `json:"txUrl"`
		TransactionHash string `json:"transactionHash"`
		TransactionUrl  string `json:"transactionUrl"`
	} `json:"activityList"`
}

func (c *BClient) PoolList(ctx context.Context, chain, contract string, page, size int) ([]PoolList, error) {
	params := map[string]interface{}{
		"chain":    chain,
		"contract": contract,
		"page":     page,
		"size":     size,
	}

	body, err := c.sendPostRequest(ctx, PoolListPath, params)
	if err != nil {
		return nil, err
	}
	type PoolListResponse struct {
		Status string `json:"status"`
		Data   struct {
			List []PoolList `json:"list"`
		} `json:"data"`
	}
	var response PoolListResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response.Data.List, nil
}
