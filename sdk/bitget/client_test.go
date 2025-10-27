package bitget

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client *BClient
}

func (s *ClientSuite) SetupTest() {
	apiKey := os.Getenv("BITGET_API_KEY")
	apiSecret := os.Getenv("BITGET_API_SECRET")

	s.Require().NotEmpty(apiKey, "BITGET_API_KEY environment variable is required")
	s.Require().NotEmpty(apiSecret, "BITGET_API_SECRET environment variable is required")

	s.client = NewClient(apiKey, apiSecret, 5*time.Second)
}

// func (s *ClientSuite) TestKline() {
// 	ctx := context.Background()
// 	klines, err := s.client.Kline(ctx, "sol", "73UdJevxaNKXARgkvPHQGKuv8HCZARszuKW2LTL3pump", "1h", 10)
// 	s.NoError(err)
// 	s.NotNil(klines)
// }

// func (s *ClientSuite) TestCoinSecurityAudits() {
// 	ctx := context.Background()
// 	request := AuditsRequest{
// 		List: []Token{
// 			{
// 				ChainId:  1,
// 				Chain:    "ethereum",
// 				Contract: "0x5a98fcbea516cf06857215779fd812ca3bef1b32",
// 			},
// 		},
// 		Source: "bg",
// 	}

// 	audits, err := s.client.CoinSecurityAudits(ctx, request)
// 	s.NoError(err)
// 	s.NotNil(audits)
// }

func (s *ClientSuite) TestCoin() {
	ctx := context.Background()
	coins, err := s.client.Coin(ctx, "sol", "73UdJevxaNKXARgkvPHQGKuv8HCZARszuKW2LTL3pump")
	s.NoError(err)
	s.NotNil(coins)

}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
