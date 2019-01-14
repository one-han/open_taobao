// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/one-han/open_taobao"
)

/* 淘宝客商品详情查询（简版） */
type TbkItemInfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供 */
func (r *TbkItemInfoGetRequest) SetIp(value string) {
	r.SetValue("ip", value)
}

/* 商品ID串，用,分割，最大40个 */
func (r *TbkItemInfoGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemInfoGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

func (r *TbkItemInfoGetRequest) GetResponse(accessToken string) (*TbkItemInfoGetResponse, []byte, error) {
	var resp TbkItemInfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.info.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemInfoGetResponse struct {
	Results []*NTbkItem `json:"results"`
}

type TbkItemInfoGetResponseResult struct {
	Response *TbkItemInfoGetResponse `json:"tbk_item_info_get_response"`
}

/* 淘宝客商品关联推荐查询 */
type TbkItemRecommendGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 返回数量，默认20，最大值40 */
func (r *TbkItemRecommendGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 需返回的字段列表 */
func (r *TbkItemRecommendGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品Id */
func (r *TbkItemRecommendGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemRecommendGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

func (r *TbkItemRecommendGetRequest) GetResponse(accessToken string) (*TbkItemRecommendGetResponse, []byte, error) {
	var resp TbkItemRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemRecommendGetResponse struct {
	Results []*NTbkItem `json:"results"`
}

type TbkItemRecommendGetResponseResult struct {
	Response *TbkItemRecommendGetResponse `json:"tbk_item_recommend_get_response"`
}

/* 淘宝客商品查询 */
type TbkItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (r *TbkItemGetRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 折扣价范围上限，单位：元 */
func (r *TbkItemGetRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 淘客佣金比率下限，如：1234表示12.34% */
func (r *TbkItemGetRequest) SetEndTkRate(value string) {
	r.SetValue("end_tk_rate", value)
}

/* 需返回的字段列表 */
func (r *TbkItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemGetRequest) SetIsOverseas(value string) {
	r.SetValue("is_overseas", value)
}

/* 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkItemGetRequest) SetIsTmall(value string) {
	r.SetValue("is_tmall", value)
}

/* 所在地 */
func (r *TbkItemGetRequest) SetItemloc(value string) {
	r.SetValue("itemloc", value)
}

/* 第几页，默认：１ */
func (r *TbkItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkItemGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi） */
func (r *TbkItemGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 折扣价范围下限，单位：元 */
func (r *TbkItemGetRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}

/* 淘客佣金比率上限，如：1234表示12.34% */
func (r *TbkItemGetRequest) SetStartTkRate(value string) {
	r.SetValue("start_tk_rate", value)
}

func (r *TbkItemGetRequest) GetResponse(accessToken string) (*TbkItemGetResponse, []byte, error) {
	var resp TbkItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkItemGetResponse struct {
	Results      []*NTbkItem `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkItemGetResponseResult struct {
	Response *TbkItemGetResponse `json:"tbk_item_get_response"`
}

/* 枚举出淘宝客在淘宝联盟超级搜索，特色频道中，创建的选品库列表 */
type TbkUatmFavoritesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段列表，不能为空，字段名之间使用逗号分隔 */
func (r *TbkUatmFavoritesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，从1开始计数 */
func (r *TbkUatmFavoritesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 默认20，页大小，即每一页的活动个数 */
func (r *TbkUatmFavoritesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 默认值-1；选品库类型，1：普通选品组，2：高佣选品组，-1，同时输出所有类型的选品组 */
func (r *TbkUatmFavoritesGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

func (r *TbkUatmFavoritesGetRequest) GetResponse(accessToken string) (*TbkUatmFavoritesGetResponse, []byte, error) {
	var resp TbkUatmFavoritesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.favorites.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmFavoritesGetResponse struct {
	Results      []*TbkFavorites `json:"results"`
	TotalResults int             `json:"total_results"`
}

type TbkUatmFavoritesGetResponseResult struct {
	Response *TbkUatmFavoritesGetResponse `json:"tbk_uatm_favorites_get_response"`
}

/* 拉新活动汇总API */
type TbkScNewuserOrderSumRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 活动ID，现有活动id包括： 2月手淘拉新：119013_2；3月手淘拉新：119013_3； */
func (r *TbkScNewuserOrderSumRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkScNewuserOrderSumRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 页码，默认1 */
func (r *TbkScNewuserOrderSumRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkScNewuserOrderSumRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 结算月份，需按照活动的结算月份传入具体的值：201807 */
func (r *TbkScNewuserOrderSumRequest) SetSettleMonth(value string) {
	r.SetValue("settle_month", value)
}

/* mm_xxx_xxx_xxx的第二位 */
func (r *TbkScNewuserOrderSumRequest) SetSiteId(value string) {
	r.SetValue("site_id", value)
}

func (r *TbkScNewuserOrderSumRequest) GetResponse(accessToken string) (*TbkScNewuserOrderSumResponse, []byte, error) {
	var resp TbkScNewuserOrderSumResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.sc.newuser.order.sum", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkScNewuserOrderSumResponse struct {
	Results *Data `json:"results"`
}

type TbkScNewuserOrderSumResponseResult struct {
	Response *TbkScNewuserOrderSumResponse `json:"tbk_sc_newuser_order_sum_response"`
}

/* 淘宝客店铺关联推荐查询 */
type TbkShopRecommendGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 返回数量，默认20，最大值40 */
func (r *TbkShopRecommendGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 需返回的字段列表 */
func (r *TbkShopRecommendGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkShopRecommendGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 卖家Id */
func (r *TbkShopRecommendGetRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}

func (r *TbkShopRecommendGetRequest) GetResponse(accessToken string) (*TbkShopRecommendGetResponse, []byte, error) {
	var resp TbkShopRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shop.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopRecommendGetResponse struct {
	Results []*NTbkShop `json:"results"`
}

type TbkShopRecommendGetResponseResult struct {
	Response *TbkShopRecommendGetResponse `json:"tbk_shop_recommend_get_response"`
}

/* 拉新活动汇总API */
type TbkDgNewuserOrderSumRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277 */
func (r *TbkDgNewuserOrderSumRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgNewuserOrderSumRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 页码，默认1 */
func (r *TbkDgNewuserOrderSumRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgNewuserOrderSumRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 结算月份 */
func (r *TbkDgNewuserOrderSumRequest) SetSettleMonth(value string) {
	r.SetValue("settle_month", value)
}

/* mm_xxx_xxx_xxx的第二位 */
func (r *TbkDgNewuserOrderSumRequest) SetSiteId(value string) {
	r.SetValue("site_id", value)
}

func (r *TbkDgNewuserOrderSumRequest) GetResponse(accessToken string) (*TbkDgNewuserOrderSumResponse, []byte, error) {
	var resp TbkDgNewuserOrderSumResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.newuser.order.sum", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgNewuserOrderSumResponse struct {
	Results *Data `json:"results"`
}

type TbkDgNewuserOrderSumResponseResult struct {
	Response *TbkDgNewuserOrderSumResponse `json:"tbk_dg_newuser_order_sum_response"`
}

/* 指定选品库id，获取该选品库的宝贝信息 */
type TbkUatmFavoritesItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 推广位id，需要在淘宝联盟后台创建；且属于appkey备案的媒体id（siteid），如何获取adzoneid，请参考，http://club.alimama.com/read-htm-tid-6333967.html?spm=0.0.0.0.msZnx5 */
func (r *TbkUatmFavoritesItemGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 选品库的id */
func (r *TbkUatmFavoritesItemGetRequest) SetFavoritesId(value string) {
	r.SetValue("favorites_id", value)
}

/* 需要输出则字段列表，逗号分隔 */
func (r *TbkUatmFavoritesItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认：1，从1开始计数 */
func (r *TbkUatmFavoritesItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkUatmFavoritesItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkUatmFavoritesItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道 */
func (r *TbkUatmFavoritesItemGetRequest) SetUnid(value string) {
	r.SetValue("unid", value)
}

func (r *TbkUatmFavoritesItemGetRequest) GetResponse(accessToken string) (*TbkUatmFavoritesItemGetResponse, []byte, error) {
	var resp TbkUatmFavoritesItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.favorites.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmFavoritesItemGetResponse struct {
	Results      []*UatmTbkItem `json:"results"`
	TotalResults int            `json:"total_results"`
}

type TbkUatmFavoritesItemGetResponseResult struct {
	Response *TbkUatmFavoritesItemGetResponse `json:"tbk_uatm_favorites_item_get_response"`
}

/* 通过指定定向招商活动id，获取该活动id下的宝贝信息；
宝贝信息中包含了适用于定向招商活动的高佣金淘客点击串; 注意，只能获取已经开始的定向招商id下面的宝贝信息，当天新开始的定向招商活动在凌晨2点生效； */
type TbkUatmEventItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 推广位id，需要在淘宝联盟后台创建；且属于appkey对应的备案媒体id（siteid），如何获取adzoneid，请参考：http://club.alimama.com/read-htm-tid-6333967.html?spm=0.0.0.0.msZnx5 */
func (r *TbkUatmEventItemGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 招商活动id */
func (r *TbkUatmEventItemGetRequest) SetEventId(value string) {
	r.SetValue("event_id", value)
}

/* 需要输出则字段列表，逗号分隔 */
func (r *TbkUatmEventItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认：１，从1开始计数 */
func (r *TbkUatmEventItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkUatmEventItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkUatmEventItemGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道 */
func (r *TbkUatmEventItemGetRequest) SetUnid(value string) {
	r.SetValue("unid", value)
}

func (r *TbkUatmEventItemGetRequest) GetResponse(accessToken string) (*TbkUatmEventItemGetResponse, []byte, error) {
	var resp TbkUatmEventItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.event.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmEventItemGetResponse struct {
	Results      []*UatmTbkItem `json:"results"`
	TotalResults int            `json:"total_results"`
}

type TbkUatmEventItemGetResponseResult struct {
	Response *TbkUatmEventItemGetResponse `json:"tbk_uatm_event_item_get_response"`
}

/* 淘宝客店铺查询 */
type TbkShopGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 累计推广商品上限 */
func (r *TbkShopGetRequest) SetEndAuctionCount(value string) {
	r.SetValue("end_auction_count", value)
}

/* 淘客佣金比率上限，1~10000 */
func (r *TbkShopGetRequest) SetEndCommissionRate(value string) {
	r.SetValue("end_commission_rate", value)
}

/* 信用等级上限，1~20 */
func (r *TbkShopGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 店铺商品总数上限 */
func (r *TbkShopGetRequest) SetEndTotalAction(value string) {
	r.SetValue("end_total_action", value)
}

/* 需返回的字段列表 */
func (r *TbkShopGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性 */
func (r *TbkShopGetRequest) SetIsTmall(value string) {
	r.SetValue("is_tmall", value)
}

/* 第几页，默认1，1~100 */
func (r *TbkShopGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkShopGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkShopGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkShopGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction） */
func (r *TbkShopGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 累计推广商品下限 */
func (r *TbkShopGetRequest) SetStartAuctionCount(value string) {
	r.SetValue("start_auction_count", value)
}

/* 淘客佣金比率下限，1~10000 */
func (r *TbkShopGetRequest) SetStartCommissionRate(value string) {
	r.SetValue("start_commission_rate", value)
}

/* 信用等级下限，1~20 */
func (r *TbkShopGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 店铺商品总数下限 */
func (r *TbkShopGetRequest) SetStartTotalAction(value string) {
	r.SetValue("start_total_action", value)
}

func (r *TbkShopGetRequest) GetResponse(accessToken string) (*TbkShopGetResponse, []byte, error) {
	var resp TbkShopGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.shop.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkShopGetResponse struct {
	Results      []*NTbkShop `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkShopGetResponseResult struct {
	Response *TbkShopGetResponse `json:"tbk_shop_get_response"`
}

/* 通用物料推荐，传入官方公布的物料id，可获取指定物料 */
type TbkDgOptimusMaterialRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgOptimusMaterialRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 内容详情ID */
func (r *TbkDgOptimusMaterialRequest) SetContentId(value string) {
	r.SetValue("content_id", value)
}

/* 内容渠道信息 */
func (r *TbkDgOptimusMaterialRequest) SetContentSource(value string) {
	r.SetValue("content_source", value)
}

/* 设备号加密类型：MD5 */
func (r *TbkDgOptimusMaterialRequest) SetDeviceEncrypt(value string) {
	r.SetValue("device_encrypt", value)
}

/* 设备号类型：IMEI，或者IDFA，或者UTDID */
func (r *TbkDgOptimusMaterialRequest) SetDeviceType(value string) {
	r.SetValue("device_type", value)
}

/* 设备号加密后的值 */
func (r *TbkDgOptimusMaterialRequest) SetDeviceValue(value string) {
	r.SetValue("device_value", value)
}

/* 官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096) */
func (r *TbkDgOptimusMaterialRequest) SetMaterialId(value string) {
	r.SetValue("material_id", value)
}

/* 第几页，默认：1 */
func (r *TbkDgOptimusMaterialRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgOptimusMaterialRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *TbkDgOptimusMaterialRequest) GetResponse(accessToken string) (*TbkDgOptimusMaterialResponse, []byte, error) {
	var resp TbkDgOptimusMaterialResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.optimus.material", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgOptimusMaterialResponse struct {
	ResultList []*MapData `json:"result_list"`
}

type TbkDgOptimusMaterialResponseResult struct {
	Response *TbkDgOptimusMaterialResponse `json:"tbk_dg_optimus_material_response"`
}

/* 好券清单API【导购】 */
type TbkDgItemCouponGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgItemCouponGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (r *TbkDgItemCouponGetRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 第几页，默认：1（当后台类目和查询词均不指定的时候，最多出10000个结果，即page_no*page_size不能超过10000；当指定类目或关键词的时候，则最多出100个结果） */
func (r *TbkDgItemCouponGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgItemCouponGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 1：PC，2：无线，默认：1 */
func (r *TbkDgItemCouponGetRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkDgItemCouponGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

func (r *TbkDgItemCouponGetRequest) GetResponse(accessToken string) (*TbkDgItemCouponGetResponse, []byte, error) {
	var resp TbkDgItemCouponGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.item.coupon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgItemCouponGetResponse struct {
	Results      []*TbkCoupon `json:"results"`
	TotalResults int          `json:"total_results"`
}

type TbkDgItemCouponGetResponseResult struct {
	Response *TbkDgItemCouponGetResponse `json:"tbk_dg_item_coupon_get_response"`
}

/* 拉新API */
type TbkDgNewuserOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277 */
func (r *TbkDgNewuserOrderGetRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgNewuserOrderGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间 */
func (r *TbkDgNewuserOrderGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码，默认1 */
func (r *TbkDgNewuserOrderGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgNewuserOrderGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间 */
func (r *TbkDgNewuserOrderGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TbkDgNewuserOrderGetRequest) GetResponse(accessToken string) (*TbkDgNewuserOrderGetResponse, []byte, error) {
	var resp TbkDgNewuserOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.newuser.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgNewuserOrderGetResponse struct {
	Results *Results `json:"results"`
}

type TbkDgNewuserOrderGetResponseResult struct {
	Response *TbkDgNewuserOrderGetResponse `json:"tbk_dg_newuser_order_get_response"`
}

/* 通用物料搜索API（导购） */
type TbkDgMaterialOptionalRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkDgMaterialOptionalRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (r *TbkDgMaterialOptionalRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 设备号加密类型：MD5 */
func (r *TbkDgMaterialOptionalRequest) SetDeviceEncrypt(value string) {
	r.SetValue("device_encrypt", value)
}

/* 设备号类型：IMEI，或者IDFA，或者UTDID */
func (r *TbkDgMaterialOptionalRequest) SetDeviceType(value string) {
	r.SetValue("device_type", value)
}

/* 设备号加密后的值 */
func (r *TbkDgMaterialOptionalRequest) SetDeviceValue(value string) {
	r.SetValue("device_value", value)
}

/* KA媒体淘客佣金比率上限，如：1234表示12.34% */
func (r *TbkDgMaterialOptionalRequest) SetEndKaTkRate(value string) {
	r.SetValue("end_ka_tk_rate", value)
}

/* 折扣价范围上限，单位：元 */
func (r *TbkDgMaterialOptionalRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 淘客佣金比率上限，如：1234表示12.34% */
func (r *TbkDgMaterialOptionalRequest) SetEndTkRate(value string) {
	r.SetValue("end_tk_rate", value)
}

/* 是否有优惠券，设置为true表示该商品有优惠券，设置为false或不设置表示不判断这个属性 */
func (r *TbkDgMaterialOptionalRequest) SetHasCoupon(value string) {
	r.SetValue("has_coupon", value)
}

/* 好评率是否高于行业均值 */
func (r *TbkDgMaterialOptionalRequest) SetIncludeGoodRate(value string) {
	r.SetValue("include_good_rate", value)
}

/* 成交转化是否高于行业均值 */
func (r *TbkDgMaterialOptionalRequest) SetIncludePayRate30(value string) {
	r.SetValue("include_pay_rate_30", value)
}

/* 退款率是否低于行业均值 */
func (r *TbkDgMaterialOptionalRequest) SetIncludeRfdRate(value string) {
	r.SetValue("include_rfd_rate", value)
}

/* ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供 */
func (r *TbkDgMaterialOptionalRequest) SetIp(value string) {
	r.SetValue("ip", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkDgMaterialOptionalRequest) SetIsOverseas(value string) {
	r.SetValue("is_overseas", value)
}

/* 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性 */
func (r *TbkDgMaterialOptionalRequest) SetIsTmall(value string) {
	r.SetValue("is_tmall", value)
}

/* 所在地 */
func (r *TbkDgMaterialOptionalRequest) SetItemloc(value string) {
	r.SetValue("itemloc", value)
}

/* 官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)，不传时默认为2836 */
func (r *TbkDgMaterialOptionalRequest) SetMaterialId(value string) {
	r.SetValue("material_id", value)
}

/* 是否包邮，true表示包邮，空或false表示不限 */
func (r *TbkDgMaterialOptionalRequest) SetNeedFreeShipment(value string) {
	r.SetValue("need_free_shipment", value)
}

/* 是否加入消费者保障，true表示加入，空或false表示不限 */
func (r *TbkDgMaterialOptionalRequest) SetNeedPrepay(value string) {
	r.SetValue("need_prepay", value)
}

/* 牛皮癣程度，取值：1:不限，2:无，3:轻微 */
func (r *TbkDgMaterialOptionalRequest) SetNpxLevel(value string) {
	r.SetValue("npx_level", value)
}

/* 第几页，默认：１ */
func (r *TbkDgMaterialOptionalRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkDgMaterialOptionalRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (r *TbkDgMaterialOptionalRequest) SetPlatform(value string) {
	r.SetValue("platform", value)
}

/* 查询词 */
func (r *TbkDgMaterialOptionalRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price） */
func (r *TbkDgMaterialOptionalRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 店铺dsr评分，筛选高于等于当前设置的店铺dsr评分的商品0-50000之间 */
func (r *TbkDgMaterialOptionalRequest) SetStartDsr(value string) {
	r.SetValue("start_dsr", value)
}

/* KA媒体淘客佣金比率下限，如：1234表示12.34% */
func (r *TbkDgMaterialOptionalRequest) SetStartKaTkRate(value string) {
	r.SetValue("start_ka_tk_rate", value)
}

/* 折扣价范围下限，单位：元 */
func (r *TbkDgMaterialOptionalRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}

/* 淘客佣金比率下限，如：1234表示12.34% */
func (r *TbkDgMaterialOptionalRequest) SetStartTkRate(value string) {
	r.SetValue("start_tk_rate", value)
}

func (r *TbkDgMaterialOptionalRequest) GetResponse(accessToken string) (*TbkDgMaterialOptionalResponse, []byte, error) {
	var resp TbkDgMaterialOptionalResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.dg.material.optional", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkDgMaterialOptionalResponse struct {
	ResultList   []*MapData `json:"result_list"`
	TotalResults int        `json:"total_results"`
}

type TbkDgMaterialOptionalResponseResult struct {
	Response *TbkDgMaterialOptionalResponse `json:"tbk_dg_material_optional_response"`
}

/* 枚举指定淘客自己发起的，*正在进行中的*定向招商的活动列表；每天新开始的定向招商活动，在凌晨2点后生效；即凌晨2点后可以枚举到当天开始的定向招商活动列表；时间过早不能取到当天开始的定向招商活动； */
type TbkUatmEventGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段列表，不能为空，字段名之间使用逗号分隔 */
func (r *TbkUatmEventGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 默认1，第几页，从1开始计数 */
func (r *TbkUatmEventGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 默认20,  页大小，即每一页的活动个数 */
func (r *TbkUatmEventGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *TbkUatmEventGetRequest) GetResponse(accessToken string) (*TbkUatmEventGetResponse, []byte, error) {
	var resp TbkUatmEventGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.uatm.event.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkUatmEventGetResponse struct {
	Results      []*TbkEvent `json:"results"`
	TotalResults int         `json:"total_results"`
}

type TbkUatmEventGetResponseResult struct {
	Response *TbkUatmEventGetResponse `json:"tbk_uatm_event_get_response"`
}

/* 通用物料推荐，传入官方公布的物料id，可获取指定物料 */
type TbkScOptimusMaterialRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkScOptimusMaterialRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 内容详情ID */
func (r *TbkScOptimusMaterialRequest) SetContentId(value string) {
	r.SetValue("content_id", value)
}

/* 内容渠道信息 */
func (r *TbkScOptimusMaterialRequest) SetContentSource(value string) {
	r.SetValue("content_source", value)
}

/* 设备号加密类型：MD5 */
func (r *TbkScOptimusMaterialRequest) SetDeviceEncrypt(value string) {
	r.SetValue("device_encrypt", value)
}

/* 设备号加密后的值 */
func (r *TbkScOptimusMaterialRequest) SetDeviceType(value string) {
	r.SetValue("device_type", value)
}

/* 设备号类型：IMEI，或者IDFA，或者UTDID */
func (r *TbkScOptimusMaterialRequest) SetDeviceValue(value string) {
	r.SetValue("device_value", value)
}

/* 官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096) */
func (r *TbkScOptimusMaterialRequest) SetMaterialId(value string) {
	r.SetValue("material_id", value)
}

/* 第几页，默认：1 */
func (r *TbkScOptimusMaterialRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkScOptimusMaterialRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* mm_xxx_xxx_xxx的第二位 */
func (r *TbkScOptimusMaterialRequest) SetSiteId(value string) {
	r.SetValue("site_id", value)
}

func (r *TbkScOptimusMaterialRequest) GetResponse(accessToken string) (*TbkScOptimusMaterialResponse, []byte, error) {
	var resp TbkScOptimusMaterialResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.sc.optimus.material", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkScOptimusMaterialResponse struct {
	ResultList []*MapData `json:"result_list"`
}

type TbkScOptimusMaterialResponseResult struct {
	Response *TbkScOptimusMaterialResponse `json:"tbk_sc_optimus_material_response"`
}

/* 阿里妈妈推广券信息查询。传入商品ID+券ID，或者传入me参数，均可查询券信息。 */
type TbkCouponGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 券ID */
func (r *TbkCouponGetRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}

/* 商品ID */
func (r *TbkCouponGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 带券ID与商品ID的加密串 */
func (r *TbkCouponGetRequest) SetMe(value string) {
	r.SetValue("me", value)
}

func (r *TbkCouponGetRequest) GetResponse(accessToken string) (*TbkCouponGetResponse, []byte, error) {
	var resp TbkCouponGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.coupon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkCouponGetResponse struct {
	Data *MapData `json:"data"`
}

type TbkCouponGetResponseResult struct {
	Response *TbkCouponGetResponse `json:"tbk_coupon_get_response"`
}

/* 拉新API */
type TbkScNewuserOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 活动id， 现有活动id包括： 2月手淘拉新：119013_2 3月手淘拉新：119013_3 4月手淘拉新：119013_4 拉手机支付宝新用户_赚赏金：200000_5 */
func (r *TbkScNewuserOrderGetRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}

/* mm_xxx_xxx_xxx的第三位 */
func (r *TbkScNewuserOrderGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间 */
func (r *TbkScNewuserOrderGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码，默认1 */
func (r *TbkScNewuserOrderGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认20，1~100 */
func (r *TbkScNewuserOrderGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* mm_xxx_xxx_xxx的第二位 */
func (r *TbkScNewuserOrderGetRequest) SetSiteId(value string) {
	r.SetValue("site_id", value)
}

/* 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间 */
func (r *TbkScNewuserOrderGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TbkScNewuserOrderGetRequest) GetResponse(accessToken string) (*TbkScNewuserOrderGetResponse, []byte, error) {
	var resp TbkScNewuserOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.sc.newuser.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkScNewuserOrderGetResponse struct {
	Results *Results `json:"results"`
}

type TbkScNewuserOrderGetResponseResult struct {
	Response *TbkScNewuserOrderGetResponse `json:"tbk_sc_newuser_order_get_response"`
}

/* 提供淘客生成淘口令接口，淘客提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播 */
type TbkTpwdCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 扩展字段JSON格式 */
func (r *TbkTpwdCreateRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 口令弹框logoURL */
func (r *TbkTpwdCreateRequest) SetLogo(value string) {
	r.SetValue("logo", value)
}

/* 口令弹框内容 */
func (r *TbkTpwdCreateRequest) SetText(value string) {
	r.SetValue("text", value)
}

/* 口令跳转目标页 */
func (r *TbkTpwdCreateRequest) SetUrl(value string) {
	r.SetValue("url", value)
}

/* 生成口令的淘宝用户ID */
func (r *TbkTpwdCreateRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}

func (r *TbkTpwdCreateRequest) GetResponse(accessToken string) (*TbkTpwdCreateResponse, []byte, error) {
	var resp TbkTpwdCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.tpwd.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkTpwdCreateResponse struct {
	Data *MapData `json:"data"`
}

type TbkTpwdCreateResponseResult struct {
	Response *TbkTpwdCreateResponse `json:"tbk_tpwd_create_response"`
}

/* 获取淘抢购的数据，淘客商品转淘客链接，非淘客商品输出普通链接 */
type TbkJuTqgGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 推广位id（推广位申请方式：http://club.alimama.com/read.php?spm=0.0.0.0.npQdST&tid=6306396&ds=1&page=1&toread=1） */
func (r *TbkJuTqgGetRequest) SetAdzoneId(value string) {
	r.SetValue("adzone_id", value)
}

/* 最晚开团时间 */
func (r *TbkJuTqgGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 需返回的字段列表 */
func (r *TbkJuTqgGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 第几页，默认1，1~100 */
func (r *TbkJuTqgGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页大小，默认40，1~40 */
func (r *TbkJuTqgGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 最早开团时间 */
func (r *TbkJuTqgGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TbkJuTqgGetRequest) GetResponse(accessToken string) (*TbkJuTqgGetResponse, []byte, error) {
	var resp TbkJuTqgGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.tbk.ju.tqg.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TbkJuTqgGetResponse struct {
	Results      []*Results `json:"results"`
	TotalResults int        `json:"total_results"`
}

type TbkJuTqgGetResponseResult struct {
	Response *TbkJuTqgGetResponse `json:"tbk_ju_tqg_get_response"`
}
