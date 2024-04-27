package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionRiskService) Symbol(symbol string) *GetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       decimal.Decimal  `json:"entryPrice"`
	BreakEvenPrice   string           `json:"breakEvenPrice"`
	MarginType       string           `json:"marginType"`
	IsAutoAddMargin  string           `json:"isAutoAddMargin"`
	IsolatedMargin   string           `json:"isolatedMargin"`
	Leverage         decimal.Decimal  `json:"leverage"`
	LiquidationPrice decimal.Decimal  `json:"liquidationPrice"`
	MarkPrice        string           `json:"markPrice"`
	MaxNotionalValue string           `json:"maxNotionalValue"`
	PositionAmt      decimal.Decimal  `json:"positionAmt"`
	Symbol           string           `json:"symbol"`
	UnRealizedProfit string           `json:"unRealizedProfit"`
	PositionSide     PositionSideType `json:"positionSide"`
	Notional         string           `json:"notional"`
	IsolatedWallet   string           `json:"isolatedWallet"`
}
