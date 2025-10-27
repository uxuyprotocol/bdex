package bitget

import "context"

// 批量获取代币安全检测
// /bgw-pro/market/v1/coin/security/audits
const (
	CoinSecurityAuditsPath = "/bgw-pro/market/v1/coin/security/audits"
)

type Token struct {
	ChainId  int64  `json:"chain_id"`
	Chain    string `json:"chain"`
	Contract string `json:"contract"`
}

type AuditsRequest struct {
	List   []Token `json:"list"`
	Source string  `json:"source"`
}

// Audit represents the token security audit data structure
type Audit struct {
	Chain                string      `json:"chain"`
	ChainID              int64       `json:"chain_id"`
	Contract             string      `json:"contract"`
	RiskCount            int         `json:"riskCount"`
	WarnCount            int         `json:"warnCount"`
	CheckStatus          int         `json:"checkStatus"`
	Checking             bool        `json:"checking"`
	Support              int         `json:"support"`
	HighRisk             bool        `json:"highRisk"`
	BuyTax               int         `json:"buyTax"`
	SellTax              int         `json:"sellTax"`
	FreezeAuth           bool        `json:"freezeAuth"`
	MintAuth             bool        `json:"mintAuth"`
	Token2022            bool        `json:"token2022"`
	LpLock               bool        `json:"lpLock"`
	Top10HolderRiskLevel int         `json:"top_10_holder_risk_level"`
	RiskChecks           []LabelPair `json:"riskChecks"`
	WarnChecks           []LabelPair `json:"warnChecks"`
	LowChecks            []LabelPair `json:"lowChecks"`
}

type LabelPair struct {
	LabelName string `json:"labelName"`
	Status    int    `json:"status"`
}

func (c *BClient) CoinSecurityAudits(ctx context.Context, request AuditsRequest) ([]Audit, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := c.sendPostRequest(ctx, CoinSecurityAuditsPath, jsonBody)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Status string  `json:"status"`
		Data   []Audit `json:"data"`
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
