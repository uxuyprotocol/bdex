package bitget

import (
	"context"
	"errors"
	"strings"
)

// 代币信息
const (
	CoinPath        = "/market/v1/coin"
	CoinListPath    = "/market/v1/coinList"
	CoinHistoryPath = "/market/v1/historical-coins"
)

// "coin": "dogehq", //代币符号
// "name": "Dept Of Gov Efficiency HQ", //代币名称
// "chain": "sol", // 链简写
// "chainName": "Solana", // 链name
// "chainIcon": "https://cdn.bitkeep.vip/u_b_bbb3aa00-9afd-11ec-aac8-
// bf8a172584ab.png",
// "contract": "CFbEmC3JJ5HqXwfFNrKMzhAnFaMp64QkbBegcTw3Hpzh", // 代币合约
// 地址
// "price": "0.004907", // 代币价值（U）
// "percent": "554.07%", // 涨幅
// "unit": "$", // 价格单位
// "icon": "https://staticweb.jjdsn.vip/cqc/sol/174585c348f9f71c2b20dde748201243",
// "about": "",
// "totalSupply": "", // 总供应量
// "secureInfo": {
// "riskLevel": "low",
// "riskCount": 0,
// "contractCheckPic": "https://cdn.bitkeep.vip/u_b_83c5af10-cdb8-
// 11ec-984f-85fca8e9d3b6.png",
// "contractCheckUrl":
// "https://tools.bknode.vip/en/tools/ContractDetectionv3?
// chainId=100278&contractAddress=CFbEmC3JJ5HqXwfFNrKMzhAnFaMp64QkbBegcTw3Hpzh"
// },
// "source": { // 媒体信息
// "twitter": "https://x.com/search?q=$dogehq",
// "telegram": "https://t.me/DogeHQSol"
// },
// "shareUrl":
// "https://web3.bitget.com/en/swap/sol/CFbEmC3JJ5HqXwfFNrKMzhAnFaMp64QkbBegcTw3Hp
// zh",
// "chainRate": "1 SOL=44884.6 DOGEHQ",
// "tabVersion": "1.0",
// "issueDate": "2024-11-16 07:44:39", // 代币的创建时间（⾏情库数据的创建时间,
// 不⼀定是链上时间）
// "token_security": {
// "riskLevel": "low",
// "riskCount": 0,
// "contractCheckPic": "https://cdn.bitkeep.vip/u_b_83c5af10-cdb8-
// 11ec-984f-85fca8e9d3b6.png",
// "contractCheckUrl":
// "https://tools.bknode.vip/en/tools/ContractDetectionv3/100278/CFbEmC3JJ5HqXwfFN
// rKMzhAnFaMp64QkbBegcTw3Hpzh"
// },
// "fdv": 49069.857380361005

type SecureInfo struct {
	RiskLevel        string `json:"riskLevel"`
	RiskCount        int    `json:"riskCount"`
	ContractCheckPic string `json:"contractCheckPic"`
	ContractCheckUrl string `json:"contractCheckUrl"`
}

type Source struct {
	Twitter  string `json:"twitter,omitempty"`
	Telegram string `json:"telegram,omitempty"`
}

type TokenSecurity struct {
	RiskLevel        string `json:"riskLevel"`
	RiskCount        int    `json:"riskCount"`
	ContractCheckPic string `json:"contractCheckPic"`
	ContractCheckUrl string `json:"contractCheckUrl"`
}

type Coin struct {
	Coin          string        `json:"coin"`           // 代币符号
	Name          string        `json:"name"`           // 代币名称
	Chain         string        `json:"chain"`          // 链简写
	ChainName     string        `json:"chainName"`      // 链name
	ChainIcon     string        `json:"chainIcon"`      // 链图标
	Contract      string        `json:"contract"`       // 代币合约
	Price         string        `json:"price"`          // 代币价值（U）
	Percent       string        `json:"percent"`        // 涨幅
	Unit          string        `json:"unit"`           // 价格单位
	Icon          string        `json:"icon"`           // 代币图标
	About         string        `json:"about"`          // 关于代币
	TotalSupply   string        `json:"totalSupply"`    // 总供应量
	SecureInfo    SecureInfo    `json:"secureInfo"`     // 安全信息
	Source        Source        `json:"source"`         // 媒体信息
	ShareUrl      string        `json:"shareUrl"`       // 分享链接
	ChainRate     string        `json:"chainRate"`      // 链汇率
	TabVersion    string        `json:"tabVersion"`     // 版本
	IssueDate     string        `json:"issueDate"`      // 代币的创建时间
	TokenSecurity TokenSecurity `json:"token_security"` // 代币安全信息
	FDV           float64       `json:"fdv"`            // 完全稀释估值
}

func (c *BClient) Coin(ctx context.Context, chain, contract string) (*Coin, error) {
	req := map[string]interface{}{
		"chain":    chain,
		"contract": contract,
	}

	body, err := c.sendGetRequest(ctx, CoinPath, req)
	if err != nil {
		return nil, err
	}
	type CoinResponse struct {
		Status int    `json:"status"`
		Errmsg string `json:"errmsg"`
		Data   Coin   `json:"data"`
	}

	var coinResp CoinResponse
	if err = json.Unmarshal(body, &coinResp); err != nil {
		return nil, err
	}

	return &coinResp.Data, nil
}

