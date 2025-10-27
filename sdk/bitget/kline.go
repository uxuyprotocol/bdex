package bitget

import (
	"context"
)

// /bgw-pro/market/v2/kline

const (
	KlinePath = "/bgw-pro/market/v2/kline"
)

// Kline represents the kline data structure
// as described in the API documentation comments
type Kline struct {
	Timestamp  int64   `json:"ts"`         // 秒级时间戳，10位
	High       float64 `json:"high"`       // 最⾼价
	Low        float64 `json:"low"`        // 最低价
	Open       float64 `json:"open"`       // 开盘价
	Close      float64 `json:"close"`      // 收盘价
	Volume     float64 `json:"volume"`     // 交易额
	Amount     float64 `json:"amount"`     // 交易量
	Txn        int     `json:"txn"`        // 交易个数
	BuyVolume  float64 `json:"buyVolume"`  // 交易额-买
	SellVolume float64 `json:"sellVolume"` // 交易额-卖
	BuyAmount  float64 `json:"buyAmount"`  // 交易量-买
	SellAmount float64 `json:"sellAmount"` // 交易量-卖
}

func (c *BClient) Kline(ctx context.Context, chain, contract, period string, size int64) ([]Kline, error) {
	params := map[string]interface{}{
		"chain":    chain,
		"contract": contract,
		"period":   period,
		"size":     size,
	}

	body, err := c.sendPostRequest(ctx, KlinePath, params)
	if err != nil {
		return nil, err
	}

	type KlineResponse struct {
		Status string `json:"status"`
		Data   struct {
			List []Kline `json:"list"`
		} `json:"data"`
	}

	var response KlineResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response.Data.List, nil
}
