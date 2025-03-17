package wildberies

import (
	"context"
	"net/http"
)

type (
	SupplierService struct {
		client *Client
	}

	IncomesResp []struct {
		IncomeID        int    `json:"incomeId"`
		Number          string `json:"number"`
		Date            string `json:"date"`
		LastChangeDate  string `json:"lastChangeDate"`
		SupplierArticle string `json:"supplierArticle"`
		TechSize        string `json:"techSize"`
		Barcode         string `json:"barcode"`
		Quantity        int    `json:"quantity"`
		TotalPrice      int    `json:"totalPrice"`
		DateClose       string `json:"dateClose"`
		WarehouseName   string `json:"warehouseName"`
		NmID            int    `json:"nmId"`
		Status          string `json:"status"`
	}

	StocksResp []struct {
		LastChangeDate  string `json:"lastChangeDate"`
		WarehouseName   string `json:"warehouseName"`
		SupplierArticle string `json:"supplierArticle"`
		NmID            int    `json:"nmId"`
		Barcode         string `json:"barcode"`
		Quantity        int    `json:"quantity"`
		InWayToClient   int    `json:"inWayToClient"`
		InWayFromClient int    `json:"inWayFromClient"`
		QuantityFull    int    `json:"quantityFull"`
		Category        string `json:"category"`
		Subject         string `json:"subject"`
		Brand           string `json:"brand"`
		TechSize        string `json:"techSize"`
		Price           int    `json:"Price"`
		Discount        int    `json:"Discount"`
		IsSupply        bool   `json:"isSupply"`
		IsRealization   bool   `json:"isRealization"`
		SCCode          string `json:"SCCode"`
	}

	OrdersResp []struct {
		Date            string `json:"date"`
		LastChangeDate  string `json:"lastChangeDate"`
		WarehouseName   string `json:"warehouseName"`
		WarehouseType   string `json:"warehouseType"`
		CountryName     string `json:"countryName"`
		OblastOkrugName string `json:"oblastOkrugName"`
		RegionName      string `json:"regionName"`
		SupplierArticle string `json:"supplierArticle"`
		NmID            int    `json:"nmId"`
		Barcode         string `json:"barcode"`
		Category        string `json:"category"`
		Subject         string `json:"subject"`
		Brand           string `json:"brand"`
		TechSize        string `json:"techSize"`
		IncomeID        int    `json:"incomeID"`
		IsSupply        bool   `json:"isSupply"`
		IsRealization   bool   `json:"isRealization"`
		TotalPrice      int    `json:"totalPrice"`
		DiscountPercent int    `json:"discountPercent"`
		Spp             int    `json:"spp"`
		FinishedPrice   int    `json:"finishedPrice"`
		PriceWithDisc   int    `json:"priceWithDisc"`
		IsCancel        bool   `json:"isCancel"`
		CancelDate      string `json:"cancelDate"`
		OrderType       string `json:"orderType"`
		Sticker         string `json:"sticker"`
		GNumber         string `json:"gNumber"`
		Srid            string `json:"srid"`
	}

	SalesResp []struct {
		Date              string  `json:"date"`
		LastChangeDate    string  `json:"lastChangeDate"`
		WarehouseName     string  `json:"warehouseName"`
		WarehouseType     string  `json:"warehouseType"`
		CountryName       string  `json:"countryName"`
		OblastOkrugName   string  `json:"oblastOkrugName"`
		RegionName        string  `json:"regionName"`
		SupplierArticle   string  `json:"supplierArticle"`
		NmID              int     `json:"nmId"`
		Barcode           string  `json:"barcode"`
		Category          string  `json:"category"`
		Subject           string  `json:"subject"`
		Brand             string  `json:"brand"`
		TechSize          string  `json:"techSize"`
		IncomeID          int     `json:"incomeID"`
		IsSupply          bool    `json:"isSupply"`
		IsRealization     bool    `json:"isRealization"`
		TotalPrice        int     `json:"totalPrice"`
		DiscountPercent   int     `json:"discountPercent"`
		Spp               int     `json:"spp"`
		PaymentSaleAmount int     `json:"paymentSaleAmount"`
		ForPay            float64 `json:"forPay"`
		FinishedPrice     int     `json:"finishedPrice"`
		PriceWithDisc     int     `json:"priceWithDisc"`
		SaleID            string  `json:"saleID"`
		OrderType         string  `json:"orderType"`
		Sticker           string  `json:"sticker"`
		GNumber           string  `json:"gNumber"`
		Srid              string  `json:"srid"`
	}
)

func (s *SupplierService) Incomes(
	ctx context.Context,
) (*IncomesResp, *http.Response, error) {
	uri := "https://statistics-api.wildberries.ru/api/v1/supplier/incomes"
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	incomes := new(IncomesResp)
	resp, err := s.client.Do(ctx, req, incomes)
	return incomes, resp, err
}

func (s *SupplierService) Stocks(
	ctx context.Context,
) (*StocksResp, *http.Response, error) {
	uri := "https://statistics-api.wildberries.ru/api/v1/supplier/stocks"
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	stocks := new(StocksResp)
	resp, err := s.client.Do(ctx, req, stocks)
	return stocks, resp, err
}

func (s *SupplierService) Orders(
	ctx context.Context,
) (*OrdersResp, *http.Response, error) {
	uri := "https://statistics-api.wildberries.ru/api/v1/supplier/orders"
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	orders := new(OrdersResp)
	resp, err := s.client.Do(ctx, req, orders)
	return orders, resp, err
}

func (s *SupplierService) Sales(
	ctx context.Context,
) (*SalesResp, *http.Response, error) {
	uri := "https://statistics-api.wildberries.ru/api/v1/supplier/sales"
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	sales := new(SalesResp)
	resp, err := s.client.Do(ctx, req, sales)
	return sales, resp, err
}