// 获取多个代币基础信息
// /bgw-pro/market/v1/coinList
// param chains
// param contracts
// {
// "chain": "sol", // 链名称简写
// "contract": "7wti9XBn5L3gV815sovZUx7UDFxQydcSgCeBTXfHpump",
// "icon": "https://staticweb.jjdsn.vip/cqc/sol/89c2085ccf180f2ab957a91d812fbaa2",
// "name": "DOG IN POOL", //代币全称
// "symbol": "DIP", //币名称
// "price": "0.01032", // 价格
// "tvl": "111150.09", // 流动
// "fdv": "11110.32", // fdv 市值
// "vol24h": "11110.32", // 24 ⼩时交易量
// "change24h": "-0.5", // 24⼩时交易变化率
// "tokenSecurityStatus": 0 // 0 是安全 1是危险 2 告警 （有时效性， 5
// 分钟）
// "holders": 1233, // 持有⼈数
// "top10holderPercent": "0.1"
// }

type TokenSecurityStatus int

const (
	SecurityStatusSafe    TokenSecurityStatus = 0
	SecurityStatusDanger  TokenSecurityStatus = 1
	SecurityStatusWarning TokenSecurityStatus = 2
)

// CoinBase represents the base coin information structure
// as described in the API documentation comments
type CoinBase struct {
	Chain               string              `json:"chain"`
	Contract            string              `json:"contract"`
	Icon                string              `json:"icon"`
	Name                string              `json:"name"`
	Symbol              string              `json:"symbol"`
	Price               string              `json:"price"`
	TVL                 string              `json:"tvl"`
	FDV                 string              `json:"fdv"`
	Vol24h              string              `json:"vol24h"`
	Change24h           string              `json:"change24h"`
	TokenSecurityStatus TokenSecurityStatus `json:"tokenSecurityStatus"`
	Holders             int                 `json:"holders"`
	Top10HolderPercent  string              `json:"top10holderPercent"`
}

func (c *BClient) CoinList(ctx context.Context, chains, contracts []string) ([]CoinBase, error) {
	if len(chains) == 0 || len(contracts) > 50 {
		return nil, errors.New("chains and contracts  equal 0 or more than 50")
	}
	if len(chains) != len(contracts) {
		return nil, errors.New("chains and contracts length not equal")
	}

	req := map[string]string{
		"chains":    strings.Join(chains, ","),
		"contracts": strings.Join(contracts, ","),
	}
	reqByte, _ := json.Marshal(req)

	body, err := c.sendPostRequest(ctx, CoinListPath, reqByte)
	if err != nil {
		return nil, err
	}

	type CoinListResponse struct {
		Status int        `json:"status"`
		Errmsg string     `json:"errmsg"`
		Data   []CoinBase `json:"data"`
	}

	var coinListResp CoinListResponse
	if err = json.Unmarshal(body, &coinListResp); err != nil {
		return nil, err
	}

	return coinListResp.Data, nil
}

// 分⻋获取历史代币列表
// /bgw-pro/market/v1/historical-coins

// {
// "tokenList": [
// {
// "chain": "sol",
// "contract": "PhiLR4JDZB9z92rYT5xBXKCxmq4pGB1LYjtybii7aiS",
// "symbol": "POVT",
// "name": "philodmusic | Youtube",
// "decimals": 5,
// "icon":
// "https://dt0q2gt1jpyrq.cloudfront.net/cqc/sol/5a9f73da3057ce9bdfdfaa2ae6a3154d.
// png",
// "createTime": "2025-06-17 07:08:56"
// },
// {
// "chain": "sol",
// "contract": "E8bGWbEQvK7xkmdRVoZhMjFvPEwxyS5nKQMfppVtpump",
// "symbol": "VC",
// "name": "MAKE VC GREAT AGAIN",
// "decimals": 6,
// "icon":
// "https://dt0q2gt1jpyrq.cloudfront.net/cqc/sol/1157de64a64337a8e6910405a8da6762.
// jpg",
// "createTime": "2025-06-17 08:23:52"
// }
// ],
// "lastTime": "2025-06-17 08:23:52"
// },

// TokenBase represents the base token information structure
// as described in the historical coins API documentation comments
type TokenBase struct {
	Chain      string `json:"chain"`
	Contract   string `json:"contract"`
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Decimals   int    `json:"decimals"`
	Icon       string `json:"icon"`
	CreateTime string `json:"createTime"`
}

type TokenHistory struct {
	TokenList []TokenBase `json:"tokenList"`
	LastTime  string      `json:"lastTime"`
}

func (c *BClient) HistoricalCoins(ctx context.Context, createTime string, limit uint64) (*TokenHistory, error) {
	params := map[string]interface{}{
		"createTime": createTime,
		"limit":      limit,
	}

	body, err := c.sendGetRequest(ctx, CoinHistoryPath, params)
	if err != nil {
		return nil, err
	}

	type HistoricalCoinsResponse struct {
		Status int    `json:"status"`
		Errmsg string `json:"errmsg"`
		Data   TokenHistory
	}

	var coinHistoryResp HistoricalCoinsResponse
	if err = json.Unmarshal(body, &coinHistoryResp); err != nil {
		return nil, err
	}
	return &coinHistoryResp.Data, nil
}
