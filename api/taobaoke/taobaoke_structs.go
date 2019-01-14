// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

const VersionNo = "20181116"

/* 淘宝客商品 */
type NTbkItem struct {
	ClickUrl        string   `json:"click_url"`
	ItemUrl         string   `json:"item_url"`
	Nick            string   `json:"nick"`
	NumIid          int      `json:"num_iid"`
	PictUrl         string   `json:"pict_url"`
	Provcity        string   `json:"provcity"`
	ReservePrice    string   `json:"reserve_price"`
	SellerId        int      `json:"seller_id"`
	ShopTitle       string   `json:"shop_title"`
	SmallImages     []string `json:"small_images"`
	Title           string   `json:"title"`
	TkRate          string   `json:"tk_rate"`
	UserType        int      `json:"user_type"`
	Volume          int      `json:"volume"`
	ZkFinalPrice    string   `json:"zk_final_price"`
	ZkFinalPriceWap string   `json:"zk_final_price_wap"`
}

/* 淘宝客选品库 */
type TbkFavorites struct {
	FavoritesId    int    `json:"favorites_id"`
	FavoritesTitle string `json:"favorites_title"`
	Type           int    `json:"type"`
}

/* data */
type Data struct {
	HasNext  bool       `json:"has_next"`
	PageNo   int        `json:"page_no"`
	PageSize int        `json:"page_size"`
	Results  []*MapData `json:"results"`
}

/* data */
type MapData struct {
	Model string `json:"model"`
}

/* 淘宝客店铺 */
type NTbkShop struct {
	ClickUrl   string `json:"click_url"`
	PictUrl    string `json:"pict_url"`
	SellerNick string `json:"seller_nick"`
	ShopTitle  string `json:"shop_title"`
	ShopType   string `json:"shop_type"`
	ShopUrl    string `json:"shop_url"`
	UserId     int    `json:"user_id"`
}

/* 淘宝联盟选品和招商宝贝信息 */
type UatmTbkItem struct {
	Category          int      `json:"category"`
	ClickUrl          string   `json:"click_url"`
	CommissionRate    string   `json:"commission_rate"`
	CouponClickUrl    string   `json:"coupon_click_url"`
	CouponEndTime     string   `json:"coupon_end_time"`
	CouponInfo        string   `json:"coupon_info"`
	CouponRemainCount int      `json:"coupon_remain_count"`
	CouponStartTime   string   `json:"coupon_start_time"`
	CouponTotalCount  int      `json:"coupon_total_count"`
	EventEndTime      string   `json:"event_end_time"`
	EventStartTime    string   `json:"event_start_time"`
	ItemUrl           string   `json:"item_url"`
	Nick              string   `json:"nick"`
	NumIid            int      `json:"num_iid"`
	PictUrl           string   `json:"pict_url"`
	Provcity          string   `json:"provcity"`
	ReservePrice      string   `json:"reserve_price"`
	SellerId          int      `json:"seller_id"`
	ShopTitle         string   `json:"shop_title"`
	SmallImages       []string `json:"small_images"`
	Status            int      `json:"status"`
	Title             string   `json:"title"`
	TkRate            string   `json:"tk_rate"`
	Type              int      `json:"type"`
	UserType          int      `json:"user_type"`
	Volume            int      `json:"volume"`
	ZkFinalPrice      string   `json:"zk_final_price"`
	ZkFinalPriceWap   string   `json:"zk_final_price_wap"`
}

/* TbkCoupon */
type TbkCoupon struct {
	Category          int      `json:"category"`
	CommissionRate    string   `json:"commission_rate"`
	CouponClickUrl    string   `json:"coupon_click_url"`
	CouponEndTime     string   `json:"coupon_end_time"`
	CouponInfo        string   `json:"coupon_info"`
	CouponRemainCount int      `json:"coupon_remain_count"`
	CouponStartTime   string   `json:"coupon_start_time"`
	CouponTotalCount  int      `json:"coupon_total_count"`
	ItemDescription   string   `json:"item_description"`
	ItemUrl           string   `json:"item_url"`
	Nick              string   `json:"nick"`
	NumIid            int      `json:"num_iid"`
	PictUrl           string   `json:"pict_url"`
	SellerId          int      `json:"seller_id"`
	ShopTitle         string   `json:"shop_title"`
	SmallImages       []string `json:"small_images"`
	Title             string   `json:"title"`
	UserType          int      `json:"user_type"`
	Volume            int      `json:"volume"`
	ZkFinalPrice      string   `json:"zk_final_price"`
}

/* 淘抢购对象 */
type Results struct {
	CategoryName string `json:"category_name"`
	ClickUrl     string `json:"click_url"`
	EndTime      string `json:"end_time"`
	NumIid       int    `json:"num_iid"`
	PicUrl       string `json:"pic_url"`
	ReservePrice string `json:"reserve_price"`
	SoldNum      int    `json:"sold_num"`
	StartTime    string `json:"start_time"`
	Title        string `json:"title"`
	TotalAmount  int    `json:"total_amount"`
	ZkFinalPrice string `json:"zk_final_price"`
}

/* 淘客定向招商活动基本信息 */
type TbkEvent struct {
	EndTime    string `json:"end_time"`
	EventId    int    `json:"event_id"`
	EventTitle string `json:"event_title"`
	StartTime  string `json:"start_time"`
}
