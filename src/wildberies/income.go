package wildberies

import (
	"context"
	"net/http"
)

type (
	// IncomeService is a service to deal with incomes.
	IncomeService struct {
		client *Client
	}

	// IncomeResp describes income method response.
	SellerInfoResp struct {
		Name      string `json:"name"`
		Sid       string `json:"sid"`
		TradeMark string `json:"tradeMark"`
	}
)

// List returns list of existed incomes.
func (s *IncomeService) SellerInfo(
	ctx context.Context,
) (*SellerInfoResp, *http.Response, error) {
	uri := "https://common-api.wildberries.ru/api/v1/seller-info"
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	sellerInfo := new(SellerInfoResp)
	resp, err := s.client.Do(ctx, req, sellerInfo)
	return sellerInfo, resp, err
}
