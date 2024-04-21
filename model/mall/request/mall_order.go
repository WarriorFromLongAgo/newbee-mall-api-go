package request

import "github.com/shopspring/decimal"

type PaySuccessParams struct {
	OrderNo string `json:"orderNo"`
	PayType int    `json:"payType"`
}

type OrderSearchParams struct {
	Status     string `form:"status"`
	PageNumber int    `form:"pageNumber"`
}

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds"`
	AddressId   int   `json:"addressId"`
}

type PayParams struct {
	OrderNo string `json:"orderNo"`
	//GoodsId       int    `json:"goodsId"`
	//GoodsName     string `json:"goodsName"`
	OriginalPrice decimal.Decimal `json:"originalPrice"`
	PayType       int             `json:"payType"`
}
