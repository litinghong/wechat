package pay

import (
	"encoding/xml"
	"github.com/silenceper/wechat/util"
)

var orderQueryGateway = "https://api.mch.weixin.qq.com/pay/orderquery"

// 查询订单参数
type orderQueryParam struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	SignType      string `xml:"sign_type"`
	Sign          string `xml:"sign"`
	NonceStr      string `xml:"nonce_str"`
}

// 查询订单结果
type OrderQueryResult struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	AppId         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      string `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
	TradeState    string `xml:"trade_state"`
}

func (pcf *Pay) OrderQuery(outTradeNo string) (result *OrderQueryResult, err error) {
	nonceStr := util.RandomStr(32)

	request := orderQueryParam{
		AppID:      pcf.Context.AppID,
		MchID:      pcf.Context.PayMchID,
		OutTradeNo: outTradeNo,
		SignType:   "MD5",
		NonceStr:   nonceStr,
	}

	param := make(map[string]interface{})
	param["appid"] = pcf.AppID
	param["mch_id"] = pcf.PayMchID
	param["out_trade_no"] = outTradeNo
	param["sign_type"] = "MD5"
	param["nonce_str"] = nonceStr

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)
	request.Sign = sign

	rawRet, err := util.PostXML(orderQueryGateway, request)

	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(rawRet, &result)
	if err != nil {
		return
	}

	return
}
