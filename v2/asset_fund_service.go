package binance

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
)

type GetFundingAssetService struct {
	c     *Client
	asset *string
}

func (s *GetFundingAssetService) Asset(asset string) *GetFundingAssetService {
	s.asset = &asset
	return s
}

// Do send request
func (s *GetFundingAssetService) Do(ctx context.Context, opts ...RequestOption) (assets []*FundingAsset, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setFormParam("asset", *s.asset)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	assets = make([]*FundingAsset, 0)
	err = json.Unmarshal(data, &assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

// FundingAsset define funding asset
type FundingAsset struct {
	Asset       string          `json:"asset"`
	Free        decimal.Decimal `json:"free"`
	Locked      decimal.Decimal `json:"locked"`
	Freeze      decimal.Decimal `json:"freeze"`
	Withdrawing decimal.Decimal `json:"withdrawing"`
}
